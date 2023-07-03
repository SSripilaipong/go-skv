package replcontroller

import (
	"go-skv/client/clientconnection"
	goutil2 "go-skv/common/util/goutil"
	"strings"
)

func New(connectionFactory clientconnection.ConnectionFactory) Interface {
	ctrl := &controller{
		connectionFactory: connectionFactory,
	}
	ctrl.generateCommandMapper()
	return ctrl
}

type controller struct {
	connectionFactory clientconnection.ConnectionFactory
	connection        clientconnection.Interface
	commandMapper     map[string]func([]string) (string, error)
}

func (c *controller) Connect(address string) (err error) {
	c.connection, err = c.connectionFactory(address)
	goutil2.PanicUnhandledError(err)
	return nil
}

func (c *controller) Input(s string) (string, error) {
	tokens := strings.Split(strings.Trim(s, "\n"), " ")
	command, err := goutil2.ElementAt(tokens, 0)
	goutil2.PanicUnhandledError(err)

	params := tokens[1:]
	handle, matches := c.commandMapper[strings.ToLower(command)]
	if !matches {
		return "", nil
	}

	return handle(params)
}
