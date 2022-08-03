package template

var (
	Module = `module {{.Dir}}-service

go 1.18

require (
	proto v0.0.0-00010101000000-000000000000
	comm v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
)

replace (
	comm => ../comm
	proto => ../proto
)


// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1

`
)
