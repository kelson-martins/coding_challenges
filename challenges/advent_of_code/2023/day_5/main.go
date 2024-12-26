package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SEED_TO_SOIL = "seed-to-soil"
const SOIL_TO_FERTILIZER = "soil-to-fertilizer"
const FERTILIZER_TO_WATER = "fertilizer-to-water"
const WATER_TO_LIGHT = "water-to-light"
const LIGHT_TO_TEMPERATURE = "light-to-temperature"
const TEMPERATURE_TO_HUMIDITY = "temperature-to-humidity"
const HUMIDITY_TO_LOCATION = "humidity-to-location"

type conversionMap struct {
	mapType          string
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
	rangeLength      int
}

func main() {

	lowestLocation := 0
	seeds := []int{1482445116, 339187393, 3210489476, 511905836, 42566461, 51849137, 256584102, 379575844, 3040181568, 139966026, 4018529087, 116808249, 2887351536, 89515778, 669731009, 806888490, 2369242654, 489923931, 2086168596, 82891253}
	// seeds := []int{79, 14, 55, 13}

	lines, _ := readLines("input.txt")

	seedToSoilMaps := loadMap(lines, SEED_TO_SOIL)
	soilToFertilizerMaps := loadMap(lines, SOIL_TO_FERTILIZER)
	fertilizerToWaterMaps := loadMap(lines, FERTILIZER_TO_WATER)
	waterToLightMaps := loadMap(lines, WATER_TO_LIGHT)
	lightToTemperatureMaps := loadMap(lines, LIGHT_TO_TEMPERATURE)
	temperatureToHumidityMaps := loadMap(lines, TEMPERATURE_TO_HUMIDITY)
	humidityToLocationMaps := loadMap(lines, HUMIDITY_TO_LOCATION)

	fmt.Println(seedToSoilMaps)
	fmt.Println(soilToFertilizerMaps)
	fmt.Println(fertilizerToWaterMaps)
	fmt.Println(waterToLightMaps)
	fmt.Println(lightToTemperatureMaps)
	fmt.Println(temperatureToHumidityMaps)
	fmt.Println(humidityToLocationMaps)

	// jumping 2 steps
	for i := 0; i < len(seeds); i += 2 {
		// reached the end
		if i+2 > len(seeds) {
			break
		}

		seedRoot := seeds[i]
		seedSteps := seeds[i+1]
		seedMax := seedRoot + seedSteps

		for seed := seedRoot; seed <= seedMax; seed++ {
			// fmt.Println("seed", seed, "seedStep", seedSteps)
			soilVal := 0
			// 1. soil
			foundSoil := false
			for _, soil := range seedToSoilMaps {
				if seed >= soil.sourceStart && seed <= soil.sourceEnd {
					soilVal = soil.destinationStart + (seed - soil.sourceStart)
					foundSoil = true
					break
				}
			}

			if foundSoil == false {
				soilVal = seed
			}
			// fmt.Println("Soil:", soilVal)

			// 2. fertilizing
			fertilizerVal := 0
			foundFertilizer := false
			for _, fertilizing := range soilToFertilizerMaps {
				if soilVal >= fertilizing.sourceStart && soilVal <= fertilizing.sourceEnd {
					fertilizerVal = fertilizing.destinationStart + (soilVal - fertilizing.sourceStart)
					foundFertilizer = true
					break
				}
			}

			if !foundFertilizer {
				fertilizerVal = soilVal
			}

			// fmt.Println(fertilizerVal)

			// 3. water
			waterVal := 0
			foundWater := false
			for _, water := range fertilizerToWaterMaps {
				if fertilizerVal >= water.sourceStart && fertilizerVal <= water.sourceEnd {
					waterVal = water.destinationStart + (fertilizerVal - water.sourceStart)
					foundWater = true
					break
				}
			}

			if !foundWater {
				waterVal = fertilizerVal
			}

			// fmt.Println("water:", waterVal)

			// 4. light
			lightVal := 0
			foundLight := false
			for _, light := range waterToLightMaps {
				if waterVal >= light.sourceStart && waterVal <= light.sourceEnd {
					lightVal = light.destinationStart + (waterVal - light.sourceStart)
					foundLight = true
					break
				}
			}

			if !foundLight {
				lightVal = waterVal
			}

			// fmt.Println("light:", lightVal)

			// 5. temperature
			temperatureVal := 0
			foundTemperature := false
			for _, temperature := range lightToTemperatureMaps {
				if lightVal >= temperature.sourceStart && lightVal <= temperature.sourceEnd {
					temperatureVal = temperature.destinationStart + (lightVal - temperature.sourceStart)
					foundTemperature = true
					break
				}
			}

			if !foundTemperature {
				temperatureVal = lightVal
			}

			// fmt.Println("temperature:", temperatureVal)

			// 6. humidity
			humidityVal := 0
			foundHumidity := false
			for _, humidity := range temperatureToHumidityMaps {
				if temperatureVal >= humidity.sourceStart && temperatureVal <= humidity.sourceEnd {
					humidityVal = humidity.destinationStart + (temperatureVal - humidity.sourceStart)
					foundHumidity = true
					break
				}
			}

			if !foundHumidity {
				humidityVal = temperatureVal
			}

			// fmt.Println("humidity:", humidityVal)

			// 7. location
			locationVal := 0
			foundLocation := false
			for _, location := range humidityToLocationMaps {
				if humidityVal >= location.sourceStart && humidityVal <= location.sourceEnd {
					locationVal = location.destinationStart + (humidityVal - location.sourceStart)
					foundLocation = true
					break
				}
			}

			if !foundLocation {
				locationVal = humidityVal
			}

			// fmt.Println("location:", locationVal)

			if i == 0 && lowestLocation == 0 { // adding lowestLocation == 0 as we want only the first initialization
				lowestLocation = locationVal
			}

			if locationVal < lowestLocation {
				// fmt.Println("new lowest location", "old location", lowestLocation, "new location", locationVal)
				lowestLocation = locationVal
			}
		}
	}

	fmt.Println("Lowest Location: ", lowestLocation)

}

func loadMap(lines []string, mapType string) []conversionMap {

	seedToSoilBlock := false

	listMap := []conversionMap{}

	for _, line := range lines {
		if strings.Contains(line, mapType) {
			seedToSoilBlock = true
			continue
		}

		// iterate until breakline if I am in the block of interest
		if seedToSoilBlock {
			if len(line) == 0 {
				// finished
				break
			}

			// line of interest
			mp := conversionMap{
				mapType: mapType,
			}

			splitLine := strings.Split(line, " ")
			mp.destinationStart, _ = strconv.Atoi(splitLine[0])
			mp.sourceStart, _ = strconv.Atoi(splitLine[1])
			mp.rangeLength, _ = strconv.Atoi(splitLine[2])
			mp.sourceEnd = mp.sourceStart + mp.rangeLength
			mp.destinationEnd = mp.destinationStart + mp.rangeLength

			listMap = append(listMap, mp)

		}
	}

	return listMap
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
