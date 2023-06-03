package replcontroller

import (
	"go-skv/client/clientconnection"
	"go-skv/util/goutil"
	"strings"
)

func NewController(connectionFactory clientconnection.ConnectionFactory) *Controller {
	ctrl := &Controller{
		connectionFactory: connectionFactory,
	}
	ctrl.generateCommandMapper()
	return ctrl
}

type Controller struct {
	connectionFactory clientconnection.ConnectionFactory
	connection        clientconnection.Interface
	commandMapper     map[string]func([]string) (string, error)
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

	params := tokens[1:]
	handle, matches := c.commandMapper[strings.ToLower(command)]
	if !matches {
		return "", nil
	}

	return handle(params)
}
