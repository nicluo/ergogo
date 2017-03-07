package cmd

import "github.com/urfave/cli"

// CmdRun is a placeholder default command and runs the GUI
var CmdRun = cli.Command{
	Name:   "run",
	Usage:  "Starts Ergogo main application (default)",
	Action: runApp,
}

func runApp(ctx *cli.Context) error {
	return nil
}
