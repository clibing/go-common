package cmd

import (
	"github.com/urfave/cli/v2"
)

func client() *cli.Command {
	return &cli.Command{
		Name: "client",
		Flags: []cli.Flag{
			configFlag(),
		},
		Action: func(c *cli.Context) error {
			_ = cli.ShowAppHelp(c)
			return nil
		},
		Subcommands: []*cli.Command{},
	}
}
