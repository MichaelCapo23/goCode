package main

import (
	"fmt"
)

func getMaxHeight(r int, n int) int {
	max := r
	if r < n {
		max = r
	} else if n < r {
		max = n
	}
	return max
}

//width * max height combo
func maxArea(height []int) int {
	//store the first combo in max to save a loop
	max := 0

	//loop over the array, keep track of max area
	for index, row := range height {

		//loop over the rest of the array after the current index
		for index2, next := range height[index+1 : len(height)] {
			//get the next number to get calcs from
			maxHeight := getMaxHeight(row, next)
			currentWidth := index2 + 1
			value := maxHeight * currentWidth

			if value > max {
				max = value
			}
		}
	}

	return max
}

func main() {
	digits := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(digits))
}
