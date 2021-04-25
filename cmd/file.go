package cmd

import (
	"os"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/files"
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
	//Setup Input
	r := files.NewFakeFileReader()

	//Setup OutStream
	o := os.Stdout

	//Execute Instruction Pipeline
	err := app.RunInstructionsPipeline(r, o)
	if err != nil {
		panic(err)
	}
}
