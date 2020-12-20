package robot

import (
	"fmt"
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

func TestCalculatePath(t *testing.T) {
	robot := AssembleRobot("1 1 E")
	planetMap := planet.MapGenerator("5 3")
	result := CalculatePath(planetMap, robot, "RFRFRFRF")
	fmt.Println(result)
	if result != "1 1 E" {
		t.Errorf("Robot final position. Expected: 1 1 E, actual %s.", result)
	}
}
