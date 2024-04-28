package commands

import (
	"fmt"
	"scan2epub/service"
	"scan2epub/utils"

	cli "github.com/urfave/cli/v2"
)

var ConvertCommand = cli.Command{
	Name:      "conv",
	Usage:     "Convert scan to epub",
	ArgsUsage: "<chap> [chap...]",
	Aliases:   []string{"c"},
	Action:    convertAction,
}

func convertAction(c *cli.Context) error {
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return fmt.Errorf("no chapter specified")
	}

	log.Printf("Converting %d chapters\n", c.NArg())
	chaps := c.Args().Slice()
	if err := service.Scan2Epub(chaps); err != nil {
		return err
	}

	if err := utils.RmTmpDir(); err != nil {
		return err
	}

	return nil
}
