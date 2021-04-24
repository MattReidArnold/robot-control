package cmd

import (
	"fmt"

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
	//Read input
	r := files.NewFakeFileReader()
	d, p, i, err := r.Read()
	if err != nil {
		panic(err)
	}
	s := app.NewScenario(d, p)
	//Execute instructions
	app.ExecuteInstructions(s, i...)
	//print result
	fmt.Println(*s)

}
