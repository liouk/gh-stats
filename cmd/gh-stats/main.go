package main

import (
	"os"

	"github.com/liouk/gh-stats/pkg/cmd"
)

func main() {
	app := cmd.NewCLIApp()
	if app == nil {
		panic("could not create CLI app")
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
