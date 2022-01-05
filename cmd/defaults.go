package cmd

import (
	"github.com/2637309949/micro/v3/service/auth"
	authSrv "github.com/2637309949/micro/v3/service/auth/client"
	"github.com/2637309949/micro/v3/service/broker"
	brokerSrv "github.com/2637309949/micro/v3/service/broker/client"
	"github.com/2637309949/micro/v3/service/client"
	grpcCli "github.com/2637309949/micro/v3/service/client/grpc"
	"github.com/2637309949/micro/v3/service/events"
	eventsSrv "github.com/2637309949/micro/v3/service/events/client"
	"github.com/2637309949/micro/v3/service/metrics"
	noopMet "github.com/2637309949/micro/v3/service/metrics/noop"
	"github.com/2637309949/micro/v3/service/network"
	mucpNet "github.com/2637309949/micro/v3/service/network/mucp"
	"github.com/2637309949/micro/v3/service/registry"
	registrySrv "github.com/2637309949/micro/v3/service/registry/client"
	"github.com/2637309949/micro/v3/service/router"
	routerSrv "github.com/2637309949/micro/v3/service/router/client"
	"github.com/2637309949/micro/v3/service/runtime"
	runtimeSrv "github.com/2637309949/micro/v3/service/runtime/client"
	"github.com/2637309949/micro/v3/service/server"
	grpcSvr "github.com/2637309949/micro/v3/service/server/grpc"
	"github.com/2637309949/micro/v3/service/store"
	storeSrv "github.com/2637309949/micro/v3/service/store/client"
)

// setupDefaults sets the default auth, broker etc implementations incase they arent configured by
// a profile. The default implementations are always the RPC implementations.
func setupDefaults() {
	client.DefaultClient = grpcCli.NewClient()
	server.DefaultServer = grpcSvr.NewServer()
	network.DefaultNetwork = mucpNet.NewNetwork()
	metrics.DefaultMetricsReporter = noopMet.New()

	// setup rpc implementations after the client is configured
	auth.DefaultAuth = authSrv.NewAuth()
	broker.DefaultBroker = brokerSrv.NewBroker()
	events.DefaultStream = eventsSrv.NewStream()
	events.DefaultStore = eventsSrv.NewStore()
	registry.DefaultRegistry = registrySrv.NewRegistry()
	router.DefaultRouter = routerSrv.NewRouter()
	store.DefaultStore = storeSrv.NewStore()
	store.DefaultBlobStore = storeSrv.NewBlobStore()
	runtime.DefaultRuntime = runtimeSrv.NewRuntime()
}
