package e2e

import (
	"bufio"
	"os"
	"testing"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/files"
	"github.com/mattreidarnold/robot-control/test/helpers"
)

func TestExampleFile(t *testing.T) {
	want := "{S {4 99}}"
	//Setup File reader
	f, err := os.Open("../fixtures/input_files/example.txt")
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
	o := NewMockWriter()

	//Configure Robot
	robotControls := app.NewRobotControl(app.NewDefaultRobot())

	//Configure World
	world := robotControls.World(app.WrappedGridWorld(100, 100))

	//Execute Instruction Pipeline
	err = app.RunScenario(r, o, world)

	helpers.AssertNil(t, err)

	got := o.Received()

	if want != got {
		helpers.FailAssertion(t, "expected output", got, want)
	}

}
