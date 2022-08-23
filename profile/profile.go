// Package profile is for specific profiles
// @todo this package is the definition of cruft and
// should be rewritten in a more elegant way
package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/2637309949/micro/v3/plugin/etcd"
	"github.com/2637309949/micro/v3/plugin/postgres"
	"github.com/2637309949/micro/v3/plugin/prometheus"
	redisBroker "github.com/2637309949/micro/v3/plugin/redis/broker"
	redisstream "github.com/2637309949/micro/v3/plugin/redis/stream"
	microAuth "github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/service/auth/jwt"
	"github.com/2637309949/micro/v3/service/auth/noop"
	"github.com/2637309949/micro/v3/service/broker"
	memBroker "github.com/2637309949/micro/v3/service/broker/memory"
	microBuilder "github.com/2637309949/micro/v3/service/build"
	"github.com/2637309949/micro/v3/service/build/golang"
	"github.com/2637309949/micro/v3/service/client"
	grpcClient "github.com/2637309949/micro/v3/service/client/grpc"
	"github.com/2637309949/micro/v3/service/config"
	storeConfig "github.com/2637309949/micro/v3/service/config/store"
	microEvents "github.com/2637309949/micro/v3/service/events"
	evStore "github.com/2637309949/micro/v3/service/events/store"
	memStream "github.com/2637309949/micro/v3/service/events/stream/memory"
	"github.com/2637309949/micro/v3/service/logger"
	"github.com/2637309949/micro/v3/service/metrics"
	"github.com/2637309949/micro/v3/service/model"
	"github.com/2637309949/micro/v3/service/registry"
	"github.com/2637309949/micro/v3/service/registry/memory"
	"github.com/2637309949/micro/v3/service/router"
	k8sRouter "github.com/2637309949/micro/v3/service/router/kubernetes"
	regRouter "github.com/2637309949/micro/v3/service/router/registry"
	microRuntime "github.com/2637309949/micro/v3/service/runtime"
	"github.com/2637309949/micro/v3/service/runtime/kubernetes"
	"github.com/2637309949/micro/v3/service/runtime/local"
	"github.com/2637309949/micro/v3/service/server"
	grpcServer "github.com/2637309949/micro/v3/service/server/grpc"
	microStore "github.com/2637309949/micro/v3/service/store"
	"github.com/2637309949/micro/v3/service/store/file"
	mem "github.com/2637309949/micro/v3/service/store/memory"
	inAuth "github.com/2637309949/micro/v3/util/auth"
	"github.com/2637309949/micro/v3/util/opentelemetry"
	"github.com/2637309949/micro/v3/util/opentelemetry/jaeger"
	"github.com/2637309949/micro/v3/util/user"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/urfave/cli/v2"
)

// profiles which when called will configure micro to run in that environment
var profiles = map[string]*Profile{
	// built in profiles
	"client":     Client,
	"service":    Service,
	"server":     Server,
	"test":       Test,
	"local":      Local,
	"staging":    Staging,
	"kubernetes": Kubernetes,
}

// Profile configures an environment
type Profile struct {
	// name of the profile
	Name string
	// function used for setup
	Setup func(*cli.Context) error
	// TODO: presetup dependencies
	// e.g start resources
}

// Register a profile
func Register(name string, p *Profile) error {
	if _, ok := profiles[name]; ok {
		return fmt.Errorf("profile %s already exists", name)
	}
	profiles[name] = p
	return nil
}

// Load a profile
func Load(name string) (*Profile, error) {
	v, ok := profiles[name]
	if !ok {
		return nil, fmt.Errorf("profile %s does not exist", name)
	}
	return v, nil
}

// Client profile is for any entrypoint that behaves as a client
var Client = &Profile{
	Name:  "client",
	Setup: func(ctx *cli.Context) error { return nil },
}

