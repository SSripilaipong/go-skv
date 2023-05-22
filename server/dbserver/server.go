package dbserver

import (
	"fmt"
	"go-skv/server/dbmanager"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
)

func New(port int, dep Dependency) dbmanager.DbServer {
	return &server{
		port:            port,
		getValueUsecase: dep.GetValueUsecase,
		setValueUsecase: dep.SetValueUsecase,
	}
}

type server struct {
	port            int
	getValueUsecase dbusecase.GetValueFunc
	setValueUsecase dbusecase.SetValueFunc

	grpcServer *grpc.Server
}

func (s *server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}

	s.updatePort(lis)

	s.grpcServer = grpc.NewServer()
	dbgrpc.RegisterDbServiceServer(s.grpcServer, &controller{
		getValueUsecase: s.getValueUsecase,
		setValueUsecase: s.setValueUsecase,
	})

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			panic(fmt.Errorf("unhandled error"))
		}
	}()

	return nil
}

func (s *server) updatePort(lis net.Listener) {
	tokens := strings.Split(lis.Addr().String(), ":")
	port, err := strconv.Atoi(tokens[len(tokens)-1])
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	s.port = port
}

func (s *server) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}

func (s *server) Port() int {
	return s.port
}
