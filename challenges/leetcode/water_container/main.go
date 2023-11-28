package main

import (
	"fmt"
)

func main() {
	// container := []int{3, 2, 4, 8}
	// container := []int{1, 1}
	container := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(container)
}
func maxArea(height []int) int {

	highestArea := 0
	counter := 0

	end := len(height)
	start := 0
	for start < end {

		rightWallValue := height[end-1]
		rightWallPos := len(height) - 1 - start

		leftWallValue := height[start]
		fmt.Println("iteration: ", counter+1)
		fmt.Println("Left Wall Value:", leftWallValue, "Right Wall Value:", rightWallValue)
		fmt.Println("Left Wall Pointer:", start, "Right Wall Pointer:", rightWallPos)

		area := 0

		if leftWallValue < rightWallValue {
			area = leftWallValue * (end - start - 1)
			start++
		} else {
			area = rightWallValue * (end - start - 1)
			end--
		}

		if highestArea < area {
			highestArea = area
		}

		fmt.Println("Highest so far", highestArea)
		fmt.Println("-----------------------------")

		counter++

		// if counter == 3 {
		// 	os.Exit(1)
		// }

	}

	fmt.Println("Highest Area: ", highestArea)

	return highestArea
}
