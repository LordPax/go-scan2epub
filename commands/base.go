package commands

import (
	"fmt"
	"os"
	"scan2epub/utils"

	cli "github.com/urfave/cli/v2"
)

var MainFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "output directory",
		Action: func(c *cli.Context, value string) error {
			if value == "" {
				return fmt.Errorf("Output directory is empty")
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

func MainAction(c *cli.Context) error {
	return fmt.Errorf("No command specified")
}
