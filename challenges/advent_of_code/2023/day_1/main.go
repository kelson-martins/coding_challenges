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
		for _, line := range file {

			total += getLineValues(line)

		}

	}

	log.Println("Total: ", total)
}

func getLineValues(text string) int {

	start := 0
	end := len(text) - 1

	char1 := ""
	char2 := ""

	for _, character := range text {

		if string(char1) == "" {

			if unicode.IsNumber(character) {

				char1 = string(character)

			}
		}

		if string(char2) == "" {
			if unicode.IsNumber(rune(text[end])) {
				char2 = string(text[end])
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
