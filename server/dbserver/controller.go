package dbserver

import (
	"context"
	"fmt"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
)

type Controller struct {
	dbgrpc.UnimplementedDbServiceServer
	getValueUsecase dbusecase.GetValueFunc
	setValueUsecase dbusecase.SetValueFunc
}

func NewController(dep Dependency) *Controller {
	return &Controller{
		getValueUsecase: dep.GetValueUsecase,
		setValueUsecase: dep.SetValueUsecase,
	}
}

func (c *Controller) GetValue(ctx context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	result, err := c.getValueUsecase(ctx, &dbusecase.GetValueRequest{Key: request.Key})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return &dbgrpc.GetValueResponse{Value: result.Value}, nil
}

func (c *Controller) SetValue(ctx context.Context, request *dbgrpc.SetValueRequest) (*dbgrpc.SetValueResponse, error) {
	_, err := c.setValueUsecase(ctx, &dbusecase.SetValueRequest{Key: request.Key, Value: request.Value})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return nil, nil
}
