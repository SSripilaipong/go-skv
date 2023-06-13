package dbservercontroller

import (
	"fmt"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbserver/dbusecase"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
)

func New(port int, usecase dbusecase.Interface) Interface {
	return &controller{
		port:    port,
		usecase: usecase,
	}
}

type controller struct {
	port    int
	usecase dbusecase.Interface

	grpcServer *grpc.Server
}

func (s *controller) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}

	s.updatePort(lis)

	s.grpcServer = grpc.NewServer()
	dbgrpc.RegisterDbServiceServer(s.grpcServer, newGrpcImplementation(s.usecase))

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			panic(fmt.Errorf("unhandled error"))
		}
	}()

	return nil
}

func (s *controller) updatePort(lis net.Listener) {
	tokens := strings.Split(lis.Addr().String(), ":")
	port, err := strconv.Atoi(tokens[len(tokens)-1])
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	s.port = port
}

func (s *controller) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}

func (s *controller) Port() int {
	return s.port
}
