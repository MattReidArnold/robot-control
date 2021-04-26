package cmd

import (
	"bufio"
	"os"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/files"
	"github.com/spf13/cobra"
)

var (
	inFile  string
	fileCmd = &cobra.Command{
		Use:   "file",
		Short: "Read an input file for robot instructions",
		Run:   fileRun,
	}
)

func init() {
	fileCmd.Flags().StringVarP(&inFile, "in-file", "i", "", `Input file path (required). File must contain starting position and instructions`)
	fileCmd.MarkFlagRequired("in-file")
	rootCmd.AddCommand(fileCmd)
}

func fileRun(cmd *cobra.Command, args []string) {
	//Setup File reader
	f, err := os.Open(inFile)
	if err != nil {
		cmdErrs <- err
		return
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

	//Configure World
	world := robotControls.World(app.WrappedGridWorld(100, 100))

	//Execute Instruction Pipeline
	err = app.RunInstructionsPipeline(r, o, world)
	if err != nil {
		cmdErrs <- err
		return
	}
}
