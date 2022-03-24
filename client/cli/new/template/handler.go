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

// DeleteInfo defined TODO
func (h *Handler) DeleteInfo(ctx context.Context, req *{{dehyphen .Alias}}.DeleteInfoRequest, rsp *{{dehyphen .Alias}}.DeleteInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteInfo", acc.Name)
	}
	return nil
}

// UpdateInfo defined TODO
func (h *Handler) UpdateInfo(ctx context.Context, req *{{dehyphen .Alias}}.UpdateInfoRequest, rsp *{{dehyphen .Alias}}.UpdateInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do UpdateInfo", acc.Name)
	}
	return nil
}

// InsertInfo defined TODO
func (h *Handler) InsertInfo(ctx context.Context, req *{{dehyphen .Alias}}.InsertInfoRequest, rsp *{{dehyphen .Alias}}.InsertInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do InsertInfo", acc.Name)
	}
	return nil
}

// QueryInfoDetail defined TODO
func (h *Handler) QueryInfoDetail(ctx context.Context, req *{{dehyphen .Alias}}.QueryInfoDetailRequest, rsp *{{dehyphen .Alias}}.QueryInfoDetailResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfoDetail", acc.Name)
	}
	return nil
}

// QueryInfo defined TODO
func (h *Handler) QueryInfo(ctx context.Context, req *{{dehyphen .Alias}}.QueryInfoRequest, rsp *{{dehyphen .Alias}}.QueryInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfo", acc.Name)
	}
	return nil
}
`
)
