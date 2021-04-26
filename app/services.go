package app

import (
	"errors"
	"io"
)

var (
	ErrInstructionNotFound = errors.New("instruction not found")
)

type InputProvider interface {
	StartDirection() Direction
	StartPosition() Position
	GetNextInstruction() (Instruction, error)
}

type ControlRobotFn func(*Orientation, Instruction)
type FormatResultFn func(*Orientation) []byte

func RunScenario(ip InputProvider, w io.Writer, controlRobot ControlRobotFn, format FormatResultFn) error {
	o := &Orientation{
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
		controlRobot(o, ni)
	}

	_, err := w.Write(format(o))
	return err
}
