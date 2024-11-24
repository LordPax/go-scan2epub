package commands

import (
	"scan2epub/lang"
	"scan2epub/service"

	cli "github.com/urfave/cli/v2"
)

func IntervalCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:    "inter",
		Usage:   l.Get("interval-usage"),
		Aliases: []string{"i"},
		Action:  intervalAction,
	}
}

func intervalAction(c *cli.Context) error {
	if err := service.CronDownloadChap(); err != nil {
		return err
	}

	return nil
}
