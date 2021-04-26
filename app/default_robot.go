package app

func NewDefaultRobot() Robot {
	return map[Command]ExecuteCommandFn{
		CommandMoveForward: moveForwardOneUnit,
		CommandRotateLeft:  rotateCounterClockwise90,
		CommandRotateRight: rotateClockwise90,
	}
}

func rotateCounterClockwise90(wi *WorldInhabitant) {
	switch wi.Direction {
	case DirectionNorth:
		wi.Direction = DirectionWest
	case DirectionWest:
		wi.Direction = DirectionSouth
	case DirectionSouth:
		wi.Direction = DirectionEast
	case DirectionEast:
		wi.Direction = DirectionNorth
	}
}

func rotateClockwise90(wi *WorldInhabitant) {
	switch wi.Direction {
	case DirectionNorth:
		wi.Direction = DirectionEast
	case DirectionWest:
		wi.Direction = DirectionNorth
	case DirectionSouth:
		wi.Direction = DirectionWest
	case DirectionEast:
		wi.Direction = DirectionSouth
	}
}

func moveForwardOneUnit(wi *WorldInhabitant) {
	switch wi.Direction {
	case DirectionNorth:
		wi.Y += 1
	case DirectionWest:
		wi.X -= 1
	case DirectionSouth:
		wi.Y -= 1
	case DirectionEast:
		wi.X += 1
	}
}
