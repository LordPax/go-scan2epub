package main

import (
	"fmt"
	"os"
	"scan2epub/commands"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "CLI tool to convert scan to epub"
	app.Version = "0.0.1"
	app.Commands = []*cli.Command{
		&commands.ConvertCommand,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
