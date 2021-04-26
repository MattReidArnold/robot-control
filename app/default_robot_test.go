package app_test

import (
	"fmt"
	"testing"

	"github.com/mattreidarnold/robot-control/app"
	"github.com/mattreidarnold/robot-control/test/helpers"
)

func TestDefaultRobot_Rotation(t *testing.T) {
	type DirTest struct {
		to   app.Direction
		from app.Direction
	}
	rotateTests := []struct {
		command  app.Command
		dirTests []DirTest
	}{
		{
			command: app.CommandRotateLeft,
			dirTests: []DirTest{
				{from: app.DirectionNorth, to: app.DirectionWest},
				{from: app.DirectionWest, to: app.DirectionSouth},
				{from: app.DirectionSouth, to: app.DirectionEast},
				{from: app.DirectionEast, to: app.DirectionNorth},
			},
		},
		{
			command: app.CommandRotateRight,
			dirTests: []DirTest{
				{from: app.DirectionNorth, to: app.DirectionEast},
				{from: app.DirectionWest, to: app.DirectionNorth},
				{from: app.DirectionSouth, to: app.DirectionWest},
				{from: app.DirectionEast, to: app.DirectionSouth},
			},
		},
	}
	for _, rt := range rotateTests {
		for _, dt := range rt.dirTests {
			t.Run(fmt.Sprintf("rotate %s when facing %s", rt.command, dt.from), func(t *testing.T) {
				want := dt.to
				wi := &app.WorldInhabitant{
					Direction: dt.from,
				}
				app.NewDefaultRobot()[rt.command](wi)
				got := wi.Direction

				if got != want {
					helpers.FailAssertion(t, "final direction", got, want)
				}
			})
		}
	}
}

func TestDefaultRobot_Movement(t *testing.T) {
	startFrom := app.Position{
		X: 0,
		Y: 0,
	}
	type MoveTest struct {
		direction app.Direction
		endAt     app.Position
	}

	moveTests := []MoveTest{
		{direction: app.DirectionNorth, endAt: app.Position{X: 0, Y: 1}},
		{direction: app.DirectionWest, endAt: app.Position{X: -1, Y: 0}},
		{direction: app.DirectionSouth, endAt: app.Position{X: 0, Y: -1}},
		{direction: app.DirectionEast, endAt: app.Position{X: 1, Y: 0}},
	}

	for _, mt := range moveTests {
		t.Run(fmt.Sprintf("move forward when facing %s", mt.direction), func(t *testing.T) {
			want := mt.endAt
			wi := &app.WorldInhabitant{
				Direction: mt.direction,
				Position:  startFrom,
			}
			app.NewDefaultRobot()[app.CommandMoveForward](wi)
			got := wi.Position

			if got != want {
				helpers.FailAssertion(t, "final position", got, want)
			}
		})
	}

}
