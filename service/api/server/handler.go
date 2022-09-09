package server

import (
	"context"
	"net/http"

	"github.com/2637309949/micro/v3/service/api"
	"github.com/2637309949/micro/v3/service/api/handler"
	"github.com/2637309949/micro/v3/service/api/handler/event"
	"github.com/2637309949/micro/v3/service/api/handler/rpc"
	"github.com/2637309949/micro/v3/service/api/handler/web"
	"github.com/2637309949/micro/v3/service/api/router"
	"github.com/2637309949/micro/v3/service/client"
	"github.com/2637309949/micro/v3/service/errors"

	// TODO: only import handler package
	aapi "github.com/2637309949/micro/v3/service/api/handler/api"
	ahttp "github.com/2637309949/micro/v3/service/api/handler/http"
	xhttp "github.com/2637309949/micro/v3/util/http"
)

type metaHandler struct {
	c  client.Client
	r  router.Router
	ns string
}

var (
	// built in handlers
	handlers = map[string]handler.Handler{
		"rpc":   rpc.NewHandler(),
		"web":   web.NewHandler(),
		"http":  ahttp.NewHandler(),
		"event": event.NewHandler(),
		"api":   aapi.NewHandler(),
	}
)

// Register a handler
func Register(handler string, hd handler.Handler) {
	handlers[handler] = hd
}

// serverContext
type serverContext struct {
	context.Context
	domain  string
	client  client.Client
	service *api.Service
}

func (c *serverContext) Service() *api.Service {
	return c.service
}

func (c *serverContext) Client() client.Client {
	return c.client
}

func (c *serverContext) Domain() string {
	return c.domain
}

func (m *metaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service, err := m.r.Route(r)
	if err != nil {
		err := errors.InternalServerError("go.micro.api", err.Error())
		xhttp.WriteError(w, r, err)
		return
	}

	// inject service into context
	ctx := r.Context()
	// create a new server context
	srvContext := &serverContext{
		Context: ctx,
		domain:  m.ns,
		client:  m.c,
		service: service,
	}
	// clone request with new context
	req := r.Clone(srvContext)

	// get the necessary handler
	hd := service.Endpoint.Handler
	// retrieve the handler for the request
	if len(hd) == 0 {
		hd = "rpc"
	}

	hdr, ok := handlers[hd]
	if !ok {
		// use the catch all rpc handler
		hdr = handlers["rpc"]
	}

	// serve the request
	hdr.ServeHTTP(w, req)
}

// Meta is a http.Handler that routes based on endpoint metadata
func Meta(c client.Client, r router.Router, ns string) http.Handler {
	return &metaHandler{
		c:  c,
		r:  r,
		ns: ns,
	}
}
