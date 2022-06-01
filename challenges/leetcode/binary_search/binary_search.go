package main

import "fmt"

func main() {
	fmt.Println(searchRecursive([]int{-1, 0, 3, 5, 9, 12}, 9))
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

func searchRecursive(nums []int, target int) int {

	return (recursive(nums, target, 0, len(nums)-1))
}

func recursive(nums []int, target int, left int, right int) int {

	pivot := left + (right-left)/2

	if left > right {
		return -1
	}

	if nums[pivot] == target {
		return pivot
	} else if nums[pivot] < target {
		return recursive(nums, target, pivot+1, right)
	} else {
		return recursive(nums, target, left, pivot-1)
	}
}
