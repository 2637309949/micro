package handler

import (
	"github.com/2637309949/micro/v3/service/api"
	"github.com/2637309949/micro/v3/service/client"
)

type Context interface {
	Client() client.Client
	Service() *api.Service
	Domain() string
}