// Local profile to run as a single process
var Local = &Profile{
	Name: "local",
	Setup: func(ctx *cli.Context) error {
		// set client/server
		client.DefaultClient = grpcClient.NewClient()
		server.DefaultServer = grpcServer.NewServer()

		microAuth.DefaultAuth = jwt.NewAuth()
		microStore.DefaultStore = file.NewStore(file.WithDir(filepath.Join(user.Dir, "server", "store")))
		SetupConfigSecretKey(ctx)
		SetupJWT(ctx)

		// the registry service uses the memory registry, the other core services will use the default
		// rpc client and call the registry service
		if ctx.Args().Get(1) == "registry" {
			SetupRegistry(memory.NewRegistry())
		} else {
			// set the registry address
			registry.DefaultRegistry.Init(
				registry.Addrs("localhost:8000"),
			)
			SetupRegistry(registry.DefaultRegistry)
		}
		config.DefaultConfig, _ = storeConfig.NewConfig(microStore.DefaultStore, "")

		SetupJWT(ctx)
		SetupRegistry(memory.NewRegistry())
		SetupBroker(memBroker.NewBroker())

		// set the store in the model
		model.DefaultModel = model.NewModel(
			model.WithStore(microStore.DefaultStore),
		)

		microRuntime.DefaultRuntime = local.NewRuntime()

		var err error
		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream: %v", err)
		}
		microEvents.DefaultStore = evStore.NewStore(
			evStore.WithStore(microStore.DefaultStore),
		)

		microStore.DefaultBlobStore, err = file.NewBlobStore()
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		return nil
	},
}

// staging profile to run in staging env
// reference: https://github.com/m3o/platform/blob/master/profile/platform/platform.go
var Staging = &Profile{
	Name: "staging",
	Setup: func(ctx *cli.Context) error {
		microAuth.DefaultAuth = jwt.NewAuth()
		// the cockroach store will connect immediately so the address must be passed
		// when the store is created. The cockroach store address contains the location
		// of certs so it can't be defaulted like the broker and registry.
		microStore.DefaultStore = postgres.NewStore(microStore.Nodes(ctx.String("store_address")))
		SetupBroker(redisBroker.NewBroker(broker.Addrs(ctx.String("broker_address"))))
		SetupRegistry(etcd.NewRegistry(registry.Addrs(ctx.String("registry_address"))))
		SetupJWT(ctx)
		SetupConfigSecretKey(ctx)

		// Set up a default metrics reporter (being careful not to clash with any that have already been set):
		if !metrics.IsSet() {
			prometheusReporter, err := prometheus.New()
			if err != nil {
				logger.Fatalf("Error configuring prometheus: %v", err)
			}
			metrics.SetDefaultMetricsReporter(prometheusReporter)
		}

		var err error
		if ctx.Args().Get(1) == "events" {
			microEvents.DefaultStream, err = redisstream.NewStream(redisStreamOpts(ctx)...)
			if err != nil {
				logger.Fatalf("Error configuring stream: %v", err)
			}
			microEvents.DefaultStore = evStore.NewStore(
				evStore.WithStore(microStore.DefaultStore),
			)
		}

		// only configure the blob store for the store and runtime services
		if ctx.Args().Get(1) == "runtime" || ctx.Args().Get(1) == "store" {
			microStore.DefaultBlobStore, err = file.NewBlobStore()
			if err != nil {
				logger.Fatalf("Error configuring file blob store: %v", err)
			}
		}

		// set the store in the model
		model.DefaultModel = model.NewModel(
			model.WithStore(microStore.DefaultStore),
		)

		reporterAddress := os.Getenv("MICRO_TRACING_REPORTER_ADDRESS")
		if len(reporterAddress) == 0 {
			reporterAddress = jaeger.DefaultReporterAddress
		}

		// Configure tracing with Jaeger (forced tracing):
		tracingServiceName := ctx.Args().Get(1)
		if len(tracingServiceName) == 0 {
			tracingServiceName = "Micro"
		}
		openTracer, _, err := jaeger.New(
			opentelemetry.WithServiceName(tracingServiceName),
			opentelemetry.WithSamplingRate(1),
			opentelemetry.WithTraceReporterAddress(reporterAddress),
		)
		if err != nil {
			logger.Fatalf("Error configuring opentracing: %v", err)
		}
		opentelemetry.DefaultOpenTracer = openTracer
		opentracing.SetGlobalTracer(openTracer)

		// use the local runtime, note: the local runtime is designed to run source code directly so
		// the runtime builder should NOT be set when using this implementation
		microRuntime.DefaultRuntime = local.NewRuntime()
		return nil
	},
}

