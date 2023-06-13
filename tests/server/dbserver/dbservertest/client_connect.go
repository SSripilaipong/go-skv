package dbservertest

import (
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/util/goutil"
	"go-skv/util/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectWithPort(port int, execute func(client dbgrpc.DbServiceClient) error) error {
	conn, err := grpc.Dial(grpcutil.LocalAddress(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil.PanicUnhandledError(err)

	defer func() { _ = conn.Close() }()
	return execute(dbgrpc.NewDbServiceClient(conn))
}
