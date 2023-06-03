package clientconnection

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectionFactory func(string) (Interface, error)

func New(address string) (Interface, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil.PanicUnhandledError(err)

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
	goutil.PanicUnhandledError(err)
	if clientErr != nil {
		return "", clientErr
	}

	return *goutil.Coalesce(response.Value, goutil.Pointer("")), nil
}

func (c *connection) SetValue(ctx context.Context, key string, value string) error {
	_, grpcErr := c.service.SetValue(ctx, &dbgrpc.SetValueRequest{Key: key, Value: value})

	clientErr, err := parseGrpcError(grpcErr)
	goutil.PanicUnhandledError(err)
	if clientErr != nil {
		return clientErr
	}

	return nil
}

func (c *connection) Close() error {
	return c.conn.Close()
}
