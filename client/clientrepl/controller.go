package clientrepl

import (
	"context"
	"fmt"
	"go-skv/client/clientconnection"
	"go-skv/util/goutil"
	"strings"
)

func NewController(connectionFactory clientconnection.ConnectionFactory) *Controller {
	return &Controller{
		connectionFactory: connectionFactory,
	}
}

type Controller struct {
	connectionFactory clientconnection.ConnectionFactory
	connection        clientconnection.Interface
}

func (c *Controller) Connect(address string) (err error) {
	c.connection, err = c.connectionFactory(address)
	goutil.PanicUnhandledError(err)
	return nil
}

func (c *Controller) Input(s string) (string, error) {
	tokens := strings.Split(strings.Trim(s, "\n"), " ")
	keyWithQuotes, err := goutil.ElementAt(tokens, 1)
	goutil.PanicUnhandledError(err)

	key := strings.Trim(keyWithQuotes, "\"")

	value, err := c.connection.GetValue(context.Background(), key)
	goutil.PanicUnhandledError(err)

	return fmt.Sprintf("%#v => %#v\n", key, value), nil
}
