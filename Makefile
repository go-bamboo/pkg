GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
PROTO_FILES=$(shell find . -name *.proto)
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)
PWD := $(shell pwd)

.PHONY: init
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

.PHONY: install
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.13
	go install github.com/go-kratos/kratos/cmd/kratos/v2@v2.0.0-20221102101533-2a65502be27b
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@v2.0.0-20221102101533-2a65502be27b
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@v2.0.0-20221102101533-2a65502be27b
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/gnostic@latest
	go install github.com/googleapis/gnostic-go-generator@latest
	go install github.com/google/wire/cmd/wire@latest


.PHONY: upgrade
upgrade:
	kratos upgrade

.PHONY: api
api:
	protoc --proto_path=. \
		   --proto_path=$(KRATOS) \
		   --proto_path=$(KRATOS)/api \
		   --proto_path=$(KRATOS)/third_party \
		   --proto_path=$(PWD)/../third_party \
		   --go_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
		   --go-http_out=paths=source_relative:. \
		   --go-errors_out=paths=source_relative:. \
		   --validate_out=lang=go,paths=source_relative:. \
		   --openapi_out=. \
		   ./api/status/status.proto \
		   ./api/sys/sys.proto

.PHONY: proto
proto:
	protoc --proto_path=. \
		   --proto_path=$(KRATOS) \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(PWD)/../third_party \
           --go_out=paths=source_relative:. \
           --validate_out=lang=go,paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           ./otel/conf.proto \
           ./meta/meta.proto \
           ./registry/conf.proto \
           ./log/core/conf.proto \
           ./rpc/conf.proto \
           ./queue/conf.proto \
           ./rest/conf.proto \
           ./store/mongox/conf.proto \
           ./store/gormx/conf/conf.proto \
           ./store/redis/conf.proto \
           ./fs/s3/conf.proto \
           ./fs/google/conf.proto

.PHONY: swift
swift:
	gnostic openapi.yaml --swift-generator-out=client

.PHONY: gengo
gengo:
	gnostic openapi.yaml --plugin-request-out=.

.PHONY: test
test:
	go test -v ./... -cover
