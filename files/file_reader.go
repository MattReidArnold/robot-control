package files

import (
	"github.com/mattreidarnold/robot-control/app"
)

// N 0 0
// M1RM4L3M2

type fakeFileReader struct {
	curInst      int
	instructions []app.Instruction
}

func NewFakeFileReader() app.InputProvider {
	return &fakeFileReader{
		curInst: 0,
		instructions: []app.Instruction{
			{Command: app.CommandMoveForward, Repetitions: 1},
			{Command: app.CommandRotateRight, Repetitions: 1},
			{Command: app.CommandMoveForward, Repetitions: 4},
			{Command: app.CommandRotateLeft, Repetitions: 3},
			{Command: app.CommandMoveForward, Repetitions: 2},
		},
	}
}

func (f *fakeFileReader) StartDirection() app.Direction {
	return app.DirectionNorth
}

func (f *fakeFileReader) StartPosition() app.Position {
	return app.Position{X: 0, Y: 0}
}

func (f *fakeFileReader) GetNextInstruction() (app.Instruction, error) {
	if f.curInst >= len(f.instructions) {
		return app.Instruction{}, app.ErrInstructionNotFound
	}
	i := f.instructions[f.curInst]
	f.curInst++
	return i, nil
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
