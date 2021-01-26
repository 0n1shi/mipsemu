package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/0n1shi/mipsemu/pkg/mips"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

var (
	version = "unset"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "A file path of MIPS binary",
			Required: true,
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Running in debug mode",
		},
	}

	app := cli.App{
		Name:    "Misper",
		Usage:   `A MIPS CPU Emulator written in Golang`,
		Version: version,
		Action:  run,
		Flags:   flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	file := c.String("file")
	debug := c.Bool("debug")

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.WithStack(err)
	}

	emu, err := mips.NewEmulator(data, debug)
	if err != nil {
		return errors.WithStack(err)
	}

	return emu.Run()
}
