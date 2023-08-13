package cli

import (
	"github.com/urfave/cli/v2"
)

func NewCommands(connectToServer func(string) error) []*cli.Command {
	return []*cli.Command{
		{
			Name: "connect",
			Action: func(ctx *cli.Context) error {
				serverIp := ctx.Args().First()
				return connectToServer(serverIp)
			},
		},
	}
}
