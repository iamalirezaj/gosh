package flags

import "github.com/codegangsta/cli"

type RequiredFlag struct {
	cli.Flag
}

func (RequiredFlag) Required() bool {
	return true
}