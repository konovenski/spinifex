package gateway

import (
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

// restRoute binds one HTTP method and chi path pattern to an AWS action and the
// handler that serves it. Params are positional: a handler indexes them in the
// order their {name} segments appear in the pattern.
type restRoute[H any] struct {
	method  string
	pattern string
	action  string
	handler H
}

// restEntry is one registered route plus the param names read off its pattern.
type restEntry[H any] struct {
	route  restRoute[H]
	params []string
}

// restRouter resolves a method and path to a route through a chi trie, which
// orders specific paths ahead of {param} ones itself rather than relying on the
// order the table lists them in.
type restRouter[H any] struct {
	service string
	mux     *chi.Mux
	entries map[string]restEntry[H]
}

// newRESTRouter builds the router for one service's dispatch table. A duplicate
// method+pattern would shadow its twin here as it does in chi; the route tables
// are checked for one by test.
func newRESTRouter[H any](service string, routes []restRoute[H]) *restRouter[H] {
	rr := &restRouter[H]{
		service: service,
		mux:     chi.NewMux(),
		entries: make(map[string]restEntry[H], len(routes)),
	}
	for _, route := range routes {
		rr.entries[restRouteKey(route.method, route.pattern)] = restEntry[H]{
			route:  route,
			params: restPatternParams(route.pattern),
		}
		// The trie is consulted through Match, which never serves the handler:
		// the table above holds the real one.
		rr.mux.MethodFunc(route.method, route.pattern, func(http.ResponseWriter, *http.Request) {})
	}
	return rr
}

// lookup matches method+path, returning the action, path params, and handler, or
// ("", nil, zero, false) on no match. path must be r.URL.EscapedPath(): ARNs in
// path segments are percent-encoded by the CLI (e.g. user%2Fadmin), and matching
// the decoded path would split one param across two segments. Params are
// PathUnescape'd before returning.
func (rr *restRouter[H]) lookup(method, path string) (string, []string, H, bool) {
	var zero H

	rctx := chi.NewRouteContext()
	if !rr.mux.Match(rctx, method, path) {
		return "", nil, zero, false
	}
	entry, ok := rr.entries[restRouteKey(method, rctx.RoutePattern())]
	if !ok {
		slog.Debug(rr.service+": matched route has no table entry", "method", method, "pattern", rctx.RoutePattern())
		return "", nil, zero, false
	}

	var params []string
	if len(entry.params) > 0 {
		params = make([]string, 0, len(entry.params))
		for _, name := range entry.params {
			raw := rctx.URLParam(name)
			// A {name} segment never matches empty, but a trailing wildcard
			// does; every param here stands for a resource identifier.
			if raw == "" {
				return "", nil, zero, false
			}
			decoded, err := url.PathUnescape(raw)
			if err != nil {
				slog.Debug(rr.service+": bad percent-encoding in path param", "param", raw, "err", err)
				decoded = raw
			}
			params = append(params, decoded)
		}
	}
	return entry.route.action, params, entry.route.handler, true
}

// restRouteKey keys the entry table. The trailing slash is trimmed because chi
// reports a pattern registered as "/knowledgebases/" as "/knowledgebases", while
// still matching only the slashed path.
func restRouteKey(method, pattern string) string {
	if len(pattern) > 1 {
		pattern = strings.TrimSuffix(pattern, "/")
	}
	return method + " " + pattern
}

// restPatternParams lists the param names of a chi pattern in path order, the
// order handlers index them in.
func restPatternParams(pattern string) []string {
	var names []string
	for segment := range strings.SplitSeq(pattern, "/") {
		switch {
		case strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}"):
			names = append(names, segment[1:len(segment)-1])
		case segment == "*":
			names = append(names, "*")
		}
	}
	return names
}
