package peergrpccontroller

import (
	"context"
	"fmt"
	"go-skv/server/dbpeerconnector/peergrpc"
	"go-skv/util/goutil"
	"go-skv/util/grpcutil"
	"google.golang.org/grpc"
	"net"
)

func (c *controller) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.port))
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}

	c.updatePort(lis)

	grpcServer := grpc.NewServer()
	peergrpc.RegisterPeerServiceServer(grpcServer, newGrpcImplementation(c.usecase))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(fmt.Errorf("unhandled error"))
		}
	}()

	c.wg.Add(1)
	go func() {
		select {
		case <-ctx.Done():
			grpcServer.GracefulStop()
			c.wg.Done()
		}
	}()

	return nil
}

func (c *controller) updatePort(lis net.Listener) {
	port, err := grpcutil.GetPortFromAddress(lis.Addr())
	if err != nil {
		goutil.PanicUnhandledError(err)
	}
	c.port = port
}
