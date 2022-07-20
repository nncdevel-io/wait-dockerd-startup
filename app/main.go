package app

import (
	"fmt"
	"time"
)

type Options struct {
	InitialDelay     time.Duration
	Period           time.Duration
	Timeout          time.Duration
	SuccessThreshold int
	FailureThreshold int
}

func Main(options *Options) error {

	cli, err := NewClientProvider(options.Timeout).Provide()
	if err != nil {
		return err
	}

	checker := NewChecker(cli)

	poller := NewPoller(checker, &PollingOptions{
		Period:           options.Period,
		SuccessThreshold: options.SuccessThreshold,
		FailureThreshold: options.FailureThreshold,
	})

	fmt.Printf("wait %s.\n", options.InitialDelay)

	time.Sleep(options.InitialDelay)

	if poller.Poll() != nil {
		return err
	}

	return nil

}
