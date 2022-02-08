package template

var (
	Module = `module {{.Dir}}-service

go 1.15

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

`
)
