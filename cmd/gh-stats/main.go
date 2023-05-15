package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/liouk/gh-stats/pkg/cmd"
)

func main() {
	app := cmd.NewCLIApp()
	if app == nil {
		panic("could not create CLI app")
	}

	if err := app.Run(os.Args); err != nil {
		if strings.Contains(err.Error(), "Bad credentials") {
			fmt.Printf("could not login to GitHub: %v\n", err)
		} else {
			panic(err)
		}
	}
}
