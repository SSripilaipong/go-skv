package clientcli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func New(dep Dependency) Interface {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands:             commands(dep),
	}
	return &clientCli{app: app}
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
