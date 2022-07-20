//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package app

import (
	"context"
	"github.com/docker/docker/client"
)

type Checker interface {
	Check() (string, error)
}

type DefaultChecker struct {
	client client.SystemAPIClient
}

func NewChecker(client client.SystemAPIClient) Checker {
	return DefaultChecker{
		client: client,
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
