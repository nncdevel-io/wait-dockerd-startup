package main

import (
	"fmt"
	"os/exec"
	"time"
)

type Poller interface {
	Poll() error
}

type DefaultPoller struct {
	Interval time.Duration
	Timeout  time.Duration
	Checker  Checker
}

func NewPoller(checker Checker, options *Options) Poller {
	return DefaultPoller{
		Interval: options.Interval,
		Timeout:  options.Timeout,
		Checker:  checker,
	}
}

func (p DefaultPoller) Poll() error {
	elapsed := time.Duration(0)

	for {
		fmt.Printf("[%s] Trying to retrieve dockerd info. ", elapsed)

		elapsed = elapsed + p.Interval
		last := elapsed > p.Timeout

		version, err := p.Checker.Check()

		if err != nil {
			if last {
				return &exec.Error{
					Name: "Timeout occurred.",
					Err:  err,
				}
			} else {
				fmt.Printf("Waiting for dockerd startup. %s\n", err)
				time.Sleep(p.Interval)
			}
		} else {
			fmt.Printf("Server (%s) started.\n", version)
			fmt.Println("Docker daemon is Ready.")
			return nil
		}
	}
}
