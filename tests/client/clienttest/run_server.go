package clienttest

import (
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc"
	"net"
)

func RunServerWithService(service dbgrpc.DbServiceServer, execute func(net.Addr)) {
	lis, err := net.Listen("tcp", ":0")
	goutil.PanicUnhandledError(err)

	server := grpc.NewServer()
	defer server.GracefulStop()

	dbgrpc.RegisterDbServiceServer(server, service)
	go func() {
		goutil.PanicUnhandledError(server.Serve(lis))
	}()

	execute(lis.Addr())
}
