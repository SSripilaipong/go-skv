package clientrepltest

import (
	"context"
	"go-skv/client/clientconnection"
)

type ConnectionFactoryMock struct {
	Address string
	Return  clientconnection.Interface
}

func (f *ConnectionFactoryMock) New() clientconnection.ConnectionFactory {
	return func(address string) (clientconnection.Interface, error) {
		f.Address = address
		return f.Return, nil
	}
}

type ConnectionMock struct {
	GetValue_key   string
	GetValue_Value string
}

func (c *ConnectionMock) GetValue(_ context.Context, key string) (string, error) {
	c.GetValue_key = key
	return c.GetValue_Value, nil
}

func (c *ConnectionMock) SetValue(context.Context, string, string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ConnectionMock) Close() error {
	//TODO implement me
	panic("implement me")
}
