package client

import (
	"go-skv/client/clientcli"
	"go-skv/client/clientconnection"
	"go-skv/client/clientrepl"
)

func RunCli() {
	cli := clientcli.New(
		clientrepl.NewReplRunner(
			clientconnection.New,
		),
	)
	cli.Run()
}
