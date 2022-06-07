package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{3, 2, 4}
	target := -6

	fmt.Println(twoSum(input, target))
}

func twoSum(nums []int, target int) []int {

	left := 0
	right := len(nums) - 1

	for {
		if left > right {
			break
		}

		n1 := int(math.Abs(float64(nums[left])))
		n2 := int(math.Abs(float64(nums[right])))

		newTarget := int(math.Abs(float64(target)))

		if n1+n2 == newTarget {
			return []int{left, right}
		} else if n1+n2 > newTarget {
			right--
			continue
		} else {
			left++
			continue
		}
	}

	return []int{}
}
