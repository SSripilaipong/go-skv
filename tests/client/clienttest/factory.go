package clienttest

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
)

type DbServiceServerMock struct {
	dbgrpc.UnimplementedDbServiceServer
	GetValue_Request *dbgrpc.GetValueRequest
}

func (s *DbServiceServerMock) GetValue(_ context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	s.GetValue_Request = request
	return &dbgrpc.GetValueResponse{}, nil
}
