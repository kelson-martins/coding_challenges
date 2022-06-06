package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{-4, -1, 0, 3, 10}

	fmt.Println(squares_root(nums))
}

func squares_root(nums []int) []int {

	left := 0
	right := len(nums) - 1

	result := append([]int{}, nums...) // making a copy as I do not want a pointer to the original input
	index := len(nums) - 1             // start to pupoluate the result array from the highest to lowest

	for {
		if left > right {
			break
		}

		leftSquare := math.Pow(float64(nums[left]), 2)
		rightSquare := math.Pow(float64(nums[right]), 2)
		if leftSquare > rightSquare {
			result[index] = int(leftSquare)
			left++
		} else {
			result[index] = int(rightSquare)
			right--
		}

		index--
	}

	return result
}
