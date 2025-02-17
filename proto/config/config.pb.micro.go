// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/config/config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/2637309949/micro/v3/service/api"
	client "github.com/2637309949/micro/v3/service/client"
	server "github.com/2637309949/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Config service

func NewConfigEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Config service

type ConfigService interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...client.CallOption) (*SetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	// These methods are here for backwards compatibility reasons
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
}

type configService struct {
	c    client.Client
	name string
}

func NewConfigService(name string, c client.Client) ConfigService {
	return &configService{
		c:    c,
		name: name,
	}
}

func (c *configService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Config.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) Set(ctx context.Context, in *SetRequest, opts ...client.CallOption) (*SetResponse, error) {
	req := c.c.NewRequest(c.name, "Config.Set", in)
	out := new(SetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Config.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Config.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Config service

type ConfigHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
	Set(context.Context, *SetRequest, *SetResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	// These methods are here for backwards compatibility reasons
	Read(context.Context, *ReadRequest, *ReadResponse) error
}

func RegisterConfigHandler(s server.Server, hdlr ConfigHandler, opts ...server.HandlerOption) error {
	type config interface {
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Set(ctx context.Context, in *SetRequest, out *SetResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
	}
	type Config struct {
		config
	}
	h := &configHandler{hdlr}
	return s.Handle(s.NewHandler(&Config{h}, opts...))
}

type configHandler struct {
	ConfigHandler
}

func (h *configHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ConfigHandler.Get(ctx, in, out)
}

func (h *configHandler) Set(ctx context.Context, in *SetRequest, out *SetResponse) error {
	return h.ConfigHandler.Set(ctx, in, out)
}

func (h *configHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ConfigHandler.Delete(ctx, in, out)
}

func (h *configHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ConfigHandler.Read(ctx, in, out)
}
