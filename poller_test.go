package main

import (
	"fmt"
	_ "github.com/docker/docker/client"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"testing"
	"time"
)

type SuccessChecker struct {
}

func (d SuccessChecker) Check() (string, error) {
	return "ok", nil
}

type FailedChecker struct {
}

func (f FailedChecker) Check() (string, error) {
	return "", fmt.Errorf("dummy Error")
}

func testPoller(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("NewChecker returns new Instance", func() {
		poller := NewPoller(SuccessChecker{}, &Options{
			Interval:       time.Duration(100),
			Timeout:        time.Duration(100),
			RequestTimeout: time.Duration(100),
		})

		Expect(poller).ToNot(BeNil())
	})

	it("Poll() success", func() {
		poller := NewPoller(SuccessChecker{}, &Options{
			Interval:       time.Duration(100),
			Timeout:        time.Duration(100),
			RequestTimeout: time.Duration(100),
		})

		Expect(poller.Poll()).To(BeNil())
	})

	it("Poll() failed", func() {
		poller := NewPoller(FailedChecker{}, &Options{
			Interval:       time.Duration(100),
			Timeout:        time.Duration(100),
			RequestTimeout: time.Duration(100),
		})

		Expect(poller.Poll()).ToNot(BeNil())
	})
}
