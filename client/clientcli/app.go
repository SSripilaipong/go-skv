package clientcli

import (
	"github.com/urfave/cli/v2"
	"go-skv/client/clientrepl"
	"log"
	"os"
)

var app = buildApp()

func Run() {
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func buildApp() *cli.App {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands:             commands(),
	}
	return app
}

func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name: "connect",
			Action: func(ctx *cli.Context) error {
				serverIp := ctx.Args().First()
				return clientrepl.RunRuntimeRepl(serverIp)
			},
		},
	}
}
