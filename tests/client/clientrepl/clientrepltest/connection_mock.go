package clientrepltest

import (
	"context"
	"go-skv/client"
)

type ConnectionFactoryMock struct {
	Address string
	Return  client.Connection
}

func (f *ConnectionFactoryMock) New() client.ConnectionFactory {
	return func(address string) (client.Connection, error) {
		f.Address = address
		return f.Return, nil
	}
}

type ConnectionMock struct {
	GetValue_key string
}

func (c *ConnectionMock) GetValue(_ context.Context, key string) (string, error) {
	c.GetValue_key = key
	return "", nil
}

func (c *ConnectionMock) SetValue(context.Context, string, string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ConnectionMock) Close() error {
	//TODO implement me
	panic("implement me")
}
