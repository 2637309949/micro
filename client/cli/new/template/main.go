package template

var (
	MainSRV = `package main

import (
	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"

	"github.com/2637309949/micro/v3/service"
	"github.com/2637309949/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("{{lower .Alias}}"),
		service.Version("latest"),
	)

	// Register handler
	pb.Register{{title .Alias}}Handler(srv.Server(), new(handler.{{title .Alias}}))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
`
)
