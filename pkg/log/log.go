package log

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type stdoutLogger struct {
	verbose bool
}

var logger *stdoutLogger

func Init(clictx *cli.Context) {
	if logger == nil {
		logger = &stdoutLogger{
			clictx.Bool("verbose"),
		}
	}
}

func Logf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Logvf(format string, a ...any) {
	if !logger.verbose {
		return
	}

	fmt.Printf(format, a...)
}

func Verbose() bool {
	return logger.verbose
}
