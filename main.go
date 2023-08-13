package main

import (
	"go-skv/cli"
	"go-skv/client"
	"go-skv/server"
)

func main() {
	cli.RunCli(
		server.NewCliCommands(),
		client.NewCliCommands(),
	)
}
