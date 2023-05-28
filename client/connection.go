package client

import (
	"context"
	"go-skv/server/dbserver/dbgrpc"
	"go-skv/util/goutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnection(address string) *Connection {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	goutil.PanicUnhandledError(err)

	service := dbgrpc.NewDbServiceClient(conn)
	return &Connection{service: service, conn: conn}
}

type Connection struct {
	service dbgrpc.DbServiceClient
	conn    *grpc.ClientConn
}

func (c *Connection) GetValue(_ context.Context, key string) (string, error) {
	response, err := c.service.GetValue(context.Background(), &dbgrpc.GetValueRequest{Key: key})
	goutil.PanicUnhandledError(err)
	return *goutil.Coalesce(response.Value, goutil.Pointer("")), nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}
