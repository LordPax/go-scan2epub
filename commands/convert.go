package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var ConvertCommand = cli.Command{
	Name:      "conv",
	Usage:     "Convert scan to epub",
	ArgsUsage: "<chap> [chap...]",
	Aliases:   []string{"c"},
	Flags:     convertFlags,
	Action:    convertAction,
}

var convertFlags = []cli.Flag{}

func convertAction(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("no chapter specified")
	}

	fmt.Printf("Converting %d chapters\n", c.NArg())
	chaps := c.Args().Slice()

	fmt.Printf("Chapters: %v\n", chaps)

	return nil
}
