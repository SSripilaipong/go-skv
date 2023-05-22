package dbserver

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
)

type controller struct {
	dbgrpc.UnimplementedDbServiceServer
	getValueUsecase dbusecase.GetValueFunc
	setValueUsecase dbusecase.SetValueFunc
}

func (c *controller) GetValue(_ context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	_, _ = c.getValueUsecase(&dbusecase.GetValueRequest{Key: request.Key})
	return nil, nil
}

func (c *controller) SetValue(_ context.Context, request *dbgrpc.SetValueRequest) (*dbgrpc.SetValueResponse, error) {
	_, _ = c.setValueUsecase(&dbusecase.SetValueRequest{Key: request.Key, Value: request.Value})
	return nil, nil
}
