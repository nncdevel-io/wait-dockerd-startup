package main

import (
	"fmt"
	"os"
)

func main() {
	options := OptionsFromEnv()

	cli, err := NewClientProvider(options.RequestTimeout).Provide()
	if err != nil {
		panic(err)
	}

	checker := NewChecker(*cli)
	poller := NewPoller(checker, options)

	if poller.Poll() != nil {
		fmt.Printf("Error occcured. %s\n", err)
		os.Exit(1)
	}

}
