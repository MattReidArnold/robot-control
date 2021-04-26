package files_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/files"
	"github.com/mattreidarnold/robot-control/test/helpers"
)

const exampleFilePath = "../test/fixtures/input_files/example.txt"

func OpenFile(path string) (*bufio.Reader, func()) {
	f, err := os.Open(exampleFilePath)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(f), func() {
		f.Close()
	}
}

func TestFileReader_New_with_example_file(t *testing.T) {
	f, close := OpenFile(exampleFilePath)
	defer close()

	_, got := files.NewFileReader(f)

	helpers.AssertNil(t, got)
}

func TestFileReader_StartDirection_with_example_file(t *testing.T) {
	var want app.Direction = app.DirectionNorth
	f, close := OpenFile(exampleFilePath)
	defer close()

	r, _ := files.NewFileReader(f)

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
	f, close := OpenFile(exampleFilePath)
	defer close()

	r, _ := files.NewFileReader(f)

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

	f, close := OpenFile(exampleFilePath)
	defer close()

	r, _ := files.NewFileReader(f)

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
