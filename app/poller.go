//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package app

import (
	"fmt"
	"time"
)

type Poller interface {
	Poll() error
}

type DefaultPoller struct {
	Period           time.Duration
	SuccessThreshold int
	FailureThreshold int
	Checker          Checker
}

type PollingOptions struct {
	Period           time.Duration
	SuccessThreshold int
	FailureThreshold int
}

func NewPoller(checker Checker, options *PollingOptions) Poller {
	return DefaultPoller{
		Period:           options.Period,
		SuccessThreshold: options.SuccessThreshold,
		FailureThreshold: options.FailureThreshold,
		Checker:          checker,
	}
}

func (p DefaultPoller) Poll() error {
	result := NewResult(p.SuccessThreshold, p.FailureThreshold)
	elapsed := time.Duration(0)

	for {
		fmt.Printf("[%s] Trying to retrieve dockerd info.", elapsed)

		serverVersion, err := p.Checker.Check()

		if err != nil {
			fmt.Printf("Docker Response Error. %s\n", err)
			result = result.IncrementFailure()
		} else {
			fmt.Printf("Server (%s) started.\n", serverVersion)
			result = result.IncrementSuccess()
		}

		switch result.State {
		case Continue:
			fmt.Printf("[%s] Continue. %s\n", elapsed, result.Format())
			time.Sleep(p.Period)
			elapsed = elapsed + p.Period
			continue
		case Failed:
			fmt.Printf("[%s] Failued. %s\n", elapsed, result.Format())
			return err
		case Succeeded:
			fmt.Printf("[%s] Succeeded. %s\n", elapsed, result.Format())
			return nil
		}
	}
}
