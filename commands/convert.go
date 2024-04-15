package commands

import (
	"fmt"
	"os"
	"scan2epub/service"
	"scan2epub/utils"

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

var convertFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "output directory",
		Action: func(c *cli.Context, value string) error {
			if value != "" {
				os.Unsetenv("EPUB_DIR")
				os.Setenv("EPUB_DIR", value)
			}
			return nil
		},
	},
}

func convertAction(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("no chapter specified")
	}

	fmt.Printf("Converting %d chapters\n", c.NArg())
	chaps := c.Args().Slice()
	if err := service.Scan2Epub(chaps); err != nil {
		return err
	}

	if err := utils.RmTmpDir(); err != nil {
		return err
	}

	return nil
}
