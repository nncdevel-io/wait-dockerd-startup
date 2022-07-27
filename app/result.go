//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package app

import "fmt"

type Result interface {
	IncrementSuccess() *Result
	IncrementFailure() *Result
}

type State int

const (
	Continue  State = iota
	Succeeded State = iota
	Failed    State = iota
)

type DefaultResult struct {
	State            State
	Success          int
	Failure          int
	SuccessThreshold int
	FailureThreshold int
}

func NewResult(successThreshold int, failureThreshold int) *DefaultResult {
	return &DefaultResult{
		State:            Continue,
		Success:          0,
		Failure:          0,
		SuccessThreshold: successThreshold,
		FailureThreshold: failureThreshold,
	}
}

func (r DefaultResult) IncrementSuccess() *DefaultResult {
	success := r.Success + 1

	state := Continue
	if success >= r.SuccessThreshold {
		state = Succeeded
	}

	return &DefaultResult{
		State:            state,
		Success:          success,
		Failure:          0,
		SuccessThreshold: r.SuccessThreshold,
		FailureThreshold: r.FailureThreshold,
	}
}

func (r DefaultResult) IncrementFailure() *DefaultResult {
	failure := r.Failure + 1

	state := Continue
	if failure >= r.FailureThreshold {
		state = Failed
	}

	return &DefaultResult{
		State:            state,
		Success:          0,
		Failure:          failure,
		SuccessThreshold: r.SuccessThreshold,
		FailureThreshold: r.FailureThreshold,
	}
}

func (r DefaultResult) Format() string {
	return fmt.Sprintf("Threshold: Success=%d/%d, Failure=%d/%d", r.Success, r.SuccessThreshold, r.Failure, r.FailureThreshold)
}
