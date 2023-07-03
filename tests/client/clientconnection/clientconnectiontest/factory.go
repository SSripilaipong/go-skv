package clientconnectiontest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbgrpc"
)

type DbServiceServerMock struct {
	dbgrpc.UnimplementedDbServiceServer
	GetValue_Request *dbgrpc.GetValueRequest
	GetValue_Return  *dbgrpc.GetValueResponse
	SetValue_Request *dbgrpc.SetValueRequest
}

func (s *DbServiceServerMock) GetValue(_ context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	s.GetValue_Request = request
	response := goutil.Coalesce(s.GetValue_Return, &dbgrpc.GetValueResponse{})
	return response, nil
}

func (s *DbServiceServerMock) SetValue(_ context.Context, request *dbgrpc.SetValueRequest) (*dbgrpc.SetValueResponse, error) {
	s.SetValue_Request = request
	return &dbgrpc.SetValueResponse{}, nil
}
