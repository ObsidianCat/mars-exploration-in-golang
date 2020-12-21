package planet

import (
	"strconv"
	"strings"
)

func stringIntoInt(st string) int {
	result, _ := strconv.ParseInt(st, 10, 0)
	return int(result)
}

// Generate plant - matrix map
func MapGenerator(line string) [][]int {
	//We start from 0 to the number including. So number 5 means that there will be 6 columns

	dimensions := strings.Fields(line)
	xColumns := stringIntoInt(dimensions[1])
	yRows := stringIntoInt(dimensions[0])
	grid := [][]int{}
	for i := 0; i <= xColumns; i++ {
		//create row
		// I add one to num of cell, because value represent length of array
		curRow := make([]int, yRows+1)

		//append row
		grid = append(grid, curRow)
	}

	return grid
}
