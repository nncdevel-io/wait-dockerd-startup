package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/registry"
	"github.com/nncdevel-io/wait-dockerd-startup/app"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

type SuccessClient struct{}

func (c *SuccessClient) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error) {
	panic("implement me")
}

func (c *SuccessClient) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error) {
	panic("implement me")
}

func (c *SuccessClient) DiskUsage(ctx context.Context, opt types.DiskUsageOptions) (types.DiskUsage, error) {
	panic("implement me")
}

func (c *SuccessClient) Ping(ctx context.Context) (types.Ping, error) {
	panic("implement me")
}

func (c *SuccessClient) Info(ctx context.Context) (types.Info, error) {
	return types.Info{
		ServerVersion: "dummy-server-version",
	}, nil
}

type ErrorClient struct{}

func (c *ErrorClient) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error) {
	panic("implement me")
}

func (c *ErrorClient) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error) {
	panic("implement me")
}

func (c *ErrorClient) DiskUsage(ctx context.Context, opt types.DiskUsageOptions) (types.DiskUsage, error) {
	panic("implement me")
}

func (c *ErrorClient) Ping(ctx context.Context) (types.Ping, error) {
	panic("implement me")
}

func (c *ErrorClient) Info(ctx context.Context) (types.Info, error) {
	return types.Info{}, fmt.Errorf("dummy-error")
}

func testChecker(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("NewChecker returns new DefaultChecker Instance", func() {

		checker := app.NewChecker(&SuccessClient{})

		Expect(checker).ToNot(BeNil())
		_, ok := checker.(app.DefaultChecker)

		if ok == false {
			panic("error")
		}

	})

	it("Checker returns Server version", func() {

		checker := app.NewChecker(&SuccessClient{})

		version, err := checker.Check()

		Expect(version).Should(Equal("dummy-server-version"))
		Expect(err).ShouldNot(HaveOccurred())

	})

	it("Checker returns error", func() {

		checker := app.NewChecker(&ErrorClient{})

		version, err := checker.Check()

		Expect(version).Should(Equal(""))
		Expect(err).Should(HaveOccurred())

	})
}
