package clientconnection

import (
	"context"
	goutil2 "go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectionFactory func(string) (Interface, error)

func New(address string) (Interface, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil2.PanicUnhandledError(err)

	service := dbgrpc.NewDbServiceClient(conn)
	return &connection{service: service, conn: conn}, nil
}

type connection struct {
	service dbgrpc.DbServiceClient
	conn    *grpc.ClientConn
}

func (c *connection) GetValue(ctx context.Context, key string) (string, error) {
	response, grpcErr := c.service.GetValue(ctx, &dbgrpc.GetValueRequest{Key: key})

	clientErr, err := parseGrpcError(grpcErr)
	goutil2.PanicUnhandledError(err)
	if clientErr != nil {
		return "", clientErr
	}

	return *goutil2.Coalesce(response.Value, goutil2.Pointer("")), nil
}

func (c *connection) SetValue(ctx context.Context, key string, value string) error {
	_, grpcErr := c.service.SetValue(ctx, &dbgrpc.SetValueRequest{Key: key, Value: value})

	clientErr, err := parseGrpcError(grpcErr)
	goutil2.PanicUnhandledError(err)
	if clientErr != nil {
		return clientErr
	}

	return nil
}

func (c *connection) Close() error {
	return c.conn.Close()
}
