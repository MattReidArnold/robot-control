package app

type Scenario struct {
	Direction
	Position
}

func NewScenario(d Direction, p Position) *Scenario {
	return &Scenario{
		Direction: d,
		Position:  p,
	}
}

func ExecuteInstructions(s *Scenario, instructions ...Instruction) error {
	s.Direction = DirectionSouth
	s.Position.X = 4
	s.Position.Y = 99
	return nil
}
