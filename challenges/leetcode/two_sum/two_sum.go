package main

import (
	"fmt"
	"sort"
)

func main() {

	target := 17

	input := []int{2, 7, 11, 15}

	var inputON []int
	inputON = append(inputON, input...)

	fmt.Println(twoSum(input, target))

	// fmt.Println(twoSumON(inputON, target))
}

func twoSumON(nums []int, target int) []int {

	mapNums := make(map[int]int)

	var v1 int
	var v2 int

	for i := 0; i < len(nums); i++ {

		toCheck := target - nums[i]
		val, ok := mapNums[toCheck]
		if ok {
			v1 = val
			v2 = i
			return []int{v1, v2}
		}

		mapNums[nums[i]] = i
	}

	return nil
}

func twoSum_1(nums []int, target int) []int {

	left := 0
	right := len(nums) - 1

	var copyNums []int
	copyNums = append(copyNums, nums...)

	sort.Ints(nums)

	for {
		if left > right {
			break
		}

		n1 := nums[left]
		n2 := nums[right]

		if n1+n2 == target {
			retN1 := -1
			retN2 := -1
			for i, _ := range copyNums {
				if copyNums[i] == n1 && retN1 == -1 {
					retN1 = i
					continue
				}
				if copyNums[i] == n2 && retN2 == -1 {
					retN2 = i
					continue
				}

				if retN1 != -1 && retN2 != -1 {
					break
				}
			}

			return []int{retN1, retN2}
		} else if n1+n2 > target {
			right--
			continue
		} else {
			left++
			continue
		}
	}

	return []int{}
}

// 11/09/2022
func twoSum(nums []int, target int) []int {

	for i := 0; i <= len(nums); i++ {

		toFind := target - nums[i]

		for j := i; j <= len(nums); j++ {
			if nums[j] == toFind {
				return []int{i, j}
			}
		}

	}

	return []int{}
}
