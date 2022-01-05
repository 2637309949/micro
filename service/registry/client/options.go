package client

import (
	"context"

	"github.com/2637309949/micro/v3/service/client"
	"github.com/2637309949/micro/v3/service/registry"
)

type clientKey struct{}

// WithClient sets the RPC client
func WithClient(c client.Client) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, clientKey{}, c)
	}
}
