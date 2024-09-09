package commands

import (
	"fmt"
	"scan2epub/lang"
	"scan2epub/service"

	cli "github.com/urfave/cli/v2"
)

func ExistsCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:      "exist",
		Usage:     l.Get("exist-usage"),
		ArgsUsage: "<chap>",
		Aliases:   []string{"e"},
		Flags:     existFlags(),
		Action:    existAction,
	}
}

func existFlags() []cli.Flag {
	return []cli.Flag{}
}

func existAction(c *cli.Context) error {
	l := lang.GetLocalize()
	if c.NArg() == 0 {
		return fmt.Errorf(l.Get("no-chapter"))
	}

	chaps := c.Args().Slice()[0]
	if service.CheckChapExist(chaps) {
		fmt.Printf(l.Get("chapter-exists"), chaps)
	} else {
		fmt.Printf(l.Get("chapter-not-exists"), chaps)
	}

	return nil
}
