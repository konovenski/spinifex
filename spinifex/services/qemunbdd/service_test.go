package qemunbdd

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mulgadc/spinifex/spinifex/ebsprovider"
	"github.com/mulgadc/spinifex/spinifex/testutil"
	"github.com/mulgadc/spinifex/spinifex/types"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLaunchService_ServesRealProviderOverNATS is launchService's one piece
// of real logic: it must root an actual qcow2 Provider at cfg.BaseDir and
// answer capabilities with that provider's own, not a stub's.
func TestLaunchService_ServesRealProviderOverNATS(t *testing.T) {
	ns, client := testutil.StartTestNATS(t)

	cfg := &Config{
		NatsHost: ns.ClientURL(),
		BaseDir:  t.TempDir(),
		NodeName: "test-node",
	}
	go func() { _ = launchService(cfg) }()

	reqBody, err := json.Marshal(ebsprovider.GetCapabilitiesRequest{Versioned: ebsprovider.NewVersioned()})
	require.NoError(t, err)

	var msg *nats.Msg
	require.Eventually(t, func() bool {
		m, reqErr := client.Request(ebsprovider.CapabilitiesSubject, reqBody, time.Second)
		if reqErr != nil {
			return false
		}
		msg = m
		return true
	}, 5*time.Second, 50*time.Millisecond, "qemunbdd never answered %s", ebsprovider.CapabilitiesSubject)

	var resp ebsprovider.GetCapabilitiesResponse
	require.NoError(t, json.Unmarshal(msg.Data, &resp))
	require.Nil(t, resp.Error)
	assert.Equal(t, capabilities, resp.Capabilities)
}

// TestLaunchService_ScopesPublishToNodeID confirms cfg.NodeName reaches
// natsserve.Options.NodeID: PublishVolume must be reachable on this node's
// own subject, not the wildcard every-node form.
func TestLaunchService_ScopesPublishToNodeID(t *testing.T) {
	ns, client := testutil.StartTestNATS(t)

	cfg := &Config{
		NatsHost: ns.ClientURL(),
		BaseDir:  t.TempDir(),
		NodeName: "node-a",
	}
	go func() { _ = launchService(cfg) }()

	publishSubject, err := ebsprovider.PublishSubject(cfg.NodeName)
	require.NoError(t, err)

	reqBody, err := json.Marshal(ebsprovider.PublishVolumeRequest{
		Versioned: ebsprovider.NewVersioned(),
		VolumeID:  "vol-does-not-exist",
		NodeID:    cfg.NodeName,
	})
	require.NoError(t, err)

	var msg *nats.Msg
	require.Eventually(t, func() bool {
		m, reqErr := client.Request(publishSubject, reqBody, time.Second)
		if reqErr != nil {
			return false
		}
		msg = m
		return true
	}, 5*time.Second, 50*time.Millisecond, "qemunbdd never answered %s", publishSubject)

	var resp ebsprovider.PublishVolumeResponse
	require.NoError(t, json.Unmarshal(msg.Data, &resp))
	// The volume does not exist, so the request fails, but a reply at all
	// proves the node-scoped subject (not the wildcard) is what's subscribed.
	require.NotNil(t, resp.Error)
	assert.Equal(t, ebsprovider.ErrorCodeNotFound, resp.Error.Code)
}

func TestServeLegacyVolumeMountsBridgesPublishLifecycle(t *testing.T) {
	_, client := testutil.StartTestNATS(t)
	provider := ebsprovider.NewMemoryProvider(ebsprovider.Capabilities{VolumeEnumeration: true})
	_, err := provider.CreateVolume(t.Context(), ebsprovider.CreateVolumeRequest{
		Versioned:     ebsprovider.NewVersioned(),
		VolumeID:      "vol-legacy",
		CapacityRange: ebsprovider.CapacityRange{RequiredBytes: 1024},
	})
	require.NoError(t, err)

	stop, err := serveLegacyVolumeMounts(client, provider, "node-a")
	require.NoError(t, err)
	t.Cleanup(stop)
	require.NoError(t, client.Flush())

	mountBody, err := json.Marshal(types.EBSRequest{Name: "vol-legacy"})
	require.NoError(t, err)
	msg, err := client.Request("ebs.node-a.mount", mountBody, time.Second)
	require.NoError(t, err)
	var mounted types.EBSMountResponse
	require.NoError(t, json.Unmarshal(msg.Data, &mounted))
	require.Empty(t, mounted.Error)
	assert.True(t, mounted.Mounted)
	assert.Equal(t, "nbd+unix:///?socket=/memory/vol-legacy.sock", mounted.URI)

	msg, err = client.Request("ebs.node-a.unmount", mountBody, time.Second)
	require.NoError(t, err)
	var unmounted types.EBSUnMountResponse
	require.NoError(t, json.Unmarshal(msg.Data, &unmounted))
	require.Empty(t, unmounted.Error)
	assert.False(t, unmounted.Mounted)
	assert.Equal(t, "vol-legacy", unmounted.Volume)

	volume, err := provider.GetVolume(t.Context(), ebsprovider.GetVolumeRequest{
		Versioned: ebsprovider.NewVersioned(),
		VolumeID:  "vol-legacy",
	})
	require.NoError(t, err)
	assert.Equal(t, ebsprovider.VolumeStateAvailable, volume.State)

	deleteBody, err := json.Marshal(types.EBSDeleteRequest{Volume: "vol-legacy"})
	require.NoError(t, err)
	msg, err = client.Request("ebs.delete", deleteBody, time.Second)
	require.NoError(t, err)
	var deleted types.EBSDeleteResponse
	require.NoError(t, json.Unmarshal(msg.Data, &deleted))
	require.Empty(t, deleted.Error)
	assert.True(t, deleted.Success)

	volumes, err := provider.ListVolumes(t.Context(), ebsprovider.ListVolumesRequest{
		Versioned: ebsprovider.NewVersioned(),
	})
	require.NoError(t, err)
	assert.Empty(t, volumes.Volumes)
}
