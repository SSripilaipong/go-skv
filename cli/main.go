package cli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func RunCli(serverCommands, clientCommands []*cli.Command) {
	if err := (&cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "server",
				Subcommands: serverCommands,
			},
			{
				Name:        "client",
				Subcommands: clientCommands,
			},
		},
	}).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
