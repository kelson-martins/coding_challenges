package main

import "fmt"

var testCases map[string]int
var currentTest string
var expectedResult string

func main() {
	testCases = make(map[string]int)
	testCases["t1"] = 5
	testCases["r1"] = 4

	testCases["t2"] = 1
	testCases["r2"] = 1

	currentTest = "t1"
	expectedResult = "r1"

	input := testCases[currentTest]

	result := firstBadVersionRecursive(input)
	if result == testCases[expectedResult] {
		fmt.Printf("Correct. Output: %d. Expected: %d.\n", result, testCases[expectedResult])
	} else {
		fmt.Printf("Failed. Output: %d. Expected: %d.\n", result, testCases[expectedResult])
	}

}

func firstBadVersionRecursive(input int) int {
	return (recursive(input, 0, input-1))
}

func recursive(input int, left int, right int) int {

	for i := 0; i <= input; i++ {

		middle := left + (right-left)/2

		if isBadVersion(middle) {

			if isBadVersion(middle - 1) {
				return recursive(input, left, middle-1)
			}
			return middle
		} else {
			left = middle + 1
			continue
		}
	}

	return -1
}

func firstBadVersion(n int) int {

	left := 0
	right := n
	firstBad := 0

	for i := 0; i <= n; i++ {

		pivot := left + (right-left)/2

		isBad := isBadVersion(pivot)

		if isBad {
			if !isBadVersion(pivot - 1) {
				firstBad = pivot
				break
			} else {
				right = pivot - 1
				continue
			}
		} else {
			left = pivot + 1
		}
	}

	return firstBad
}

func isBadVersion(version int) bool {

	bad := testCases[expectedResult]
	if version >= bad {
		return true
	} else {
		return false
	}
}
