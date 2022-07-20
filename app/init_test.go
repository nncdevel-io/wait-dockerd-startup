package app_test

import (
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"testing"
)

func TestUnit(t *testing.T) {
	suite := spec.New("wait-dockerd", spec.Report(report.Terminal{}))
	suite("Poller", testPoller)
	suite("Checker", testChecker)
	suite("ClientProvider", testClientProvider)
	suite("Result", testResult)
	suite.Run(t)
}
