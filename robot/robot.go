package robot

import (
	"fmt"
	"strconv"
	"strings"
)

const DirectionRight string = "R"
const DirectionLeft string = "L"

var sides = [4]string{"N", "E", "S", "W"}

type Robot struct {
	xPosition int
	yPosition int
	direction string
	sides     [4]string
}

func (r Robot) getFormattedPosition() string {
	return "x position: " + strconv.Itoa(r.xPosition) + ", y posiion: " + strconv.Itoa(r.yPosition)
}
func (r Robot) getDirection() string {
	return r.direction
}

func (r Robot) rotate(direction string) string {
	modifier := 0
	if direction == DirectionLeft {
		modifier = -1
	} else if direction == DirectionRight {
		modifier = 1
	}
	var possible int
	for i := 0; i < len(sides); i++ {
		if sides[i] == direction {
			possible = i + modifier
			break
		}
	}

	if possible >= len(sides) {
		r.direction = sides[0]
	} else if possible < 0 {
		r.direction = sides[len(sides)-1]
	} else {
		r.direction = sides[possible]
	}

	return r.direction
}

func (r Robot) calcNextPosition() (int, int) {
	column := r.xPosition
	row := r.yPosition
	switch r.direction {
	case "E":
		// Going right, instead of left
		column++
	case "S":
		// Up
		row--
	case "W":
		// Going left, instead of right
		column--
	case "N":
		// Down
		row++
	}
	return column, row
}

func (r Robot) setPosition(x, y int) {
	r.xPosition = x
	r.yPosition = y
}

// Calculate ful path for the Robot
func CalculatePath(planetMap [][]int, robot Robot, directions string) string {
	for i, char := range directions {
		fmt.Println(i, " => ", string(char))
		//if char == DirectionRight || char == DirectionLeft {
		//
		//}
	}
	return "It will return ${robot.getFormattedPosition()} ${robot.getDirection()} "
}

func AssembleRobot(text string) Robot {
	dimensions := strings.Fields(text)
	x, _ := strconv.Atoi(dimensions[0])
	y, _ := strconv.Atoi(dimensions[1])

	fmt.Println(dimensions)
	newRobot := Robot{direction: dimensions[2],
		xPosition: x,
		yPosition: y,
		sides:     sides,
	}

	return newRobot
}
