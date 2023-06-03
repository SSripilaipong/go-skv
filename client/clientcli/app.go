package clientcli

import (
	"github.com/urfave/cli/v2"
)

func newCliApp(dep dependency) *cli.App {
	return &cli.App{
		EnableBashCompletion: true,
		Commands:             commands(dep),
	}
}

func commands(dep dependency) []*cli.Command {
	return []*cli.Command{
		{
			Name: "connect",
			Action: func(ctx *cli.Context) error {
				serverIp := ctx.Args().First()
				return dep.ConnectToServer(serverIp)
			},
		},
	}
}
