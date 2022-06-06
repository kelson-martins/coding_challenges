package main

import "fmt"

var input []int
var target int
var output int

func main() {

	input = []int{1, 3}
	target = 2
	output = 0

	// fmt.Println(searchInsertIterative(input, target))
	fmt.Println(recursiveCall(input, target))
}

func searchInsertIterative(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for true {
		middle := left + (right-left)/2

		value := nums[middle]

		if target == value {
			return middle
		} else if value > target {
			if middle == 0 {
				return 0 // It should be inserted at the last position
			} else if nums[middle-1] > target {
				right = middle - 1
			} else if nums[middle-1] < target {
				return middle
			} else { // is equal
				return middle - 1
			}
		} else {
			if middle == len(nums)-1 {
				return len(nums)
			} else if nums[middle+1] > target {
				return middle + 1
			} else if nums[middle+1] < target {
				left = middle + 1
			} else { // is equal
				return middle + 1
			}
		}
	}

	return 0

}

func recursiveCall(nums []int, target int) int {
	return searchRecursive(input, target, 0, len(input)-1)
}

func searchRecursive(nums []int, target int, left int, right int) int {

	middle := left + (right-left)/2

	value := nums[middle]

	if target == value {
		return middle
	} else if value > target {

		if middle == 0 {
			return 0
		}

		if nums[middle-1] > target {
			right = middle - 1
			return searchRecursive(nums, target, left, right)
		} else if nums[middle-1] < target {
			return middle
		} else {
			return middle - 1
		}
	} else {

		if middle == len(nums)-1 {
			return len(nums)
		}
		if nums[middle+1] > target {
			return middle + 1
		} else if nums[middle+1] < target {
			left = middle + 1
			return searchRecursive(nums, target, left, right)
		} else {
			return middle + 1
		}
	}
}
