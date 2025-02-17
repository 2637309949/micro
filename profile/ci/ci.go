// Package ci is for continuous integration testing
package ci

import (
	"github.com/2637309949/micro/v3/plugin/etcd"
	"github.com/2637309949/micro/v3/profile"
	"github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/service/auth/jwt"
	memBroker "github.com/2637309949/micro/v3/service/broker/memory"
	"github.com/2637309949/micro/v3/service/config"
	storeConfig "github.com/2637309949/micro/v3/service/config/store"
	"github.com/2637309949/micro/v3/service/events"
	evStore "github.com/2637309949/micro/v3/service/events/store"
	memStream "github.com/2637309949/micro/v3/service/events/stream/memory"
	"github.com/2637309949/micro/v3/service/logger"
	microRuntime "github.com/2637309949/micro/v3/service/runtime"
	"github.com/2637309949/micro/v3/service/runtime/local"
	"github.com/2637309949/micro/v3/service/store"
	"github.com/2637309949/micro/v3/service/store/file"
	"github.com/urfave/cli/v2"
)

func init() {
	profile.Register("ci", Profile)
}

// CI profile to use for CI tests
var Profile = &profile.Profile{
	Name: "ci",
	Setup: func(ctx *cli.Context) error {
		auth.DefaultAuth = jwt.NewAuth()
		microRuntime.DefaultRuntime = local.NewRuntime()
		store.DefaultStore = file.NewStore()
		config.DefaultConfig, _ = storeConfig.NewConfig(store.DefaultStore, "")
		events.DefaultStream, _ = memStream.NewStream()
		events.DefaultStore = evStore.NewStore(evStore.WithStore(store.DefaultStore))
		profile.SetupBroker(memBroker.NewBroker())
		profile.SetupRegistry(etcd.NewRegistry())
		profile.SetupJWT(ctx)
		profile.SetupConfigSecretKey(ctx)

		var err error
		store.DefaultBlobStore, err = file.NewBlobStore()
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		return nil
	},
}
