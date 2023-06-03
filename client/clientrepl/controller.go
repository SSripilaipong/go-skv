package clientrepl

import (
	"go-skv/client"
	"go-skv/util/goutil"
)

func NewController(connectionFactory client.ConnectionFactory) *Controller {
	return &Controller{
		connectionFactory: connectionFactory,
	}
}

type Controller struct {
	connectionFactory client.ConnectionFactory
}

func (c *Controller) Connect(address string) error {
	_, err := c.connectionFactory(address)
	goutil.PanicUnhandledError(err)
	return nil
}
