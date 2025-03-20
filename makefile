BINARY_NAME = docker-dns-sync

# Binaries
PROTOC		?= protoc

# Protobuf definition version (folder)
PROTO_VERSION = v1

.PHONY: build
build:
	go build -C src -o ../bin/${BINARY_NAME}

.PHONY: proto
proto-build-technitium: # https://dev.to/davidsbond/golang-structuring-repositories-with-protocol-buffers-3012
	protoc \
	--proto_path=${PWD}/proto/dns-sync/${PROTO_VERSION} \
	--go_out=${PWD}/pkg/technitium \
	--go_opt=module=github.com/chris-birch/docker-dns-sync/pkg/technitium \
	--go-grpc_opt=module=github.com/chris-birch/docker-dns-sync/pkg/technitium \
	--go-grpc_out=${PWD}/pkg/technitium \
	--proto_path=/usr/local/include \
	${PWD}/proto/dns-sync/${PROTO_VERSION}/*.proto

