package commands

import (
	"fmt"
	"scan2epub/utils"

	cli "github.com/urfave/cli/v2"
)

var IntervalCommand = cli.Command{
	Name:      "inter",
	Usage:     "Convert at regular intervals and increment the chapter number",
	ArgsUsage: "<interval> <chap>",
	Aliases:   []string{"i"},
	Action:    intervalAction,
}

func intervalAction(c *cli.Context) error {
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	if c.NArg() < 2 {
		return fmt.Errorf("no chapter specified")
	}

	log.Printf("Converting %d chapters\n", c.NArg())

	if err := utils.RmTmpDir(); err != nil {
		return err
	}

	return nil
}
