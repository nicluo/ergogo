package main

import (
	"log"
	"os"

	"github.com/nicluo/ergogo/cmd"
	"github.com/nicluo/ergogo/modules/dfu"
	"github.com/nicluo/ergogo/modules/firmware"
	"github.com/urfave/cli"
)

// App Name and Version
const (
	AppName = "Ergogo"
	AppVer  = "0.0.1"
)

func startupFirmwareWatcher() {
	fw, err := firmware.NewFirmwareWatcher("/Users/nicluo/Downloads")
	if err != nil {
		log.Fatalf("Error starting firmware watcher: %v", err)
	}
	fw.Close()
}

func startupDfu() {
	_, err := dfu.Scan()
	if err != nil {
		log.Fatalf("Error running DFU command: %v", err)
	}
}

func main() {
	startupFirmwareWatcher()
	startupDfu()

	app := cli.NewApp()
	app.Name = AppName
	app.Usage = "Interactively manage Ergodox firmware"
	app.Version = AppVer
	app.Commands = []cli.Command{
		cmd.CmdRun,
		cmd.Setup,
		cmd.Reset,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
