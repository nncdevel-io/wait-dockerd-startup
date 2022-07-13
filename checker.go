package main

import (
	"context"
	"github.com/docker/docker/client"
)

type Checker interface {
	Check() (string, error)
}

type DefaultChecker struct {
	client client.Client
}

func NewChecker(cli client.Client) Checker {
	return DefaultChecker{
		client: cli,
	}
}

func (c DefaultChecker) Check() (string, error) {
	info, err := c.client.Info(context.Background())

	if err != nil {
		return "", err
	} else {
		return info.ServerVersion, nil
	}
}
