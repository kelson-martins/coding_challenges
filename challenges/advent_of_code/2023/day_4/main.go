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

		linePoints := 0
		for _, w := range winningNumbersSlice {
			for _, p := range playerNumbersSlice {
				if p == w {
					if linePoints > 1 {
						linePoints = linePoints * 2
					} else {
						linePoints++
					}

				}
			}
		}

		// break
		fmt.Println("card", card+1, "points", linePoints)
		totalPoints += linePoints
	}

	fmt.Println(totalPoints)

}

func getPoints() {

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
