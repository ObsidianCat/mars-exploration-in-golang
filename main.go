package main

import (
	"bufio"
	"fmt"
	"github.com/ObsidianCat/mars-exploration-in-golang/planet"
	//"io"
	"os"
	//"strings"
)

const (
	MAP = iota +1
	POSITION
	DIRECTIONS
)


func main() {
	currentInputState := MAP
	// To create dynamic array
	//arr := make([]string, 0)
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
			currentInputState = POSITION
			planet.MapGenerator(text)
			fmt.Println("Provide first robot initial position'")


		case POSITION:
			currentInputState = DIRECTIONS
			fmt.Println("Provide path instructions")


		case DIRECTIONS:
			fmt.Println("This robot last known position is")

			currentInputState = POSITION
			fmt.Println("Provide next robot initial position")

		}
		fmt.Println(text)

	}
	// Use collected inputs
	//fmt.Println(arr)
}