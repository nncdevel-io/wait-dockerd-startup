package main

import (
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"testing"
)

func TestUnit(t *testing.T) {
	suite := spec.New("wait-dockerd", spec.Report(report.Terminal{}))
	suite("Poller", testPoller)
	suite.Run(t)
}
