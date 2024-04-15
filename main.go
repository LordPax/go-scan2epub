package main

import (
	"fmt"
	"os"
	"scan2epub/commands"
	"scan2epub/config"
	"scan2epub/utils"

	"github.com/joho/godotenv"
	cli "github.com/urfave/cli/v2"
)

func mainAction(c *cli.Context) error {
	return fmt.Errorf("no command specified")
}

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := godotenv.Load(config.CONFIG_FILE); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := config.InitEpubDir(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = config.NAME
	app.Usage = config.USAGE
	app.Version = config.VERSION
	app.Action = mainAction
	app.Commands = []*cli.Command{
		&commands.ConvertCommand,
		&commands.ExistsCommand,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		utils.RmTmpDir()
		os.Exit(1)
	}
}
