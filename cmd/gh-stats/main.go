package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/liouk/gh-stats/cmd/gh-stats/app"
)

func main() {
	err := app.NewCLIApp().Run(os.Args)

	if err != nil {
		if strings.Contains(err.Error(), "Bad credentials") {
			fmt.Printf("could not login to GitHub: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
