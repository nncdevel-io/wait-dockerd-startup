package app_test

import (
	"github.com/nncdevel-io/wait-dockerd-startup/app"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"testing"
)

func testResult(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("NewResult returns new DefaultResult", func() {
		result := app.NewResult(10, 20)

		Expect(result).ShouldNot(BeNil())

		Expect(result.State).Should(Equal(app.Continue))
		Expect(result.Success).Should(Equal(0))
		Expect(result.Failure).Should(Equal(0))
		Expect(result.SuccessThreshold).Should(Equal(10))
		Expect(result.FailureThreshold).Should(Equal(20))
	})

	it("IncrementSuccess method return result that success count has increment", func() {
		initialResult := app.NewResult(10, 20)

		result := initialResult.IncrementSuccess()

		Expect(result.State).Should(Equal(app.Continue))
		Expect(result.Success).Should(Equal(1))
		Expect(result.Failure).Should(Equal(0))
		Expect(result.SuccessThreshold).Should(Equal(10))
		Expect(result.FailureThreshold).Should(Equal(20))
	})

	it("IncrementSuccess threshold", func() {
		initialResult := app.NewResult(1, 10)

		result := initialResult.IncrementSuccess()

		Expect(result.State).Should(Equal(app.Succeeded))
		Expect(result.Success).Should(Equal(1))
		Expect(result.Failure).Should(Equal(0))
		Expect(result.SuccessThreshold).Should(Equal(1))
		Expect(result.FailureThreshold).Should(Equal(10))
	})

	it("IncrementFailure method return result that success count has increment", func() {
		initialResult := app.NewResult(10, 20)

		result := initialResult.IncrementFailure()

		Expect(result.State).Should(Equal(app.Continue))
		Expect(result.Success).Should(Equal(0))
		Expect(result.Failure).Should(Equal(1))
		Expect(result.SuccessThreshold).Should(Equal(10))
		Expect(result.FailureThreshold).Should(Equal(20))
	})

	it("IncrementFailure threshold", func() {
		initialResult := app.NewResult(10, 1)

		result := initialResult.IncrementFailure()

		Expect(result.State).Should(Equal(app.Failed))
		Expect(result.Success).Should(Equal(0))
		Expect(result.Failure).Should(Equal(1))
		Expect(result.SuccessThreshold).Should(Equal(10))
		Expect(result.FailureThreshold).Should(Equal(1))
	})
}
