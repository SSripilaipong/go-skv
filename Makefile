PROTOBUF_FILES = \
	server/dbserver/dbgrpc/main.proto \
	server/dbpeerconnector/peergrpc/main.proto

gen: $(PROTOBUF_FILES)

$(PROTOBUF_FILES):
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $@

.PHONY: gen $(PROTOBUF_FILES)

test:
	go test ./tests/...
