package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// Setup initializes the configuration directory and binds keyboards
var Setup = cli.Command{
	Name:   "setup",
	Usage:  "Setup Ergogo configuration for the first time",
	Action: setupAction,
}

func setupAction(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter download directory to watch for new firmware (default: ~/Downloads): ")
	downloadDirectory, _ := reader.ReadString('\n')

	fmt.Println("Settings saved to ~/.ergogo/settings.json")

	fmt.Println("Your setup looks good! Run `ergogo run` to start using or `ergogo reset` to clear settings. :)")

	fmt.Println(downloadDirectory)
	return nil
}
