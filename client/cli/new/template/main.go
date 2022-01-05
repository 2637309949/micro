package template

var (
	MainSRV = `package main

import (
	"comm/logger"
	"comm/service"
	"comm/define"

	"{{.Dir}}-service/handler"
	{{dehyphen .Alias}} "proto/{{.Dir}}"
)

func main() {
	// Create handler
	hdl := handler.Handler{}

	// Create service
	srv := service.New(service.Name("{{lower .Alias}}"))

	// Register handler
	{{dehyphen .Alias}}.Register{{title .Alias}}Handler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
`
)
