package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Read an input file for robot instructions",
	Run:   fileRun,
}

func init() {
	rootCmd.AddCommand(fileCmd)
}

func fileRun(cmd *cobra.Command, args []string) {
	fmt.Println("Domo Origato Mr Roboto")
}
