package servercli

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
			Name: "start",
			Action: func(*cli.Context) error {
				return dep.Start()
			},
		},
	}
}
