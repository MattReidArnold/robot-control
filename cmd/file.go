package cmd

import (
	"bufio"
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
	//Setup File reader
	f, err := os.Open("tmp/input_files/example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//Setup Input
	r, err := files.NewFileReader(bufio.NewReader(f))
	if err != nil {
		panic(err)
	}

	//Setup OutStream
	o := os.Stdout

	//Configure Robot
	robotControls := app.NewRobotControl(app.NewDefaultRobot())

	//Execute Instruction Pipeline
	err = app.RunInstructionsPipeline(r, o, robotControls)
	if err != nil {
		panic(err)
	}
}
