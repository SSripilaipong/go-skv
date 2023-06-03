package client

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectionFactory func(string) (*Connection, error)

func NewConnection(address string) (*Connection, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil.PanicUnhandledError(err)

	service := dbgrpc.NewDbServiceClient(conn)
	return &Connection{service: service, conn: conn}, nil
}

type Connection struct {
	service dbgrpc.DbServiceClient
	conn    *grpc.ClientConn
}

func (c *Connection) GetValue(ctx context.Context, key string) (string, error) {
	response, grpcErr := c.service.GetValue(ctx, &dbgrpc.GetValueRequest{Key: key})

	clientErr, err := parseGrpcError(grpcErr)
	goutil.PanicUnhandledError(err)
	if clientErr != nil {
		return "", clientErr
	}

	return *goutil.Coalesce(response.Value, goutil.Pointer("")), nil
}

func (c *Connection) SetValue(ctx context.Context, key string, value string) error {
	_, grpcErr := c.service.SetValue(ctx, &dbgrpc.SetValueRequest{Key: key, Value: value})

	clientErr, err := parseGrpcError(grpcErr)
	goutil.PanicUnhandledError(err)
	if clientErr != nil {
		return clientErr
	}

	return nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}
