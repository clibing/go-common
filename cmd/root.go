/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	buildVersion string
	buildDate    string
	commitId     string
	execBinName  = "go-common"
)

func InitCmd() (err error) {
	app := &cli.App{
		Name:  execBinName,
		Usage: "input detail me",
		Authors: []*cli.Author{
			{
				Name:  "clibing",
				Email: "wmsjhappy@gmail.com",
			},
		},
		EnableBashCompletion: true,
		Copyright:            "Copyright (c) " + time.Now().Format("2006") + " clibing, All rights reserved.",
		Version:              fmt.Sprintf("%s/%s/%s", buildVersion, buildDate, commitId),
		Compiled:             time.Now(),
		Commands: []*cli.Command{
			install(),
			migration(),
			client(),
			server(),
		},
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println("cli run error: ", err)
	}
	return
}

func configFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "config",
		Usage: "load config file path",
		Value: "configs/config.yaml",
	}
}
