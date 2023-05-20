package dbserver

import (
	"fmt"
	"go-skv/server/dbmanager"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"google.golang.org/grpc"
	"net"
)

func New(port int, dep Dependency) dbmanager.DbServer {
	return &server{port: port, getValueUsecase: dep.GetValueUsecase}
}

type server struct {
	port            int
	grpcServer      *grpc.Server
	getValueUsecase dbusecase.GetValueFunc
}

func (s *server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	s.grpcServer = grpc.NewServer()
	dbgrpc.RegisterDbServiceServer(s.grpcServer, &controller{
		getValueUsecase: s.getValueUsecase,
	})

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			panic(fmt.Errorf("unhandled error"))
		}
	}()

	return nil
}

func (s *server) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}
