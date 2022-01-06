package template

var (
	HandlerSRV = `package handler

type Handler struct{}
`

	HandlerAPISRV = `package handler

import (
	"context"

	"comm/logger"
	"comm/auth"

	{{dehyphen .Alias}} "proto/{{.Dir}}"
)

// Call is a single request handler called via client.Call or the generated client code
func (h *Handler) Call(ctx context.Context, req *{{dehyphen .Alias}}.Request, rsp *{{dehyphen .Alias}}.Response) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Call", acc.Name)
	}
	logger.Info(ctx, "Received {{title .Alias}}.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
`
)
