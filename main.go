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

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	log, err := utils.GetLog()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer log.Close()

	if err := godotenv.Load(config.CONFIG_FILE); err != nil {
		log.PrintfErr("%v\n", err)
		os.Exit(1)
	}

	if err := config.InitEpubDir(); err != nil {
		log.PrintfErr("%v\n", err)
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = config.NAME
	app.Usage = config.USAGE
	app.Version = config.VERSION
	app.Action = commands.MainAction
	app.Flags = commands.MainFlags
	app.Commands = []*cli.Command{
		&commands.ConvertCommand,
		&commands.ExistsCommand,
		&commands.IntervalCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.PrintfErr("%v\n", err)
	}

	_ = utils.RmTmpDir()
}
