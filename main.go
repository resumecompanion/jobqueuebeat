package main

import (
	"os"

	"github.com/resumecompanion/jobqueuebeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
