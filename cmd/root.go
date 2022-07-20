package cmd

import (
	"fmt"
	"github.com/nncdevel-io/wait-dockerd-startup/app"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var (
	initialDelay     time.Duration
	period           time.Duration
	timeout          time.Duration
	successThreshold int
	failureThreshold int
)

var rootCmd = &cobra.Command{
	Use:   "wait-dockerd-startup",
	Short: "Wait until docker daemon startup.",
	RunE: func(cmd *cobra.Command, args []string) error {

		options := &app.Options{
			InitialDelay:     initialDelay,
			Period:           period,
			Timeout:          timeout,
			SuccessThreshold: successThreshold,
			FailureThreshold: failureThreshold,
		}

		return app.Main(options)
	},
}

func init() {

	rootCmd.PersistentFlags().DurationVarP(&initialDelay, "initial-delay", "i", 10*time.Second, "Initial delay")
	rootCmd.PersistentFlags().DurationVarP(&period, "period", "p", 10*time.Second, "Check period.")
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 1*time.Second, "Docker API request timeout.")
	rootCmd.PersistentFlags().IntVarP(&successThreshold, "success-threshold", "s", 1, "Threshold value for detect Succeed.")
	rootCmd.PersistentFlags().IntVarP(&failureThreshold, "failure-threshold", "f", 10, "Threshold value for detect Failure.")

	if initialDelay < 0 {
		log.Fatal("illegal initial delay seconds. must positive value")
	}

	if period < 1*time.Second {
		log.Fatal("illegal period seconds. must greater than 0")
	}

	if timeout < 1*time.Second {
		log.Fatal("illegal timeout seconds. must greater than 0")
	}

	if successThreshold < 1 {
		log.Fatal("illegal success threshold. must greater than 0")
	}

	if failureThreshold < 1 {
		log.Fatal("illegal success threshold. must greater than 0")
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
