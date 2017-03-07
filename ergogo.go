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

func prototypeFlashProcedure() {
	df, _ := dfu.NewDfuFlash("/Users/nicluo/Downloads/TestFirmware/left_kiibohd.dfu.bin", "/Users/nicluo/Downloads/TestFirmware/right_kiibohd.dfu.bin")
	for {
		dfuAvailable, _ := dfu.Scan()
		log.Printf("Scanning Left DFU: %t\n", dfuAvailable)
		if dfuAvailable {
			break
		}
	}
	log.Println("Flashing Left DFU...")
	df.FlashLeft()

	for {
		dfuAvailable, _ := dfu.Scan()
		log.Printf("Scanning Right DFU: %t\n", dfuAvailable)
		if dfuAvailable {
			break
		}
	}
	log.Println("Flashing Right DFU...")
	df.FlashRight()
}

func main() {
	startupFirmwareWatcher()
	startupDfu()
	prototypeFlashProcedure()

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
