gen:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative server/dbserver/dbgrpc/main.proto
test:
	go test ./tests/...
