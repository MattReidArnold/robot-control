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

func RunScenario(ip InputProvider, w io.Writer, controlRobot ControlRobotFn) error {
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

	_, err := w.Write([]byte(fmt.Sprintf("%v", *wi)))
	return err
}
