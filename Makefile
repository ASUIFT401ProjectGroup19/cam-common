BUF_VERSION := v1.0.0-rc5
PROTOC_VERSION := 3.18.1
PROTOC_GEN_GO_VERSION := v1.27.1
PROTOC_GEN_GO_GRPC_VERSION := v1.1.0
PROTOC_GEN_GRPC_WEB_VERSION := 1.3.0
PROTOC_GEN_VALIDATE_VERSION := v0.6.2
PROTODEP_VERSION := 0.1.5
TOOLDIR := .bin
XO_VERSION := latest

OS := $(shell uname -s)
ifeq ($(OS),Linux)
	OS_TARGET := linux
endif
ifeq ($(OS),Darwin)
	OS_TARGET := osx
endif
OS_ARCH := $(shell uname -p)

export GOBIN=$(abspath $(TOOLDIR))
export PATH := $(GOBIN):$(PATH)

toolbox:
	mkdir -p $(TOOLDIR)

buf: toolbox
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@$(BUF_VERSION) github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@$(BUF_VERSION)

grpcurl: toolbox
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

protoc: toolbox
	cd $(TOOLDIR); curl -L -o pb.zip https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(OS_TARGET)-$(OS_ARCH).zip
	cd $(TOOLDIR); unzip -o pb.zip && mv bin/protoc .

protoc-gen-validate: toolbox
	go install github.com/envoyproxy/protoc-gen-validate@$(PROTOC_GEN_VALIDATE_VERSION)

protoc-go: toolbox
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

protoc-js: toolbox
	echo https://github.com/grpc/grpc-web/releases/download/$(PROTOC_GEN_GRPC_WEB_VERSION)/protoc-gen-grpc-web-$(PROTOC_GEN_GRPC_WEB_VERSION)-$(OS_TARGET)-$(OS_ARCH)
	cd $(TOOLDIR); curl -L -o protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/$(PROTOC_GEN_GRPC_WEB_VERSION)/protoc-gen-grpc-web-$(PROTOC_GEN_GRPC_WEB_VERSION)-$(OS_TARGET)-$(OS_ARCH)

protodep: toolbox
	go install github.com/stormcat24/protodep@$(PROTODEP_VERSION)

xo: toolbox
	go install github.com/xo/xo@$(XO_VERSION)

prototools: buf protoc protoc-gen-validate protoc-go protoc-js protodep

dbtools: xo

tools: prototools dbtools grpcurl

protoreqs: prototools
	protodep up --use-https -f

lint: prototools
	buf lint proto/

genproto: protoreqs lint
	buf generate proto/

gendb: dbtools
	mkdir -p pkg/gen/xo/captureamoment
	xo schema mysql://root:localpassword@localhost/captureamoment -o pkg/gen/xo/captureamoment

generate: genproto gendb
