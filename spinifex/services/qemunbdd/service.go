package qemunbdd

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mulgadc/spinifex/spinifex/admin"
	"github.com/mulgadc/spinifex/spinifex/ebsprovider"
	"github.com/mulgadc/spinifex/spinifex/ebsprovider/natsserve"
	"github.com/mulgadc/spinifex/spinifex/types"
	"github.com/mulgadc/spinifex/spinifex/utils"
	"github.com/nats-io/nats.go"
)

// serviceName roots this daemon's PID file: baseDir/qemunbd.pid.
var serviceName = "qemunbd"

// Config is the qemunbdd service's startup configuration, mirroring
// viperblockd.Config's NATS and base-directory fields. NodeName scopes the
// natsserve PublishVolume/UnpublishVolume subjects to this node.
type Config struct {
	BaseDir    string
	NatsHost   string
	NatsToken  string
	NatsCACert string
	NodeName   string
	Debug      bool
}

// Service runs the qcow2/qemu-nbd EBSProvider behind natsserve. It and
// viperblockd both answer ebs.provider.v1.*, so only one of the two may be
// pointed at the same NATS cluster at a time.
type Service struct {
	Config *Config
}

func New(config any) (svc *Service, err error) {
	cfg, ok := config.(*Config)
	if !ok {
		return nil, fmt.Errorf("invalid config type for qemunbdd service")
	}
	return &Service{Config: cfg}, nil
}

func (svc *Service) Start() (int, error) {
	if err := utils.WritePidFileTo(svc.Config.BaseDir, serviceName, os.Getpid()); err != nil {
		return 0, fmt.Errorf("write pid file: %w", err)
	}
	if err := launchService(svc.Config); err != nil {
		slog.Error("Failed to launch service", "err", err)
		return 0, err
	}
	return os.Getpid(), nil
}

func (svc *Service) Stop() (err error) {
	return utils.StopProcessAt(svc.Config.BaseDir, serviceName)
}

func (svc *Service) Status() (string, error) {
	return utils.ServiceStatus(svc.Config.BaseDir, serviceName)
}

func (svc *Service) Shutdown() (err error) {
	return svc.Stop()
}

func (svc *Service) Reload() (err error) {
	return nil
}

// launchService connects to NATS, roots a qcow2 provider at cfg.BaseDir and
// serves ebs.provider.v1.* until SIGINT/SIGTERM, then unsubscribes. It blocks
// for the life of the process.
func launchService(cfg *Config) error {
	nc, err := utils.ConnectNATSWithRetry(admin.DialTarget(cfg.NatsHost), cfg.NatsToken, cfg.NatsCACert)
	if err != nil {
		return fmt.Errorf("connect to NATS: %w", err)
	}
	defer nc.Close()

	provider, err := NewProvider(cfg.BaseDir)
	if err != nil {
		return fmt.Errorf("construct qemunbd provider: %w", err)
	}

	stop, err := natsserve.Serve(context.Background(), nc, provider, natsserve.Options{NodeID: cfg.NodeName})
	if err != nil {
		return fmt.Errorf("serve ebs.provider.v1: %w", err)
	}
	stopLegacy, err := serveLegacyVolumeMounts(nc, provider, cfg.NodeName)
	if err != nil {
		stop()
		return fmt.Errorf("serve legacy volume mounts: %w", err)
	}

	slog.Info("qemunbdd: waiting for EBS provider events", "node", cfg.NodeName)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	slog.Info("qemunbdd: shutting down gracefully...")
	stopLegacy()
	stop()
	return nil
}

// serveLegacyVolumeMounts bridges the VM manager's node-scoped mount subjects
// to the provider-neutral publish contract. The VM manager still consumes the
// legacy response shape, but storage implementations remain behind EBSProvider.
func serveLegacyVolumeMounts(nc *nats.Conn, provider ebsprovider.EBSProvider, nodeID string) (func(), error) {
	mountSub, err := nc.Subscribe("ebs."+nodeID+".mount", func(msg *nats.Msg) {
		var req types.EBSRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			respondLegacy(msg, types.EBSMountResponse{Error: err.Error()})
			return
		}
		published, err := provider.PublishVolume(context.Background(), ebsprovider.PublishVolumeRequest{
			Versioned: ebsprovider.NewVersioned(),
			VolumeID:  req.Name,
			NodeID:    nodeID,
		})
		if err != nil {
			respondLegacy(msg, types.EBSMountResponse{Error: err.Error()})
			return
		}
		respondLegacy(msg, types.EBSMountResponse{URI: published.NBDURI, Mounted: true})
	})
	if err != nil {
		return nil, err
	}

	unmountSub, err := nc.Subscribe("ebs."+nodeID+".unmount", func(msg *nats.Msg) {
		var req types.EBSRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			respondLegacy(msg, types.EBSUnMountResponse{Volume: req.Name, Error: err.Error()})
			return
		}
		err := provider.UnpublishVolume(context.Background(), ebsprovider.UnpublishVolumeRequest{
			Versioned: ebsprovider.NewVersioned(),
			VolumeID:  req.Name,
			NodeID:    nodeID,
		})
		if err != nil {
			respondLegacy(msg, types.EBSUnMountResponse{Volume: req.Name, Error: err.Error()})
			return
		}
		respondLegacy(msg, types.EBSUnMountResponse{Volume: req.Name})
	})
	if err != nil {
		_ = mountSub.Unsubscribe()
		return nil, err
	}

	deleteSub, err := nc.Subscribe("ebs.delete", func(msg *nats.Msg) {
		var req types.EBSDeleteRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			respondLegacy(msg, types.EBSDeleteResponse{Volume: req.Volume, Error: err.Error()})
			return
		}
		err := provider.DeleteVolume(context.Background(), ebsprovider.DeleteVolumeRequest{
			Versioned: ebsprovider.NewVersioned(),
			VolumeID:  req.Volume,
		})
		if err != nil {
			respondLegacy(msg, types.EBSDeleteResponse{Volume: req.Volume, Error: err.Error()})
			return
		}
		respondLegacy(msg, types.EBSDeleteResponse{Volume: req.Volume, Success: true})
	})
	if err != nil {
		_ = mountSub.Unsubscribe()
		_ = unmountSub.Unsubscribe()
		return nil, err
	}

	return func() {
		_ = mountSub.Unsubscribe()
		_ = unmountSub.Unsubscribe()
		_ = deleteSub.Unsubscribe()
	}, nil
}

func respondLegacy(msg *nats.Msg, response any) {
	data, err := json.Marshal(response)
	if err != nil {
		slog.Error("qemunbdd: encode legacy EBS response", "err", err)
		return
	}
	if err := msg.Respond(data); err != nil {
		slog.Error("qemunbdd: send legacy EBS response", "err", err)
	}
}
