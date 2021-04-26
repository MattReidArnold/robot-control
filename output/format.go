package output

import (
	"fmt"

	"github.com/mattreidarnold/robot-control/app"
)

func FormatOrientation(o *app.Orientation) []byte {
	return []byte(fmt.Sprintf("%+v", *o))
}
