package commands

import (
	"fmt"
	"scan2epub/lang"
	"scan2epub/service"

	cli "github.com/urfave/cli/v2"
)

func IntervalCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:      "inter",
		Usage:     l.Get("interval-usage"),
		ArgsUsage: "<cron> <chap>",
		Aliases:   []string{"i"},
		Action:    intervalAction,
	}
}

func intervalAction(c *cli.Context) error {
	l := lang.GetLocalize()
	if c.NArg() < 2 {
		return fmt.Errorf(l.Get("no-chapter-cron"))
	}

	cronStr := c.Args().Get(0)
	chap := c.Args().Get(1)

	if err := service.CronDownloadChap(cronStr, chap); err != nil {
		return err
	}

	return nil
}
