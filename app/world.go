package app

func (cr ControlRobotFn) InWorld(wc ControlRobotFn) ControlRobotFn {
	return func(o *Orientation, i Instruction) {
		cr(o, i)
		wc(o, i)
	}
}

func WrappedGridWorld(xSize, ySize int64) ControlRobotFn {
	return func(o *Orientation, i Instruction) {
		if i.Command != CommandMoveForward {
			return
		}
		if o.X < 0 {
			o.X = xSize + o.X
		}
		if o.Y < 0 {
			o.Y = ySize + o.Y
		}
		if o.X > 99 {
			o.X = xSize - o.X
		}
		if o.Y > 99 {
			o.Y = ySize - o.Y
		}
	}
}
