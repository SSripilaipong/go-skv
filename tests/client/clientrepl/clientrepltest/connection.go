package clientrepltest

import (
	"go-skv/client"
)

type ConnectionFactoryMock struct {
	Address string
}

func (f *ConnectionFactoryMock) New() client.ConnectionFactory {
	return func(address string) (*client.Connection, error) {
		f.Address = address
		return nil, nil
	}
}
