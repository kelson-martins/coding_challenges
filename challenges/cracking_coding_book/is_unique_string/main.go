package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string

	fmt.Print("Enter string: ")
	fmt.Scanln(&userInput)

	hasUniqueCharsHashMap(userInput)
	hasUniqueCharsArray(userInput)

}

func hasUniqueCharsHashMap(str string) {

	strMap := map[string]bool{}

	for _, v := range strings.Split(str, "") {
		if !strMap[v] {
			strMap[v] = true
		} else {
			fmt.Println("string does not have unique characters")
			return
		}
	}

	fmt.Println("string has all unique characters")
}

func hasUniqueCharsArray(str string) {

	mem := []string{}

	for _, c := range strings.Split(str, "") {

		for _, m := range mem {
			if m == c {
				fmt.Println("string does not have unique characters")
				os.Exit(0)
			}
		}
		mem = append(mem, c)
	}

	fmt.Println("string has all unique characters")

}
