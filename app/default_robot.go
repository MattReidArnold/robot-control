package app

func NewDefaultRobot() Robot {
	return map[Command]ExecuteCommandFn{
		CommandMoveForward: moveForwardOneUnit,
		CommandRotateLeft:  rotateCounterClockwise90,
		CommandRotateRight: rotateClockwise90,
	}
}

func rotateCounterClockwise90(o *Orientation) {
	switch o.Direction {
	case DirectionNorth:
		o.Direction = DirectionWest
	case DirectionWest:
		o.Direction = DirectionSouth
	case DirectionSouth:
		o.Direction = DirectionEast
	case DirectionEast:
		o.Direction = DirectionNorth
	}
}

func rotateClockwise90(o *Orientation) {
	switch o.Direction {
	case DirectionNorth:
		o.Direction = DirectionEast
	case DirectionWest:
		o.Direction = DirectionNorth
	case DirectionSouth:
		o.Direction = DirectionWest
	case DirectionEast:
		o.Direction = DirectionSouth
	}
}

func moveForwardOneUnit(o *Orientation) {
	switch o.Direction {
	case DirectionNorth:
		o.Y += 1
	case DirectionWest:
		o.X -= 1
	case DirectionSouth:
		o.Y -= 1
	case DirectionEast:
		o.X += 1
	}
}
