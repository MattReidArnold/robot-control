package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	cmdErrs chan error
	rootCmd = &cobra.Command{
		Use: "robot-control [command]",
	}
)

// Execute runs the root command
func Execute() {
	cmdErrs = make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		cmdErrs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		cmdErrs <- rootCmd.Execute()
	}()

	err := <-cmdErrs
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
