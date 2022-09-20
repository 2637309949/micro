package main

//go:generate ./scripts/generate.sh

import (
	"github.com/2637309949/micro/v3/cmd"

	// load packages so they can register commands
	_ "github.com/2637309949/micro/v3/client/cli"
	_ "github.com/2637309949/micro/v3/client/web"
	_ "github.com/2637309949/micro/v3/cmd/server"
	_ "github.com/2637309949/micro/v3/cmd/service"
	_ "github.com/2637309949/micro/v3/cmd/usage"
)

func main() {
	cmd.Run()
}
