package clientcli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func New(connectToServer func(string) error) Interface {
	return &clientCli{
		app: newCliApp(dependency{
			ConnectToServer: connectToServer,
		}),
	}
}

type clientCli struct {
	app *cli.App
}

func (c *clientCli) Run() {
	err := c.app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
