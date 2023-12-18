package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := readLines("input.txt")

	totalPoints := 0
	totalCards := map[int]int{}

	// initialize the first ownership of card
	for card, _ := range file {
		totalCards[card+1] = 1
	}

	for card, line := range file {

		splitCard := strings.Split(line, "|")

		// PLAYER_NUMBERS  SLICE
		playerSplit := strings.Split(splitCard[0], ": ")
		playerNumbers := strings.Split(playerSplit[1], " ")
		playerNumbersSlice := []int{}

		for _, v := range playerNumbers {
			if v != "" {
				n, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println(v)
					log.Fatalln("error converting number")
				}
				playerNumbersSlice = append(playerNumbersSlice, n)
			}
		}

		// for _, n := range playerNumbersSlice {
		// 	fmt.Println(n)
		// }

		// WINNING NUMBERS SLICE
		winningNumbers := strings.Split(splitCard[1], " ")
		winningNumbersSlice := []int{}

		for _, v := range winningNumbers {
			if v != "" {
				n, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println(v)
					log.Fatalln("error converting number")
				}
				winningNumbersSlice = append(winningNumbersSlice, n)
			}
		}

		// Calculating Points for Part 1
		linePoints := 0
		lineMatches := 0
		for _, w := range winningNumbersSlice {
			for _, p := range playerNumbersSlice {
				if p == w {
					if linePoints > 1 {
						linePoints = linePoints * 2
					} else {
						linePoints++
					}
					lineMatches++
				}
			}
		}

		// Calculating extra cards for part 2
		// start at pos 1 as the current card dont get extra
		if card == 0 {
			for i := 1; i <= lineMatches; i++ {

				if card < len(file) { // stop appending if past size
					totalCards[card+1+i] += 1
				}

			}
		} else {
			for j := 0; j < totalCards[card+1]; j++ {
				for k := 1; k <= lineMatches; k++ {
					if card < len(file) { // stop appending if past size
						totalCards[card+1+k] += 1
					}
				}
			}
		}

		// fmt.Println("card", card+1, "points", linePoints)
		// fmt.Println("card", card+1, "mathinc numbers", lineMatches)
		totalPoints += linePoints
	}

	totalSumCards := 0
	for card, _ := range file {
		totalSumCards += totalCards[card+1]
	}
	fmt.Println("Part 1 total points:", totalPoints)
	fmt.Println("Part 2 total cards:", totalSumCards)

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
