package dbserver

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/server/dbusecase"
)

type controller struct {
	dbgrpc.UnimplementedDbServiceServer
	getValueUsecase dbusecase.GetValueFunc
}

func (c *controller) GetValue(_ context.Context, request *dbgrpc.GetValueRequest) (*dbgrpc.GetValueResponse, error) {
	_, _ = c.getValueUsecase(&dbusecase.GetValueRequest{Key: request.Key})
	return nil, nil
}
