package main

import (
	"github.com/docker/docker/client"
	"time"
)

type ClientProvider interface {
	Provide() (*client.Client, error)
}

type DefaultClientProvider struct {
	RequestTimeout time.Duration
}

func NewClientProvider(requestTimeout time.Duration) ClientProvider {
	return DefaultClientProvider{
		RequestTimeout: requestTimeout,
	}
}

func (p DefaultClientProvider) Provide() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithTimeout(p.RequestTimeout))
}
