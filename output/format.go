package output

import (
	"fmt"

	"github.com/mattreidarnold/robot-control/app"
)

func FormatWorldInhabitant(wi *app.WorldInhabitant) []byte {
	return []byte(fmt.Sprintf("%+v", *wi))
}
