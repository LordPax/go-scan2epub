package commands

import (
	"fmt"
	"scan2epub/service"

	cli "github.com/urfave/cli/v2"
)

var IntervalCommand = cli.Command{
	Name:      "inter",
	Usage:     "Convert at regular intervals and increment the chapter number",
	ArgsUsage: "<cron> <chap>",
	Aliases:   []string{"i"},
	Action:    intervalAction,
}

func intervalAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("No chapter or cron specified")
	}

	cronStr := c.Args().Get(0)
	chap := c.Args().Get(1)

	if err := service.CronDownloadChap(cronStr, chap); err != nil {
		return err
	}

	return nil
}
