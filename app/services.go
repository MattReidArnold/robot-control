package app

import (
	"errors"
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
type FormatResultFn func(*WorldInhabitant) []byte

func RunScenario(ip InputProvider, w io.Writer, controlRobot ControlRobotFn, format FormatResultFn) error {
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
		controlRobot(wi, ni)
	}

	_, err := w.Write(format(wi))
	return err
}
