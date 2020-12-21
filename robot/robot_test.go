package robot

import (
	"github.com/ObsidianCat/mars-exploration-in-golang/planet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssembleRobot(t *testing.T) {
	robot := New("1 2 E")
	assert := assert.New(t)

	assert.Equal(1, robot.xPosition, "Robot x position")
	assert.Equal(2, robot.yPosition, "Robot y position")
	assert.Equal("E", robot.direction, "Robot direction")
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

	for _, testCase := range calculatePathTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			robot := New(testCase.robotStartPosition)
			planetMap := planet.MapGenerator(testCase.planetSize)
			if testCase.scentModifier != nil {
				planetMap[testCase.scentModifier[0]][testCase.scentModifier[1]] = 1
			}
			result := CalculatePath(planetMap, robot, testCase.moveDirections)
			assert.Equal(t, testCase.expected, result, testCase.name)
		})
	}
}