// natsStreamOpts returns a slice of options which should be used to configure nats
func redisStreamOpts(ctx *cli.Context) []redisstream.Option {
	fullAddr := ctx.String("broker_address")
	o, err := redis.ParseURL(fullAddr)
	if err != nil {
		logger.Fatalf("Error configuring redis connection, failed to parse %s", fullAddr)
	}

	opts := []redisstream.Option{
		redisstream.Address(o.Addr),
		redisstream.User(o.Username),
		redisstream.Password(o.Password),
	}
	if o.TLSConfig != nil {
		opts = append(opts, redisstream.TLSConfig(o.TLSConfig))
	}

	return opts
}

// Kubernetes profile to run on kubernetes with zero deps. Designed for use with the micro helm chart
var Kubernetes = &Profile{
	Name: "kubernetes",
	Setup: func(ctx *cli.Context) (err error) {
		microAuth.DefaultAuth = jwt.NewAuth()
		SetupJWT(ctx)

		microRuntime.DefaultRuntime = kubernetes.NewRuntime()
		microBuilder.DefaultBuilder, err = golang.NewBuilder()
		if err != nil {
			logger.Fatalf("Error configuring golang builder: %v", err)
		}

		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream: %v", err)
		}

		microStore.DefaultStore = file.NewStore(file.WithDir("/store"))
		microStore.DefaultBlobStore, err = file.NewBlobStore(file.WithDir("/store/blob"))
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		// set the store in the model
		model.DefaultModel = model.NewModel(
			model.WithStore(microStore.DefaultStore),
		)

		// the registry service uses the memory registry, the other core services will use the default
		// rpc client and call the registry service
		if ctx.Args().Get(1) == "registry" {
			SetupRegistry(memory.NewRegistry())
		}

		// the broker service uses the memory broker, the other core services will use the default
		// rpc client and call the broker service
		if ctx.Args().Get(1) == "broker" {
			SetupBroker(memBroker.NewBroker())
		}

		SetupConfigSecretKey(ctx)

		// Use k8s routing which is DNS based
		router.DefaultRouter = k8sRouter.NewRouter()
		client.DefaultClient.Init(client.Router(router.DefaultRouter))

		// Configure tracing with Jaeger:
		tracingServiceName := ctx.Args().Get(1)
		if len(tracingServiceName) == 0 {
			tracingServiceName = "Micro"
		}
		openTracer, _, err := jaeger.New(
			opentelemetry.WithServiceName(tracingServiceName),
			opentelemetry.WithTraceReporterAddress("localhost:6831"),
		)
		if err != nil {
			logger.Fatalf("Error configuring opentracing: %v", err)
		}
		opentelemetry.DefaultOpenTracer = openTracer

		return nil
	},
}

