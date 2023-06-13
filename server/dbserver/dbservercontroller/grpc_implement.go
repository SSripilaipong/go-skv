package dbservercontroller

import (
	"context"
	"fmt"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/util/goutil"
)

type grpcImplementation struct {
	dbgrpc.UnimplementedDbServiceServer
	usecase dbusecase.Interface
}

func newGrpcImplementation(usecase dbusecase.Interface) dbgrpc.DbServiceServer {
	return &grpcImplementation{
		usecase: usecase,
	}
}

func (c *grpcImplementation) GetValue(ctx context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	result, err := c.usecase.GetValue(ctx, dbusecase.GetValueRequest{Key: request.Key})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return &dbgrpc.GetValueResponse{Value: goutil.Pointer(result.Value)}, nil
}

func (c *grpcImplementation) SetValue(ctx context.Context, request *dbgrpc.SetValueRequest) (*dbgrpc.SetValueResponse, error) {
	_, err := c.usecase.SetValue(ctx, dbusecase.SetValueRequest{Key: request.Key, Value: request.Value})
	if err != nil {
		panic(fmt.Errorf("unhandled error: %f", err))
	}
	return &dbgrpc.SetValueResponse{}, nil
}
