package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

// iterative approach
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left > right {
			break
		}
		pivot := left + (right-left)/2
		value := nums[pivot]

		if value == target {
			return pivot
		} else if value > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}
