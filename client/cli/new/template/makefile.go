package template

var (
	Makefile = `include ../.env
NAME := {{.Alias}}
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' ../.env)
GOPATH:=$(shell go env GOPATH)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_IMPORT=github.com/2637309949/go-service
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_IMPORT).BuildDate=$(BUILD_DATE) -X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG)
PROTO_FLAGS=--go_opt=paths=source_relative --micro_opt=paths=source_relative
PROTO_PATH=..:.
SRC_DIR=..
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))
.PHONY: all
all: proto mod api build

.MOD: mod
mod:
	@go mod tidy

.PHONY: proto
proto:
	@find ../proto/ -name '*.proto' -exec protoc --proto_path=$(PROTO_PATH) $(PROTO_FLAGS) --micro_out=$(SRC_DIR) --go_out=:$(SRC_DIR) {} \;

.PHONY: api
api:
	@protoc --openapi_out=. --proto_path=.. $(wildcard ../proto/$(NAME)/*.proto)
	@mv api-$(NAME).json api.json

.PHONY: build
build:
	@go build -a -installsuffix cgo -ldflags "-s -w ${LDFLAGS}" -o $(NAME)

.PHONY: up
up:
	@go run . --service_name=$(NAME) --service_version=

.PHONY: test
test:
	@go test -v ./... -cover

clean:
	@rm -rf $(NAME)

.PHONY: docker
docker:
	@docker build . -t $(NAME):latest	
`
)
