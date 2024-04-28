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

var mainFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "output directory",
		Action: func(c *cli.Context, value string) error {
			if value == "" {
				return fmt.Errorf("output directory is empty")
			}

			os.Unsetenv("EPUB_DIR")
			os.Setenv("EPUB_DIR", value)

			return nil
		},
	},
	&cli.BoolFlag{
		Name:    "silent",
		Aliases: []string{"s"},
		Usage:   "disable printing log to stdout",
		Action: func(c *cli.Context, value bool) error {
			log, err := utils.GetLog()
			if err != nil {
				return err
			}

			log.SetSilent(value)

			return nil
		},
	},
}

func main() {
	os.Setenv("SILENT", "false")

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
	app.Action = mainAction
	app.Flags = mainFlags
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
