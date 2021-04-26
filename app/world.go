package app

func (cr ControlRobotFn) World(wc ControlRobotFn) ControlRobotFn {
	return func(wi *WorldInhabitant, i Instruction) {
		cr(wi, i)
		wc(wi, i)
	}
}

func WrappedGridWorld(xSize, ySize int64) ControlRobotFn {
	return func(wi *WorldInhabitant, i Instruction) {
		if i.Command != CommandMoveForward {
			return
		}
		if wi.X < 0 {
			wi.X = xSize + wi.X
		}
		if wi.Y < 0 {
			wi.Y = ySize + wi.Y
		}
		if wi.X > 99 {
			wi.X = xSize - wi.X
		}
		if wi.Y > 99 {
			wi.Y = ySize - wi.Y
		}
	}
}
