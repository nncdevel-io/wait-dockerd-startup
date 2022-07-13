package main

import (
	"flag"
	"time"
)

type Options struct {
	Interval       time.Duration
	Timeout        time.Duration
	RequestTimeout time.Duration
}

func OptionsFromEnv() *Options {
	interval := flag.Int("interval", 1, "check interval.")
	timeout := flag.Int("timeout", 60, "start up timeout.")
	requestTimeout := flag.Int("request-timeout", 3, "docker api request timeout.")

	flag.Parse()

	return &Options{
		Interval:       time.Duration(*interval) * time.Second,
		Timeout:        time.Duration(*timeout) * time.Second,
		RequestTimeout: time.Duration(*requestTimeout) * time.Second,
	}
}
