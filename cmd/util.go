package cmd

import (
	"fmt"
	"os"
)

func failIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %s\n", err)
		os.Exit(1)
	}
}

func fatal(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("FATAL: %s\n", msg), args...)
	os.Exit(1)
}
