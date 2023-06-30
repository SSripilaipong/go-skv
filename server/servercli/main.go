package servercli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func New(start func(Config) error) Interface {
	return &serverCli{app: newCliApp(dependency{
		Start: start,
	})}
}

type serverCli struct {
	app *cli.App
}

func (c *serverCli) Run() {
	err := c.app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
