PROTOBUF_FILES = \
	server/dbserver/dbgrpc/main.proto \
	server/dbpeerconnector/peergrpc/main.proto

gen: $(PROTOBUF_FILES)

$(PROTOBUF_FILES):
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $@

.PHONY: gen $(PROTOBUF_FILES)

test:
	go test ./tests/...

server1:
	go run server.go start --db-port 5555 --peer-port 5556 --peers=localhost:6556

server2:
	go run server.go start --db-port 6555 --peer-port 6551 --peers=localhost:5556
