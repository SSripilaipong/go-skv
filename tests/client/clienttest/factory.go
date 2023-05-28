package clienttest

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/util/goutil"
)

type DbServiceServerMock struct {
	dbgrpc.UnimplementedDbServiceServer
	GetValue_Request *dbgrpc.GetValueRequest
	GetValue_Return  *dbgrpc.GetValueResponse
}

func (s *DbServiceServerMock) GetValue(_ context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	s.GetValue_Request = request
	response := goutil.Coalesce(s.GetValue_Return, &dbgrpc.GetValueResponse{})
	return response, nil
}
