package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "robot-control [command]",
}

// Execute runs the root command
func Execute() {
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errs <- fmt.Errorf("%v", r)
			}
		}()
		errs <- rootCmd.Execute()
	}()

	err := <-errs
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
