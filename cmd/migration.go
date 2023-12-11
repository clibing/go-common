package cmd

import "github.com/urfave/cli/v2"

func migration() *cli.Command {

	return &cli.Command{
		Name:      "migrate",
		Usage:     "alter table, insert update record.",
		UsageText: execBinName + " migrate [option1|option2...]",
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
