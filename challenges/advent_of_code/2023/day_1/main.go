package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := readLines("/Users/kelson/drive/repository/personal/coding_challenges/challenges/advent_of_code/2023/day_1/input.txt")
	if err != nil {
		for k, v := range f {
			fmt.Println(k, v)
		}
	}
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
