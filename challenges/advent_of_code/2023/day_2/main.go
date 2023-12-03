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

	allGames := []map[string]int{}

	games, err := readLines("input.txt")
	if err == nil {

		start := 1

		for _, game := range games {
			gameMap := getGameDetails(game)
			allGames = append(allGames, gameMap)

			fmt.Println(gameMap)

			start++
		}

		fmt.Println(checkElfBag(allGames))
	}

}

func checkElfBag(gameMaps []map[string]int) (int, int) {

	red := 12
	green := 13
	blue := 14

	idSum := 0
	powerSum := 0

	for _, game := range gameMaps {
		redOk := game["red"] <= red
		blueOk := game["blue"] <= blue
		greenOk := game["green"] <= green

		if redOk && blueOk && greenOk {
			idSum += game["id"]
		}

		powerSum += game["red"] * game["blue"] * game["green"]
	}

	return idSum, powerSum
}

func getGameDetails(line string) map[string]int {

	gameMap := make(map[string]int)

	gameMap["blue"] = 0
	gameMap["green"] = 0
	gameMap["red"] = 0

	// pos [0] will contain the Game ID
	// post [1] will contain the cubes
	gameSlice := strings.Split(line, ":")

	gameID := strings.Split(gameSlice[0], " ")[1]
	idAsInt, err := strconv.Atoi(gameID)
	if err != nil {
		log.Fatalln("Error converting String to Int")
	}
	gameMap["id"] = idAsInt

	// now taking care of the cubes
	cubeLine := strings.Replace(gameSlice[1], ",", ";", -1)
	cubeSplit := strings.Split(cubeLine, ";")

	for _, cube := range cubeSplit {

		// cube number will be position 1
		// cube color will be position 2
		cubePositions := strings.Split(cube, " ")

		cubeColor := cubePositions[2]

		cubeNumber, err := strconv.Atoi(cubePositions[1])
		if err != nil {
			log.Fatalln("Error decoding cube value to integer")
		}

		if cubeColor == "blue" {

			if gameMap["blue"] < cubeNumber {
				gameMap["blue"] = cubeNumber
			}

		} else if cubeColor == "green" {
			if gameMap["green"] < cubeNumber {
				gameMap["green"] = cubeNumber
			}

		} else if cubeColor == "red" {
			if gameMap["red"] < cubeNumber {
				gameMap["red"] = cubeNumber
			}
		} else {
			log.Fatalln("Error getting any cube color number")
		}

	}

	return gameMap
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
