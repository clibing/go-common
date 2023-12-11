package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func install() *cli.Command {

	return &cli.Command{
		Name:      "install",
		Flags: []cli.Flag{
			configFlag(),
			&cli.StringFlag{
				Name:  "path",
				Usage: "install path, default /usr/local/bin",
				Value: "/usr/local/bin",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "exec bin name, default discovery",
				Value: "discovery",
			},
		},
		Action: func(c *cli.Context) error {

			binPath, err := exec.LookPath(os.Args[0])
			if err != nil {
				fmt.Printf("failed to get bin file info: %s: %s", os.Args[0], err)
				return err
			}

			currentFile, err := os.Open(binPath)
			if err != nil {
				fmt.Printf("failed to get bin file info: %s: %s", binPath, err)
				return err
			}
			defer func() { _ = currentFile.Close() }()

			path := c.String("path")
			name := c.String("name")

			installFile, err := os.OpenFile(filepath.Join(path, name), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
			if err != nil {
				fmt.Printf("failed to create bin file: %s: %s", filepath.Join(path, name), err)
				return err
			}
			defer func() { _ = installFile.Close() }()

			_, err = io.Copy(installFile, currentFile)
			if err != nil {
				fmt.Printf("failed to copy file: %s: %s", filepath.Join(path, name), err)
				return err
			}
			fmt.Println("install success:", filepath.Join(path, name))

			return nil
		},
		Subcommands: []*cli.Command{},
	}
}
