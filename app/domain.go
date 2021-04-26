package app

type Direction string

const (
	DirectionNorth = "N"
	DirectionSouth = "S"
	DirectionEast  = "E"
	DirectionWest  = "W"
)

type Position struct {
	X int64
	Y int64
}

type Command string

const (
	CommandMoveForward = "M"
	CommandRotateRight = "R"
	CommandRotateLeft  = "L"
)

type Instruction struct {
	Command
	Repetitions int
}

type Orientation struct {
	Direction
	Position
}
