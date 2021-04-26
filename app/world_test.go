package app_test

import (
	"testing"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/test/helpers"
)

func Test_WrappedGridWorld(t *testing.T) {
	xSize, ySize := int64(100), int64(100)
	testRuns := []struct {
		runName  string
		startsAt app.Orientation
		endsAt   app.Orientation
	}{
		{
			runName:  "within grid",
			startsAt: app.Orientation{Position: app.Position{X: 20, Y: 20}},
			endsAt:   app.Orientation{Position: app.Position{X: 20, Y: 20}},
		},
		{
			runName:  "off grid West",
			startsAt: app.Orientation{Position: app.Position{X: -5, Y: 20}},
			endsAt:   app.Orientation{Position: app.Position{X: 95, Y: 20}},
		},
		{
			runName:  "off grid North",
			startsAt: app.Orientation{Position: app.Position{X: 0, Y: 105}},
			endsAt:   app.Orientation{Position: app.Position{X: 0, Y: 5}},
		},
		{
			runName:  "off grid East",
			startsAt: app.Orientation{Position: app.Position{X: 101, Y: 99}},
			endsAt:   app.Orientation{Position: app.Position{X: 1, Y: 99}},
		},
		{
			runName:  "off grid South",
			startsAt: app.Orientation{Position: app.Position{X: 99, Y: -2}},
			endsAt:   app.Orientation{Position: app.Position{X: 99, Y: 98}},
		},
	}
	i := app.Instruction{}
	for _, run := range testRuns {
		t.Run(run.runName, func(t *testing.T) {
			want := run.endsAt
			got := run.startsAt
			app.WrappedGridWorld(xSize, ySize)(&got, i)

			if got != want {
				helpers.FailAssertion(t, "final orientation", got, want)
			}
		})
	}
}
