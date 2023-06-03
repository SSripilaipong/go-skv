package clientrepl

import (
	"context"
	"go-skv/client"
	"go-skv/util/goutil"
	"strings"
)

func NewController(connectionFactory client.ConnectionFactory) *Controller {
	return &Controller{
		connectionFactory: connectionFactory,
	}
}

type Controller struct {
	connectionFactory client.ConnectionFactory
	connection        client.Connection
}

func (c *Controller) Connect(address string) (err error) {
	c.connection, err = c.connectionFactory(address)
	goutil.PanicUnhandledError(err)
	return nil
}

func (c *Controller) Input(s string) error {
	tokens := strings.Split(strings.Trim(s, "\n"), " ")
	key, err := goutil.ElementAt(tokens, 1)
	goutil.PanicUnhandledError(err)

	_, err = c.connection.GetValue(context.Background(), strings.Trim(key, "\""))
	goutil.PanicUnhandledError(err)

	return nil
}
