package commands

import (
	"fmt"
	"scan2epub/lang"
	"scan2epub/service"
	"scan2epub/utils"

	cli "github.com/urfave/cli/v2"
)

func ConvertCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:      "conv",
		Usage:     l.Get("convert-usage"),
		ArgsUsage: "<chap> [chap...]",
		Aliases:   []string{"c"},
		Action:    convertAction,
	}
}

func convertAction(c *cli.Context) error {
	l := lang.GetLocalize()
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return fmt.Errorf(l.Get("no-chapter"))
	}

	log.Printf(l.Get("converting"), c.NArg())
	chaps := c.Args().Slice()
	if err := service.Scan2Epub(chaps); err != nil {
		return err
	}

	return nil
}
