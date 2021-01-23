package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/0n1shi/mipsemu/pkg/mips"
	cli "github.com/urfave/cli/v2"
)

var (
	version = "unset"
	commit  = "unset"
	date    = "unset"
	builtBy = "unset"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "A file path of NES ROM",
			Required: true,
		},
	}

	app := cli.App{
		Name:    "Misper",
		Usage:   `A MIPS 1 CPU Emulator written in Golang`,
		Version: version,
		Action:  run,
		Flags:   flags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	romFile := c.String("file")

	data, err := ioutil.ReadFile(romFile)
	if err != nil {
		return err
	}

	emu, err := mips.NewEmulator(data)
	if err != nil {
		return err
	}

	emu.Run()

	return nil
}
