package app

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInstructionNotFound = errors.New("instruction not found")
)

type WorldInhabitant struct {
	Direction
	Position
}

type InputProvider interface {
	StartDirection() Direction
	StartPosition() Position
	GetNextInstruction() (Instruction, error)
}

type ControlRobotFn func(*WorldInhabitant, Instruction)

func RunInstructionsPipeline(ip InputProvider, w io.Writer, cr ControlRobotFn) error {
	wi := &WorldInhabitant{
		Direction: ip.StartDirection(),
		Position:  ip.StartPosition(),
	}
	for {
		ni, err := ip.GetNextInstruction()
		if err == ErrInstructionNotFound {
			break
		}
		if err != nil {
			return err
		}
		cr(wi, ni)
	}

	_, err := w.Write([]byte(fmt.Sprintf("%v", *wi)))
	return err
}

// func handleOffGrid(wi *WorldInhabitant) {
// 	if wi.X < 0 {
// 		wi.X = 100 + wi.X
// 	}
// 	if wi.Y < 0 {
// 		wi.Y = 100 + wi.Y
// 	}
// 	if wi.X > 99 {
// 		wi.X = 100 - wi.X
// 	}
// 	if wi.Y > 99 {
// 		wi.Y = 100 - wi.Y
// 	}
// }
