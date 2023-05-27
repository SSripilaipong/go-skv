package dbserver

import (
	"context"
	"fmt"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
)

type controller struct {
	dbgrpc.UnimplementedDbServiceServer
	getValueUsecase dbusecase.GetValueFunc
	setValueUsecase dbusecase.SetValueFunc
}

func (c *controller) GetValue(_ context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	result, err := c.getValueUsecase(context.Background(), &dbusecase.GetValueRequest{Key: request.Key})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return &dbgrpc.GetValueResponse{Value: result.Value}, nil
}

func (c *controller) SetValue(_ context.Context, request *dbgrpc.SetValueRequest) (*dbgrpc.SetValueResponse, error) {
	_, err := c.setValueUsecase(context.Background(), &dbusecase.SetValueRequest{Key: request.Key, Value: request.Value})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return nil, nil
}
