package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"

	"core"
)

func banner() {
}

var (
	enableAlgorithms cli.StringSlice
	filePath         cli.StringSlice
)

func initParam() {
}

func control() {
	// TODO 核心
	fmt.Println(enableAlgorithms.Value())
}

func main() {
	initParam()
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "core",
		Usage:                "calculate the core value of a file or string",
		Action: func(*cli.Context) error {
			banner()

			control()

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "string",
				Usage: "computes a string core",
				Action: func(cCtx *cli.Context) error {
					// TODO
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "enable",
				Aliases:     []string{"e"},
				Value:       cli.NewStringSlice("md5", "sha1", "sha256", "sha512"),
				Usage:       "enabled hashing algorithms",
				Destination: &enableAlgorithms,
			},
			&cli.StringSliceFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "file path",
				Destination: &filePath,
			},
			&cli.StringFlag{
				Name:    "string",
				Aliases: []string{"s"},
				Usage:   "computes a string core",
				Action: func(context *cli.Context, s string) error {
					core.StringHashReport(s, enableAlgorithms.Value())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
