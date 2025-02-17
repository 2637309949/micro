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
// Original source: github.com/micro/go-micro/v3/api/handler/api.go

// Package api provides an http-rpc handler which provides the entire http request over rpc
package api

import (
	"net/http"

	api "github.com/2637309949/micro/v3/proto/api"
	goapi "github.com/2637309949/micro/v3/service/api"
	"github.com/2637309949/micro/v3/service/api/handler"
	"github.com/2637309949/micro/v3/service/client"
	"github.com/2637309949/micro/v3/service/errors"
	xhttp "github.com/2637309949/micro/v3/util/http"
	"github.com/2637309949/micro/v3/util/router"
)

type apiHandler struct {
	opts handler.Options
	s    *goapi.Service
}

const (
	Handler = "api"
)

// API handler is the default handler which takes api.Request and returns api.Response
func (a *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bsize := handler.DefaultMaxRecvSize
	if a.opts.MaxRecvSize > 0 {
		bsize = a.opts.MaxRecvSize
	}

	r.Body = http.MaxBytesReader(w, r.Body, bsize)
	request, err := requestToProto(r)
	if err != nil {
		err := errors.InternalServerError("go.micro.api", err.Error())
		xhttp.WriteError(w, r, err)
		return
	}

	var service *goapi.Service
	var c client.Client

	if v, ok := r.Context().(handler.Context); ok {
		// we were given the service
		service = v.Service()
		c = v.Client()
	} else if a.opts.Router != nil {
		// try get service from router
		s, err := a.opts.Router.Route(r)
		if err != nil {
			err := errors.InternalServerError("go.micro.api", err.Error())
			xhttp.WriteError(w, r, err)
			return
		}
		service = s
		c = a.opts.Client
	} else {
		// we have no way of routing the request
		er := errors.InternalServerError("go.micro.api", "no route found")
		xhttp.WriteError(w, r, er)
		return
	}

	// create request and response
	req := c.NewRequest(service.Name, service.Endpoint.Name, request)
	rsp := &api.Response{}

	if err := c.Call(r.Context(), req, rsp, client.WithRouter(router.New(service.Services))); err != nil {
		xhttp.WriteError(w, r, err)
		return
	} else if rsp.StatusCode == 0 {
		rsp.StatusCode = http.StatusOK
	}

	for _, header := range rsp.GetHeader() {
		for _, val := range header.Values {
			w.Header().Add(header.Key, val)
		}
	}

	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(int(rsp.StatusCode))
	body := xhttp.Marshal(r.Context(), []byte(rsp.Body))
	w.Write(body)
}

func (a *apiHandler) String() string {
	return "api"
}

func NewHandler(opts ...handler.Option) handler.Handler {
	options := handler.NewOptions(opts...)
	return &apiHandler{
		opts: options,
	}
}
