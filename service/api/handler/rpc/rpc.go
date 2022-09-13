// Copyright 2020 Asim Aslam
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Original source: github.com/micro/go-micro/v3/api/handler/rpc/rpc.go

// Package rpc is a go-micro rpc handler.
package rpc

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/2637309949/micro/v3/service/api"
	"github.com/2637309949/micro/v3/service/api/handler"
	"github.com/2637309949/micro/v3/service/client"
	"github.com/2637309949/micro/v3/service/errors"
	"github.com/2637309949/micro/v3/service/logger"
	"github.com/2637309949/micro/v3/util/codec/bytes"
	"github.com/2637309949/micro/v3/util/ctx"
	xhttp "github.com/2637309949/micro/v3/util/http"
)

const (
	Handler = "rpc"
)

var (
	// supported json codecs
	jsonCodecs = []string{
		"application/grpc+json",
		"application/json",
		"application/json-rpc",
	}

	// support proto codecs
	protoCodecs = []string{
		"application/grpc",
		"application/grpc+proto",
		"application/proto",
		"application/protobuf",
		"application/proto-rpc",
		"application/octet-stream",
	}
)

type rpcHandler struct {
	opts handler.Options
	s    *api.Service
}

func (h *rpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bsize := handler.DefaultMaxRecvSize
	if h.opts.MaxRecvSize > 0 {
		bsize = h.opts.MaxRecvSize
	}

	r.Body = http.MaxBytesReader(w, r.Body, bsize)

	defer r.Body.Close()
	var service *api.Service
	var c client.Client

	if v, ok := r.Context().(handler.Context); ok {
		// we were given the service
		service = v.Service()
		c = v.Client()
	} else if h.opts.Router != nil {
		// try get service from router
		s, err := h.opts.Router.Route(r)
		if err != nil {
			err = errors.InternalServerError("go.micro.api", err.Error())
			xhttp.WriteError(w, r, err)
			return
		}
		service = s
		c = h.opts.Client
	} else {
		// we have no way of routing the request
		err := errors.InternalServerError("go.micro.api", "no route found")
		xhttp.WriteError(w, r, err)
		return
	}

	ct := r.Header.Get("Content-Type")

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	// create context
	cx := ctx.FromRequest(r)

	// set merged context to request
	*r = *r.Clone(cx)
	// if stream we currently only support json
	if isStream(r, service) {
		serveStream(cx, w, r, service, c)
		return
	}

	// create custom router
	var nodes []string
	for _, service := range service.Services {
		for _, node := range service.Nodes {
			nodes = append(nodes, node.Address)
		}
	}
	callOpt := client.WithAddress(nodes...)

	// walk the standard call path
	// get payload
	br, err := api.RequestPayload(r)
	if err != nil {
		err = errors.InternalServerError("go.micro.api", err.Error())
		xhttp.WriteError(w, r, err)
		return
	}

	var rsp []byte

	switch {
	// proto codecs
	case hasCodec(ct, protoCodecs):
		var request *bytes.Frame
		// if the extracted payload isn't empty lets use it
		if len(br) > 0 {
			request = &bytes.Frame{Data: br}
		}

		// create the request
		req := c.NewRequest(
			service.Name,
			service.Endpoint.Name,
			request,
			client.WithContentType(ct),
		)

		// make the call
		var response *bytes.Frame
		if err := c.Call(cx, req, response, callOpt); err != nil {
			xhttp.WriteError(w, r, err)
			return
		}
		rsp = response.Data
	default:
		// if json codec is not present set to json
		if !hasCodec(ct, jsonCodecs) {
			ct = "application/json"
		}

		// default to trying json
		var request json.RawMessage
		// if the extracted payload isn't empty lets use it
		if len(br) > 0 {
			request = json.RawMessage(br)
		}

		// create request/response
		var response json.RawMessage

		req := c.NewRequest(
			service.Name,
			service.Endpoint.Name,
			&request,
			client.WithContentType(ct),
		)
		// make the call
		if err := c.Call(cx, req, &response, callOpt); err != nil {
			xhttp.WriteError(w, r, err)
			return
		}

		// marshall response see https://play.golang.org/p/oBNxUjVTzus
		rsp, err = response.MarshalJSON()
		if err != nil {
			err = errors.InternalServerError("go.micro.api", err.Error())
			xhttp.WriteError(w, r, err)
			return
		}
		rsp = xhttp.Marshal(cx, rsp)
	}

	// write the response
	writeResponse(w, r, rsp, ct)
}

func (rh *rpcHandler) String() string {
	return "rpc"
}

func hasCodec(ct string, codecs []string) bool {
	for _, codec := range codecs {
		if ct == codec {
			return true
		}
	}
	return false
}

func writeResponse(w http.ResponseWriter, r *http.Request, rsp []byte, ct string) {
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Length", strconv.Itoa(len(rsp)))

	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "application/json")
	}

	// Set trailers
	if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
		w.Header().Set("Trailer", "grpc-status")
		w.Header().Set("Trailer", "grpc-message")
		w.Header().Set("grpc-status", "0")
		w.Header().Set("grpc-message", "")
	}

	// write 204 status if rsp is nil
	if len(rsp) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	// write response
	_, err := w.Write(rsp)
	if err != nil {
		if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
			logger.Error(err)
		}
	}
}

func NewHandler(opts ...handler.Option) handler.Handler {
	options := handler.NewOptions(opts...)
	return &rpcHandler{
		opts: options,
	}
}

func WithService(s *api.Service, opts ...handler.Option) handler.Handler {
	options := handler.NewOptions(opts...)
	return &rpcHandler{
		opts: options,
		s:    s,
	}
}
