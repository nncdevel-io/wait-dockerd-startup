package app_test

import (
	"github.com/nncdevel-io/wait-dockerd-startup/app"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"testing"
	"time"
)

func testClientProvider(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("NewClientProvider returns new instance", func() {
		provider := app.NewClientProvider(time.Second)

		Expect(provider).ShouldNot(BeNil())
	})

	it("Provide method returns new client", func() {
		provider := app.NewClientProvider(time.Second)

		client, err := provider.Provide()

		Expect(client).ShouldNot(BeNil())
		Expect(err).ShouldNot(HaveOccurred())
	})

}
