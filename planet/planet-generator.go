package planet

import (
	"fmt"
	"strconv"
	"strings"
)

// Generate plant - matrix map
func MapGenerator(line string) [][]int {
	dimensions := strings.Fields(line)
	xColumns, _ := strconv.ParseInt(dimensions[0], 10, 0 )
	yRows, _ := strconv.ParseInt(dimensions[1], 10, 0)
	grid := [][]int{}
	for i := 0; i < int(xColumns); i++ {
		fmt.Println(i)
		//create column
		curRow := make([]int, int(yRows))

		for j :=0; j< int(yRows); j++ {
			curRow[j] = 0
		}
		//create row
		grid = append(grid, curRow)
	}

	return grid
}