var Server = &Profile{
	Name: "server",
	Setup: func(ctx *cli.Context) error {
		microAuth.DefaultAuth = jwt.NewAuth()
		microStore.DefaultStore = file.NewStore(file.WithDir(filepath.Join(user.Dir, "server", "store")))
		SetupConfigSecretKey(ctx)
		config.DefaultConfig, _ = storeConfig.NewConfig(microStore.DefaultStore, "")
		SetupJWT(ctx)

		// the registry service uses the memory registry, the other core services will use the default
		// rpc client and call the registry service
		if ctx.Args().Get(1) == "registry" {
			SetupRegistry(memory.NewRegistry())
		} else {
			// set the registry address
			registry.DefaultRegistry.Init(
				registry.Addrs("localhost:8000"),
			)

			SetupRegistry(registry.DefaultRegistry)
		}

		// the broker service uses the memory broker, the other core services will use the default
		// rpc client and call the broker service
		if ctx.Args().Get(1) == "broker" {
			SetupBroker(memBroker.NewBroker())
		} else {
			broker.DefaultBroker.Init(
				broker.Addrs("localhost:8003"),
			)
			SetupBroker(broker.DefaultBroker)
		}

		// set the store in the model
		model.DefaultModel = model.NewModel(
			model.WithStore(microStore.DefaultStore),
		)

		// use the local runtime, note: the local runtime is designed to run source code directly so
		// the runtime builder should NOT be set when using this implementation
		microRuntime.DefaultRuntime = local.NewRuntime()

		var err error
		microEvents.DefaultStream, err = memStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream: %v", err)
		}
		microEvents.DefaultStore = evStore.NewStore(
			evStore.WithStore(microStore.DefaultStore),
		)

		microStore.DefaultBlobStore, err = file.NewBlobStore()
		if err != nil {
			logger.Fatalf("Error configuring file blob store: %v", err)
		}

		// Configure tracing with Jaeger (forced tracing):
		tracingServiceName := ctx.Args().Get(1)
		if len(tracingServiceName) == 0 {
			tracingServiceName = "Micro"
		}
		openTracer, _, err := jaeger.New(
			opentelemetry.WithServiceName(tracingServiceName),
			opentelemetry.WithSamplingRate(1),
		)
		if err != nil {
			logger.Fatalf("Error configuring opentracing: %v", err)
		}
		opentelemetry.DefaultOpenTracer = openTracer

		return nil
	},
}

// Service is the default for any services run
var Service = &Profile{
	Name:  "service",
	Setup: func(ctx *cli.Context) error { return nil },
}

// Test profile is used for the go test suite
var Test = &Profile{
	Name: "test",
	Setup: func(_ *cli.Context) error {
		microAuth.DefaultAuth = noop.NewAuth()
		microStore.DefaultStore = mem.NewStore()
		microStore.DefaultBlobStore, _ = file.NewBlobStore()
		SetupRegistry(memory.NewRegistry())
		// set the store in the model
		model.DefaultModel = model.NewModel(
			model.WithStore(microStore.DefaultStore),
		)
		return nil
	},
}

// SetupRegistry configures the registry
func SetupRegistry(reg registry.Registry) {
	registry.DefaultRegistry = reg
	router.DefaultRouter = regRouter.NewRouter(router.Registry(reg))
	client.DefaultClient.Init(client.Registry(reg), client.Router(router.DefaultRouter))
	server.DefaultServer.Init(server.Registry(reg))
}

// SetupBroker configures the broker
func SetupBroker(b broker.Broker) {
	broker.DefaultBroker = b
	client.DefaultClient.Init(client.Broker(b))
	server.DefaultServer.Init(server.Broker(b))
}

// SetupJWT configures the default internal system rules
func SetupJWT(_ *cli.Context) {
	for _, rule := range inAuth.SystemRules {
		if err := microAuth.Grant(rule); err != nil {
			logger.Fatal("Error creating default rule: %v", err)
		}
	}
}

func SetupConfigSecretKey(ctx *cli.Context) {
	key := ctx.String("config_secret_key")
	if len(key) == 0 {
		k, err := user.GetConfigSecretKey()
		if err != nil {
			logger.Fatal("Error getting config secret: %v", err)
		}
		os.Setenv("MICRO_CONFIG_SECRET_KEY", k)
	}
}
