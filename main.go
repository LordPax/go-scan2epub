package main

import (
	"fmt"
	"os"
	"scan2epub/commands"
	"scan2epub/config"
	"scan2epub/lang"
	"scan2epub/utils"

	cli "github.com/urfave/cli/v2"
	ini "gopkg.in/ini.v1"
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

	config.CONFIG_INI, err = ini.Load(config.CONFIG_FILE)
	if err != nil {
		log.PrintfErr("%v\n", err)
		os.Exit(1)
	}

	if err := config.InitEpubDir(); err != nil {
		log.PrintfErr("%v\n", err)
		os.Exit(1)
	}

	l := lang.GetLocalize()
	l.SetLang(os.Getenv("LANG"))
	l.AddStrings(&lang.EN_STRINGS, "en_US.UTF-8", "en_GB.UTF-8", "en")
	l.AddStrings(&lang.FR_STRINGS, "fr_FR.UTF-8", "fr_CA.UTF-8", "fr")

	app := cli.NewApp()
	app.Name = config.NAME
	app.Usage = l.Get("usage")
	app.Version = config.VERSION
	app.Action = commands.MainAction
	app.Flags = commands.MainFlags()
	app.Commands = []*cli.Command{
		commands.ConvertCommand(),
		commands.ExistsCommand(),
		commands.IntervalCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.PrintfErr("%v\n", err)
	}

	_ = utils.RmTmpDir()
}
