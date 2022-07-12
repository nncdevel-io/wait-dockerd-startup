package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/docker/docker/client"
	"os"
	"time"
)

func main() {
	intervalFlag := flag.Int("interval", 1, "check interval.")
	timeoutFlag := flag.Int("timeout", 60, "start up timeout.")
	requestTimeoutFlag := flag.Int("request-timeout", 3, "docker api request timeout.")

	elapsed := time.Duration(0)

	flag.Parse()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithTimeout(time.Duration(*requestTimeoutFlag)*time.Second))
	if err != nil {
		panic(err)
	}

	interval := time.Duration(*intervalFlag) * time.Second
	timeout := time.Duration(*timeoutFlag) * time.Second

	for {
		fmt.Printf("[%s] Trying to retrieve dockerd info. ", elapsed)

		elapsed = elapsed + interval
		last := elapsed > timeout

		info, err := cli.Info(context.Background())

		if err != nil {
			if last {
				fmt.Printf("Timeout. last Error: %s\n", err)
				os.Exit(1)
			} else {
				fmt.Printf("Waiting for dockerd startup. %s\n", err)
				time.Sleep(interval)
			}
		} else {
			fmt.Printf("Server (%s) started.\n", info.ServerVersion)
			fmt.Println("Docker daemon is Ready.")
			break
		}

	}
}
