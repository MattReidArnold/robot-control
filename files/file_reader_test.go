package files_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/files"
	"github.com/mattreidarnold/robot-control/test/helpers"
)

func TestFileReader_New_with_example_file(t *testing.T) {
	f, err := os.Open("../tmp/input_files/example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, got := files.NewFileReader(bufio.NewReader(f))

	helpers.AssertNil(t, got)
}

func TestFileReader_StartDirection_with_example_file(t *testing.T) {
	var want app.Direction = app.DirectionNorth
	f, err := os.Open("../tmp/input_files/example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, _ := files.NewFileReader(bufio.NewReader(f))

	got := r.StartDirection()

	if got != want {
		helpers.FailAssertion(t, "StartDirection", got, want)
	}
}

func TestFileReader_Position_with_example_file(t *testing.T) {
	want := app.Position{
		X: 0,
		Y: 0,
	}
	f, err := os.Open("../tmp/input_files/example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, _ := files.NewFileReader(bufio.NewReader(f))

	got := r.StartPosition()

	if got != want {
		helpers.FailAssertion(t, "StartPositionX", got, want)
	}
}

func TestFileReader_GetNextInstruction_with_example_file(t *testing.T) {
	want := []app.Instruction{
		{Command: app.CommandMoveForward, Repetitions: 1},
		{Command: app.CommandRotateRight, Repetitions: 1},
		{Command: app.CommandMoveForward, Repetitions: 4},
		{Command: app.CommandRotateLeft, Repetitions: 3},
		{Command: app.CommandMoveForward, Repetitions: 2},
	}

	f, err := os.Open("../tmp/input_files/example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, _ := files.NewFileReader(bufio.NewReader(f))

	var got []app.Instruction
	calls := 0
	for {
		i, err := r.GetNextInstruction()
		if err == app.ErrInstructionNotFound {
			break
		}
		helpers.AssertNil(t, err)
		calls = calls + 1
		if calls > len(want) {
			t.Error("called GetNextInstruction too many times")
			return
		}
		got = append(got, i)
	}

	helpers.AssertDeepEqual(t, got, want)
}
