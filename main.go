package main

import (
	"bufio"
	"fmt"
	"github.com/ObsidianCat/mars-exploration-in-golang/planet"
	"github.com/ObsidianCat/mars-exploration-in-golang/robot"

	//"io"
	"os"
	//"strings"
)

const (
	MAP = iota + 1
	POSITION
	DIRECTIONS
)

func main() {
	currentInputState := MAP
	var currentMap [][]int
	var currentRobot robot.Robot

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Provide map dimensions")

	for {
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) == 0 {
			fmt.Println("The mission completed")
			break
		}

		switch currentInputState {
		case MAP:
			planet.MapGenerator(text)
			fmt.Println("Provide first robot initial position'")
			currentInputState = POSITION

		case POSITION:
			currentRobot = robot.AssembleRobot(text)
			fmt.Println("Provide path instructions")
			currentInputState = DIRECTIONS

		case DIRECTIONS:
			fmt.Println("This robot last known position is")
			fmt.Println(robot.CalculatePath(currentMap, currentRobot, text))

			currentInputState = POSITION
			fmt.Println("Provide next robot initial position")

		}
		fmt.Println(text)

	}
}
