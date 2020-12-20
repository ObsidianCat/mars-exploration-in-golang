package robot

import (
	"github.com/ObsidianCat/mars-exploration-in-golang/planet"
	"testing"
)

func TestAssembleRobot(t *testing.T) {
	robot := AssembleRobot("1 2 E")
	if robot.xPosition != 1 {
		t.Errorf("Robot x position expected: 1, actual %d", robot.xPosition)
	}
	if robot.yPosition != 2 {
		t.Errorf("Robot x position expected: 1, actual %d", robot.yPosition)
	}
	if robot.direction != "E" {
		t.Errorf("Robot x position expected: E, actual %s", robot.direction)
	}
}

var calculatePathTestCases = []struct {
	name               string
	robotStartPosition string
	planetSize         string
	moveDirections     string
	scentModifier      []int
	expected           string
}{
	{"calculate path for Robot", "1 2 N", "5 5", "LMLMLMLMM", nil, "1 3 N"},
	{"calculate path for Robot whose final position is equal to initial position", "1 1 E", "5 3", "RFRFRFRF", nil, "1 1 E"},
	{"calculate path for Robot who got lost", "3 2 N", "5 3", "FRRFLLFFRRFLL", nil, "3 3 N LOST"},
	{"calculate path for Robot who wasn`t lost because of scent trace", "0 3 W", "5 3", "LLFFFLFLFL", []int{3, 3}, "2 3 S"},
}

func TestCalculatePath(t *testing.T) {
	for _, tt := range calculatePathTestCases {
		t.Run(tt.name, func(t *testing.T) {
			robot := AssembleRobot(tt.robotStartPosition)
			planetMap := planet.MapGenerator(tt.planetSize)
			if tt.scentModifier != nil {
				planetMap[tt.scentModifier[0]][tt.scentModifier[1]] = 1
			}
			result := CalculatePath(planetMap, robot, tt.moveDirections)
			if result != tt.expected {
				t.Errorf("Robot final position %s does not match expected position %s", result, tt.expected)
			}
		})
	}
}
