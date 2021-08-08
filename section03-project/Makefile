NAME=$(shell basename $(shell pwd))
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api/biz -name *.proto)

.PHONY: init
# init env
init:
	go get git.mysre.cn/yunlian-golang/go-hulk@v0.3.2
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go get -u github.com/google/wire/cmd/wire@v0.5.0

.PHONY: proto
proto:
	protoc --proto_path=./api \
		--go_out=paths=source_relative:./api \
		--go-grpc_out=paths=source_relative:./api \
		--grpc-gateway_out=paths=source_relative:./api \
		$(API_PROTO_FILES)

.PHONY: grpc
# generate grpc proto
grpc:
	 protoc --proto_path=./api \
           --go_out=paths=source_relative:./api \
           --go-grpc_out=paths=source_relative:./api \
           $(API_PROTO_FILES)

.PHONY: gateway
# generate gateway proto
gateway:
	 protoc --proto_path=./api \
           --go_out=paths=source_relative:./api \
           --grpc-gateway_out=paths=source_relative:./api \
           $(API_PROTO_FILES)
