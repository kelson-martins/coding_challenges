package main

import (
	"fmt"
	"sort"
)

func main() {

	target := 6

	input := []int{3, 2, 4}

	var inputON []int
	inputON = append(inputON, input...)

	fmt.Println(twoSum(input, target))

	fmt.Println(twoSumON(inputON, target))
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

func twoSum(nums []int, target int) []int {

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
