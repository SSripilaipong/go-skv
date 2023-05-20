package dbserverTest

import (
	"fmt"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/tests/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectWithPort(port int, execute func(client dbgrpc.DbServiceClient) error) error {
	conn, err := grpc.Dial(grpcTest.LocalAddress(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("unexpected error"))
	}
	defer func() { _ = conn.Close() }()
	return execute(dbgrpc.NewDbServiceClient(conn))
}
