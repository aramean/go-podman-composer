package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DescriptionVersionCommand = "Show Podman-Compose version information"
)

type VersionCommand struct {
	fs *flag.FlagSet

	format string
	short  bool
}

func NewVersionCommand() *VersionCommand {
	gc := &VersionCommand{
		fs: flag.NewFlagSet("version", flag.ExitOnError),
	}

	gc.fs.StringVar(&gc.format, "f", "", "")
	gc.fs.BoolVar(&gc.short, "short", false, "Shows only Compose's version number.")
	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionVersionCommand, "version")
	return gc
}

func (g *VersionCommand) Name() string {
	return g.fs.Name()
}

func (g *VersionCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *VersionCommand) Run() error {
	var long = binaryName + " version "
	if g.short {
		long = ""
	}
	return fmt.Errorf(long + "v" + binaryVersion)
}
