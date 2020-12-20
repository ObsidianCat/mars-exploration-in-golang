package robot

import (
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
