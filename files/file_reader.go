package files

import "github.com/mattreidarnold/robot-control/app"

// N 0 0
// M1RM4L3M2

type FileReader interface {
	Read() (app.Direction, app.Position, []app.Instruction, error)
}

type fakeFileReader struct{}

func NewFakeFileReader() FileReader {
	return &fakeFileReader{}
}

func (f *fakeFileReader) Read() (app.Direction, app.Position, []app.Instruction, error) {
	return app.DirectionNorth,
		app.Position{X: 0, Y: 0},
		[]app.Instruction{
			{Command: app.CommandMoveForward, Repetitions: 1},
			{Command: app.CommandRotateRight, Repetitions: 1},
			{Command: app.CommandMoveForward, Repetitions: 4},
			{Command: app.CommandRotateLeft, Repetitions: 3},
			{Command: app.CommandMoveForward, Repetitions: 2},
		},
		nil
}
