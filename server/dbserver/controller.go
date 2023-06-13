package dbserver

import (
	"context"
	"fmt"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
	"go-skv/util/goutil"
)

type Controller struct {
	dbgrpc.UnimplementedDbServiceServer
	usecase dbusecase.Interface
}

func NewController(usecase dbusecase.Interface) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) GetValue(ctx context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	result, err := c.usecase.GetValue(ctx, dbusecase.GetValueRequest{Key: request.Key})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return &dbgrpc.GetValueResponse{Value: goutil.Pointer(result.Value)}, nil
}

func (c *Controller) SetValue(ctx context.Context, request *dbgrpc.SetValueRequest) (*dbgrpc.SetValueResponse, error) {
	_, err := c.usecase.SetValue(ctx, dbusecase.SetValueRequest{Key: request.Key, Value: request.Value})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return &dbgrpc.SetValueResponse{}, nil
}
