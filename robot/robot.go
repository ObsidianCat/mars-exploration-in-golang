package robot

import (
	"fmt"
	"strconv"
	"strings"
)

const DirectionRight string = "R"
const DirectionLeft string = "L"

// TODO iota enum?
var sides = [4]string{"N", "E", "S", "W"}

type Robot struct {
	xPosition int
	yPosition int
	direction string
}

func (r *Robot) getFormattedPosition() string {
	return fmt.Sprintf("%s %s", strconv.Itoa(r.xPosition), strconv.Itoa(r.yPosition))
}

func (r *Robot) rotate(turnSide string) string {
	modifier := 0
	if turnSide == DirectionLeft {
		modifier = -1
	} else if turnSide == DirectionRight {
		modifier = 1
	}
	var possible int
	for i, side := range sides {
		if side == r.direction {
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

func (r *Robot) setPosition(x, y int) {
	r.xPosition = x
	r.yPosition = y
}

// Calculate full path for the Robot
func CalculatePath(planetMap [][]int, robot Robot, directions string) string {

	rotationOnly := false
	for _, curChar := range directions {
		charString := string(curChar)
		if charString == DirectionRight || charString == DirectionLeft {
			robot.rotate(charString)
		} else if !rotationOnly {

			//It is moving direction and robot can move (not frightened by smell sign)
			column, row := robot.calcNextPosition()
			isBeyondMap := false
			if column < 0 || row < 0 || column >= len(planetMap[0]) || row >= len(planetMap) {
				isBeyondMap = true
			}
			if isBeyondMap {
				//leave scent
				planetMap[robot.yPosition][robot.xPosition] = 1
				return fmt.Sprintf("%s %s LOST", robot.getFormattedPosition(), robot.direction)
			}

			nextHasScentTrace := planetMap[column][row] == 1
			if nextHasScentTrace {
				//Stopped by scent, no move movements, only rotation
				rotationOnly = true
				continue
			}
			robot.setPosition(column, row)
		}

	}
	return fmt.Sprintf("%s %s", robot.getFormattedPosition(), robot.direction)
}

func New(text string) Robot {
	dimensions := strings.Fields(text)
	x, _ := strconv.Atoi(dimensions[0])
	y, _ := strconv.Atoi(dimensions[1])

	newRobot := Robot{direction: dimensions[2],
		xPosition: x,
		yPosition: y,
	}

	return newRobot
}
