BUF_VERSION := v1.0.0-rc5
PROTOC_VERSION := 3.18.1
PROTOC_GEN_GO_VERSION := v1.27.1
PROTOC_GEN_GO_GRPC_VERSION := v1.1.0
PROTOC_GEN_GRPC_WEB_VERSION := 1.3.0
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

tools:
	mkdir -p $(TOOLDIR)
	cd $(TOOLDIR); curl -L -o pb.zip https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(OS_TARGET)-$(OS_ARCH).zip
	cd $(TOOLDIR); unzip -o pb.zip && mv bin/protoc .
	cd $(TOOLDIR); curl -L -o protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/$(PROTOC_GEN_GRPC_WEB_VERSION)/protoc-gen-grpc-web-$(PROTOC_GEN_GRPC_WEB_VERSION)-$(OS_TARGET)-$(OS_ARCH)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@$(BUF_VERSION) github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@$(BUF_VERSION)
	go install github.com/xo/xo@$(XO_VERSION)

.PHONY: lint
lint: tools
	buf lint proto/

.PHONY: generate
generate: lint
	buf generate proto/
gendb:
	mkdir -p pkg/gen/xo/captureamoment
	xo schema mysql://root@localhost/captureamoment -o pkg/gen/xo/captureamoment
