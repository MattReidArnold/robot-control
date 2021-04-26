package app

type ExecuteCommandFn func(*Orientation)

type Robot map[Command]ExecuteCommandFn

func NewRobotControl(c Robot) ControlRobotFn {
	return func(o *Orientation, i Instruction) {
		for r := 0; r < i.Repetitions; r++ {
			c[i.Command](o)
		}
	}
}
