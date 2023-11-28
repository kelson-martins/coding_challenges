package main

import (
	"fmt"
	"strings"
)

func main() {

	var input1 string
	var input2 string

	fmt.Print("enter string 1: ")
	fmt.Scanln(&input1)
	fmt.Print("enter string 2: ")
	fmt.Scanln(&input2)

	fmt.Println("Is permutation ?  ", checkPermutation(input1, input2))
}

func checkPermutation(str1 string, str2 string) bool {

	map1 := map[string]int{}
	map2 := map[string]int{}

	str1Slice := strings.Split(str1, "")
	str2Slice := strings.Split(str2, "")

	// if the sizes of strings are different, its not permutation
	if len(str1) != len(str2) {
		return false
	}

	// populate two hashmaps with the total for each character
	for i := 0; i < len(str1); i++ {

		charStr1 := str1Slice[i]
		charStr2 := str2Slice[i]

		map1[charStr1] += 1
		map2[charStr2] += 1
	}

	// compare the occurence of each character between the two maps
	for map1k, map1v := range map1 {

		map2v, ok := map2[map1k]
		if !(ok) {
			return false
		}

		if map1v != map2v {
			return false
		}

	}

	return true
}
