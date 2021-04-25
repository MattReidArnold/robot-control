package app

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInstructionNotFound = errors.New("instruction not found")
)

type Scenario struct {
	Direction
	Position
}

type InputProvider interface {
	StartDirection() Direction
	StartPosition() Position
	GetNextInstruction() (Instruction, error)
}

func NewScenario(d Direction, p Position) *Scenario {
	return &Scenario{
		Direction: d,
		Position:  p,
	}
}

func RunInstructionsPipeline(ip InputProvider, w io.Writer) error {
	s := NewScenario(ip.StartDirection(), ip.StartPosition())
	for {
		ni, err := ip.GetNextInstruction()
		if err == ErrInstructionNotFound {
			break
		}
		if err != nil {
			return err
		}
		executeInstruction(s, ni)
	}

	_, err := w.Write([]byte(fmt.Sprintf("%v", *s)))
	return err
}

func executeInstruction(s *Scenario, i Instruction) {
	for r := 0; r < i.Repetitions; r++ {
		switch i.Command {
		case CommandRotateLeft:
			rotateLeft(s)
		case CommandRotateRight:
			rotateRight(s)
		case CommandMoveForward:
			moveForward(s)
		}
	}
}

func rotateLeft(s *Scenario) {
	switch s.Direction {
	case DirectionNorth:
		s.Direction = DirectionWest
	case DirectionWest:
		s.Direction = DirectionSouth
	case DirectionSouth:
		s.Direction = DirectionEast
	case DirectionEast:
		s.Direction = DirectionNorth
	}
}

func rotateRight(s *Scenario) {
	switch s.Direction {
	case DirectionNorth:
		s.Direction = DirectionEast
	case DirectionWest:
		s.Direction = DirectionNorth
	case DirectionSouth:
		s.Direction = DirectionWest
	case DirectionEast:
		s.Direction = DirectionSouth
	}
}

func moveForward(s *Scenario) {
	switch s.Direction {
	case DirectionNorth:
		s.Y += 1
	case DirectionWest:
		s.X -= 1
	case DirectionSouth:
		s.Y -= 1
	case DirectionEast:
		s.X += 1
	}
	handleOffGrid(s)
}

func handleOffGrid(s *Scenario) {
	if s.X < 0 {
		s.X = 100 + s.X
	}
	if s.Y < 0 {
		s.Y = 100 + s.Y
	}
	if s.X > 99 {
		s.X = 100 - s.X
	}
	if s.Y > 99 {
		s.Y = 100 - s.Y
	}
}
