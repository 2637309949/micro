module github.com/2637309949/micro/service/runtime/cells/go/util/entrypoint

go 1.14

require github.com/2637309949/micro/v3 v3.8.8

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
