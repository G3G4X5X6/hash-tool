package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"main/core"
	"os"
	"time"
)

//
//func banner() {
//	banner := `
// _     ____  ____  _         _____  ____  ____  _
/// \ /|/  _ \/ ___\/ \ /|    /__ __\/  _ \/  _ \/ \
//| |_||| / \||    \| |_||_____ / \  | / \|| / \|| |
//| | ||| |-||\___ || | ||\____\| |  | \_/|| \_/|| |_/\
//\_/ \|\_/ \|\____/\_/ \|      \_/  \____/\____/\____/
//
//`
//	fmt.Print(banner)
//}

var (
	enableAlgorithms cli.StringSlice
	threads          int
)

func initParam() {
}

func main() {
	startTime := time.Now()

	initParam()
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "hash",
		Usage:                "calculate the core value of a file or string",
		Action: func(*cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "string",
				Usage: "computes a string core",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("--string value, -s value                               computes a string core")
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "enable",
				Aliases:     []string{"e"},
				Value:       cli.NewStringSlice("crc32", "md5", "sha1", "sha256", "sha512"),
				Usage:       "enabled hashing algorithms",
				Destination: &enableAlgorithms,
			},
			&cli.StringSliceFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "file path",
				Action: func(context *cli.Context, strings []string) error {
					core.FilesHashReport(strings, enableAlgorithms.Value(), threads)
					return nil
				},
			},
			&cli.IntFlag{
				Name:        "thread",
				Value:       6,
				Aliases:     []string{"t"},
				Usage:       "number of threads",
				Destination: &threads,
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

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("\nProgram runtime: %s", elapsedTime)
}
