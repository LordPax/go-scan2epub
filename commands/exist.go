package commands

import (
	"fmt"
	"scan2epub/service"

	cli "github.com/urfave/cli/v2"
)

var ExistsCommand = cli.Command{
	Name:      "exist",
	Usage:     "Check if chapter exists",
	ArgsUsage: "<chap>",
	Aliases:   []string{"e"},
	Flags:     existFlags,
	Action:    existAction,
}

var existFlags = []cli.Flag{}

func existAction(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("No chapter specified")
	}

	chaps := c.Args().Slice()[0]
	if service.CheckChapExist(chaps) {
		fmt.Printf("Chapter %s exists\n", chaps)
	} else {
		fmt.Printf("Chapter %s not exists\n", chaps)
	}

	return nil
}
