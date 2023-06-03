package replcontrollertest

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
	SetValue_key   string
	SetValue_value string
	Close_IsCalled bool
}

func (c *ConnectionMock) GetValue(_ context.Context, key string) (string, error) {
	c.GetValue_key = key
	return c.GetValue_Value, nil
}

func (c *ConnectionMock) SetValue(_ context.Context, key string, value string) error {
	c.SetValue_key = key
	c.SetValue_value = value
	return nil
}

func (c *ConnectionMock) Close() error {
	c.Close_IsCalled = true
	return nil
}
