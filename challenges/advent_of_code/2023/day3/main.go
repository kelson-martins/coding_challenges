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

	file, err := readLines("input.txt")

	partNumbers := []string{}
	sequencesPerLine := map[int][][]int{}

	if err == nil {

		for k, line := range file {

			linesToProcess := []string{}

			if k == 0 {
				// current line
				linesToProcess = append(linesToProcess, line)
				// current line + 1
				linesToProcess = append(linesToProcess, file[k+1])
			} else if k == len(file)-1 {
				// current line
				linesToProcess = append(linesToProcess, line)
				// current line - 1
				linesToProcess = append(linesToProcess, file[k-1])
			} else {
				// 3 lines. before, current and after

				// current line
				linesToProcess = append(linesToProcess, line)
				// current line - 1
				linesToProcess = append(linesToProcess, file[k-1])
				// current line + 1
				linesToProcess = append(linesToProcess, file[k+1])
			}

			// process for part numbers
			parts, sequences := processLine(linesToProcess)
			partNumbers = append(partNumbers, parts...)
			sequencesPerLine[k] = sequences

		}
	}

	// fmt.Println("part numbers:",partNumbers)

	sum := 0

	for _, v := range partNumbers {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("error parsing string to int")
		}

		sum += val
	}

	fmt.Println("sum of part numbers: ", sum)

	// computing the gears

	for k, line := range file {
		findGears(line, k, sequencesPerLine, file)
	}

}

// need to process only the first incoming line
func findGears(line string, currentLine int, mapa map[int][][]int, file []string) {

	for j := 0; j < len(line); j++ {

		char := rune(line[j])

		// find the number to start the lookup
		if checkStar(char) {
			numberAttachedParts := 0
			gears := []string{}
			// found star, now needs to find adjascent

			// 1. check same line left
			if j > 0 {
				leftChar := rune(line[j-1])
				if unicode.IsNumber(leftChar) {
					numberAttachedParts += 1

				}
			}
			// 1. check same line right
			if j+1 < len(line) {
				rightChar := rune(line[j+1])
				if unicode.IsNumber(rightChar) {
					numberAttachedParts += 1
				}
			}

			// check previous line
			content, ok := mapa[currentLine-1]
			if ok {
				for _, nums := range content {
					if j+1 >= nums[0] && j-1 <= nums[1] {
						numberAttachedParts += 1
						gears = append(gears, file[currentLine-1][nums[0]:nums[1]+1])
						// fmt.Println("gear value", file[currentLine-1][nums[0]:nums[1]])
					}
				}
			}
			// check next line
			content, ok = mapa[currentLine+1]
			if ok {
				for _, nums := range content {
					if j+1 >= nums[0] && j-1 <= nums[1] {
						numberAttachedParts += 1
						gears = append(gears, file[currentLine+1][nums[0]:nums[1]+1])
						fmt.Println("gear value", file[currentLine+1][nums[0]:nums[1]+1])
					}
				}
			}

			if numberAttachedParts == 2 {
				fmt.Println(gears)
			}
		}
	}

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

// process 2 or 3 lines at a time
func processLine(lines []string) ([]string, [][]int) {

	partNumbers := []string{}
	numberSequences := [][]int{}

	for _, line := range lines {

		for j := 0; j < len(line); j++ {

			startSequence := 0
			endSequence := 0

			char := rune(line[j])

			// find the number to start the lookup
			if unicode.IsNumber(char) {

				startSequence = j

				// finding the end of the sequence
				// TODO see if start of sequence dont already end the line
				for i := startSequence + 1; i <= len(line)-1; i++ {
					c := line[i]

					if !unicode.IsNumber(rune(c)) {
						endSequence = i - 1
						j = endSequence
						break
					}

					// if reached the end, asumme last char is the end
					if i == len(line)-1 {
						endSequence = i
						continue
					}

				}

				// once here it means we found the ending number sequence
				// now check if they are part numbers
				// also check gears

				foundPart := false // if found any single adjacente, move to true to avoid adding multiple parts inc count

				// check first the left most character
				hasLeftChar := startSequence-1 >= 0
				if hasLeftChar {
					leftChar := rune(line[startSequence-1])
					if checkSymbol(leftChar) {
						fmt.Println("Sequence for left char", startSequence, endSequence)

						foundPart = true
						partNumbers = append(partNumbers, line[startSequence:endSequence+1])

					}

				}

				// check right most char
				hasRightChar := endSequence+1 < len(line)
				if hasRightChar {
					rightChar := rune(line[endSequence+1])
					if checkSymbol(rightChar) {
						fmt.Println("Sequence for rigjt char", startSequence, endSequence)

						if foundPart == false {
							foundPart = true
							partNumbers = append(partNumbers, line[startSequence:endSequence+1])
						}
					}

				}

				// check adjascent lines
				for z := 1; z < len(lines); z++ {
					matchLine := checkParentLine(lines[z], startSequence, endSequence)
					if matchLine {
						if foundPart == false {
							partNumbers = append(partNumbers, line[startSequence:endSequence+1])
						}
					}

				}

				// adding the number range to the map
				seq := []int{startSequence, endSequence}
				numberSequences = append(numberSequences, seq)
			}

		}

		// break after the first iteration as we only want to process the first incoming line
		break
	}

	return partNumbers, numberSequences

}

func checkParentLine(text string, start int, end int) bool {

	foundSymbol := false

	for i := start; i <= end; i++ {
		char := rune(text[i])
		if checkSymbol(char) {
			foundSymbol = true

		}
	}

	// checkDiagonal
	if start-1 >= 0 {
		char := rune(text[start-1])
		if checkSymbol(char) {
			foundSymbol = true

		}
	}

	if end+1 < len(text) {
		char := rune(text[end+1])
		if checkSymbol(char) {
			foundSymbol = true

		}
	}

	return foundSymbol
}

func checkStar(c rune) bool {

	if c == '*' {
		return true
	}

	return false
}

func checkSymbol(c rune) bool {

	symbols := []rune{'#', '*', '=', '&', '/', '-', '$', '@', '%', '+'}

	for _, symbol := range symbols {
		if c == symbol {
			return true
		}
	}

	return false
}
