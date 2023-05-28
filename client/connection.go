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

func (c *Connection) GetValue(ctx context.Context, key string) (string, error) {
	_, err := c.service.GetValue(ctx, &dbgrpc.GetValueRequest{Key: key})
	goutil.PanicUnhandledError(err)
	return "", nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}
