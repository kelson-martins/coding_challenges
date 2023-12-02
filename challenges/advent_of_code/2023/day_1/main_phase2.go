package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	total := 0

	file, err := readLines("input.txt")
	if err == nil {
		for k, line := range file {

			if k == 3 {
				fmt.Println()
			}

			total += getLineValues(line)

			// if k == 3 {
			// 	os.Exit(1)
			// }

		}

	}

	log.Println("Total: ", total)
}

func getLineValues(text string) int {

	start := 0
	end := len(text) - 1

	char1 := ""
	char2 := ""

	spelledWith3 := []string{"one", "two", "six"}
	spelledWith4 := []string{"four", "five", "nine"}
	spelledWith5 := []string{"three", "seven", "eight"}

	for k, character := range text {

		if string(char1) == "" {
			if unicode.IsNumber(character) {

				char1 = string(character)

			}
		}

		if string(char1) == "" {

			// check for spelled 3 char 1
			if 3 <= len(text[k:]) {
				ahead3 := text[k : k+3]

				for _, option := range spelledWith3 {
					if ahead3 == option {
						if option == spelledWith3[0] {
							char1 = "1"
							break
						} else if option == spelledWith3[1] {
							char1 = "2"
							break
						} else {
							char1 = "6"
							break
						}
					}
				}
			}
		}
		if string(char1) == "" {
			// check for spelled 4 char1
			if 4 <= len(text[k:]) {
				ahead4 := text[k : k+4]

				for _, option := range spelledWith4 {
					if ahead4 == option {
						if option == spelledWith4[0] {
							char1 = "4"
							break
						} else if option == spelledWith4[1] {
							char1 = "5"
							break
						} else {
							char1 = "9"
							break
						}
					}
				}
			}
		}
		if string(char1) == "" {
			// check for spelled 5
			if 5 <= len(text[k:]) {
				ahead5 := text[k : k+5]

				for _, option := range spelledWith5 {
					if ahead5 == option {
						if option == spelledWith5[0] {
							char1 = "3"
							break
						} else if option == spelledWith5[1] {
							char1 = "7"
							break
						} else {
							char1 = "8"
							break
						}
					}
				}
			}

		}

		if string(char2) == "" {
			if unicode.IsNumber(rune(text[end])) {
				char2 = string(text[end])
			}
		}

		if string(char2) == "" {
			// check for spelled 3 char 2
			if 3 <= len(text[end:]) {
				ahead3 := text[end : end+3]

				for _, option := range spelledWith3 {
					if ahead3 == option {
						if option == spelledWith3[0] {
							char2 = "1"
							break
						} else if option == spelledWith3[1] {
							char2 = "2"
							break
						} else {
							char2 = "6"
							break
						}
					}
				}
			}
		}
		if string(char2) == "" {
			// check for spelled 4 char2
			if 4 <= len(text[end:]) {
				ahead4 := text[end : end+4]

				for _, option := range spelledWith4 {
					if ahead4 == option {
						if option == spelledWith4[0] {
							char2 = "4"
							break
						} else if option == spelledWith4[1] {
							char2 = "5"
							break
						} else {
							char2 = "9"
							break
						}
					}
				}
			}

			// check for spelled 5 char2
		}
		if string(char2) == "" {
			if 5 <= len(text[end:]) {
				ahead5 := text[end : end+5]

				for _, option := range spelledWith5 {
					if ahead5 == option {
						if option == spelledWith5[0] {
							char2 = "3"
							break
						} else if option == spelledWith5[1] {
							char2 = "7"
							break
						} else {
							char2 = "8"
							break
						}
					}
				}
			}
		}

		start++
		end--

	}

	combinedInt, err := strconv.Atoi(fmt.Sprintf("%v%v", char1, char2))
	if err != nil {
		log.Fatalln("Error converting String to Int")
	}
	fmt.Println(combinedInt)

	return combinedInt
}

// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
