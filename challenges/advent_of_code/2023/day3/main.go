package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := readLines("input.txt")
	if err == nil {

		start := 1
		for _, line := range file {

			start++

			if start == 5 {
				getParts(line)
				os.Exit(0)
			}

		}

	}
}

func getParts(line string) {
	fmt.Println(line)

	r, _ := regexp.Compile("([0-9]+[*$+#-])")
	fmt.Println(r.FindAllString(line, -1))
}

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
