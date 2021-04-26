package app

type ControlFn func(*WorldInhabitant)

type Robot map[Command]ControlFn

func NewRobotControl(c Robot) ControlRobotFn {
	return func(wi *WorldInhabitant, i Instruction) {
		for r := 0; r < i.Repetitions; r++ {
			c[i.Command](wi)
		}
	}
}
