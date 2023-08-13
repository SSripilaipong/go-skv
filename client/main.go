package client

import (
	"github.com/urfave/cli/v2"
	clientCli "go-skv/client/cli"
	"go-skv/client/clientconnection"
	"go-skv/client/clientrepl"
)

func NewCliCommands() []*cli.Command {
	return clientCli.NewCommands(clientrepl.NewReplRunner(
		clientconnection.New,
	))
}
