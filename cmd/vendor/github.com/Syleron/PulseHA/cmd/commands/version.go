package commands

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/cli"
)

type VersionCommand struct {
	Version        string
	VersionRelease string
	Build		   string
	Ui             cli.Ui
}

/**
 *
 */
func (c *VersionCommand) Help() string {
	return ""
}

/**
 *
 */
func (c *VersionCommand) Run(_ []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "pulseha v%s", c.Version)

	c.Ui.Output(versionString.String() + " Copyright (c) 2017 Andrew Zak <andrew@pulseha.com>")

	return 0
}

/**
 *
 */
func (c *VersionCommand) Synopsis() string {
	return "Prints the PulseHA version"
}
