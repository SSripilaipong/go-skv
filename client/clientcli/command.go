package clientcli

import (
	"github.com/urfave/cli/v2"
)

func commands(dep Dependency) []*cli.Command {
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
