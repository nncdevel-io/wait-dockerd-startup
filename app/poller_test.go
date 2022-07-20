package app_test

import (
	"fmt"
	_ "github.com/docker/docker/client"
	"github.com/golang/mock/gomock"

	"github.com/nncdevel-io/wait-dockerd-startup/app"
	mock_main "github.com/nncdevel-io/wait-dockerd-startup/mock/app"
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
		poller := app.NewPoller(SuccessChecker{}, &app.PollingOptions{
			Period:           time.Duration(100),
			SuccessThreshold: 3,
			FailureThreshold: 10,
		})

		Expect(poller).ToNot(BeNil())
	})

	it("Poll() success", func() {
		poller := app.NewPoller(SuccessChecker{}, &app.PollingOptions{
			Period:           time.Duration(100),
			SuccessThreshold: 3,
			FailureThreshold: 10,
		})

		Expect(poller.Poll()).To(BeNil())
	})

	it("Poll() failed", func() {
		poller := app.NewPoller(FailedChecker{}, &app.PollingOptions{
			Period:           time.Duration(100),
			SuccessThreshold: 3,
			FailureThreshold: 10,
		})

		Expect(poller.Poll()).ToNot(BeNil())
	})

	it("mock sample", func() {
		ctrl := gomock.NewController(t)
		checker := mock_main.NewMockChecker(ctrl)
		poller := app.NewPoller(checker, &app.PollingOptions{
			Period:           time.Duration(100),
			SuccessThreshold: 3,
			FailureThreshold: 10,
		})

		checker.EXPECT().Check().Return("ok", nil).Times(3)

		Expect(poller.Poll()).To(BeNil())
	})
}
