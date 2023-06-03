package clientrepl

import (
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
	command, err := goutil.ElementAt(tokens, 0)
	goutil.PanicUnhandledError(err)

	switch strings.ToLower(command) {
	case "getvalue":
		return c.handleGetValueCommand(tokens[1:])
	case "setvalue":
		return c.handleSetValueCommand(tokens[1:])
	}
	return "", nil
}
