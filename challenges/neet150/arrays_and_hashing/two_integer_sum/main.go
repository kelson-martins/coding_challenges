package main

import "fmt"

func two_sum(list []int, target int) []int {

	problemMap := make(map[int]int)

	for k, v := range list {

		_, ok := problemMap[target-v]

		if ok {
			idx1 := problemMap[target-v]
			idx2 := k
			fmt.Printf("The two indexes that adds to target %v are: %v, %v\n", target, idx1, idx2)
			return []int{problemMap[target-v], k}
		}

		problemMap[v] = k

	}

	return []int{}

}

func main() {

	type testCase struct {
		list   []int
		target int
	}

	testCases := []testCase{
		{
			list:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			list:   []int{3, 2, 4},
			target: 6,
		},
		{
			list:   []int{3, 3},
			target: 6,
		},
	}

	for k, v := range testCases {

		fmt.Printf("test case %v - %v\n", k+1, v.list)
		two_sum(v.list, v.target)

	}
}
