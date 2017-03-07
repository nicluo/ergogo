package cmd

import "github.com/urfave/cli"

// Reset command empties configuration folders to leave a clean slate
var Reset = cli.Command{
	Name:   "reset",
	Usage:  "Clear configuration and past firmware downloads",
	Action: resetAction,
}

func resetAction(ctx *cli.Context) error {
	return nil
}
