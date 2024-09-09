package commands

import (
	"fmt"
	"scan2epub/config"
	"scan2epub/lang"
	"scan2epub/utils"

	cli "github.com/urfave/cli/v2"
)

func MainFlags() []cli.Flag {
	l := lang.GetLocalize()
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   l.Get("output-desc"),
			Action: func(c *cli.Context, value string) error {
				if value == "" {
					return fmt.Errorf(l.Get("output-dir-empty"))
				}

				defaultSource := config.CONFIG_INI.Section("").Key("default").String()
				config.CONFIG_INI.Section(defaultSource).Key("epub_dir").SetValue(value)

				return nil
			},
		},
		&cli.StringFlag{
			Name:    "source",
			Aliases: []string{"S"},
			Usage:   l.Get("source-desc"),
			Action: func(c *cli.Context, value string) error {
				if value == "" {
					return fmt.Errorf(l.Get("source-empty"))
				}

				config.CONFIG_INI.Section("").Key("default").SetValue(value)

				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "silent",
			Aliases: []string{"s"},
			Usage:   l.Get("silent"),
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
}

func MainAction(c *cli.Context) error {
	l := lang.GetLocalize()
	return fmt.Errorf(l.Get("no-command"))
}
