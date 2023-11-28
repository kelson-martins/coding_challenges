package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isPalindrome(-121))
}

func isPalindrome(x int) bool {

	checkVals := strconv.Itoa(x)

	end := len(checkVals) - 1

	if fmt.Sprintf("%c", checkVals[0]) == "-" {
		return false
	}

	for start := 0; start <= len(checkVals)/2; start++ {

		c1 := fmt.Sprintf("%c", checkVals[start])
		c2 := fmt.Sprintf("%c", checkVals[end])
		fmt.Println(c1, c2)
		if c1 != c2 {
			return false
		}

		end--

	}

	return true

}
