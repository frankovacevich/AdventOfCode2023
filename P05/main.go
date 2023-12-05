package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var globalCache = make(map[*map[int][2]int]int)

type ProblemInput struct {
	seeds                []int
	seedToSoil           map[int][2]int
	soilToFertilizer     map[int][2]int
	fertilizerToWater    map[int][2]int
	waterToLight         map[int][2]int
	lightToTemperature   map[int][2]int
	temperatureToHumidiy map[int][2]int
	humidityToLocation   map[int][2]int
}

func sliceStringToInt(slice []string) []int {
	newSlice := make([]int, len(slice))
	for i, item := range slice {
		value, _ := strconv.Atoi(item)
		newSlice[i] = value
	}
	return newSlice
}

func sliceStringContains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func fillMap(dest int, source int, length int, pmap *map[int][2]int) {
	value := [2]int{dest, length}
	(*pmap)[source] = value
}

func getValueFromMap(key int, pmap *map[int][2]int) int {
	cachedKey, exists := globalCache[pmap]
	if exists {
		source := cachedKey
		dest := (*pmap)[cachedKey][0]
		length := (*pmap)[cachedKey][1]
		if source <= key && source+length > key {
			return key - source + dest
		}
	}

	for source, value := range *pmap {
		dest := value[0]
		length := value[1]
		if source <= key && source+length > key {
			globalCache[pmap] = source
			return key - source + dest
		}
	}
	return key
}

func main() {
	do2()
}

func do1() {
	problemInput := parseInput()
	minLocation := 99999999999999999
	for _, seed := range problemInput.seeds {
		location := getLocation(seed, &problemInput)
		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}

func do2() {
	problemInput := parseInput()
	minLocation := 999999999999999999

	c := 0
	for i := 0; i < len(problemInput.seeds); i += 2 {
		start := problemInput.seeds[i]
		length := problemInput.seeds[i+1]
		for j := 0; j < length; j++ {

			seed := start + j
			location := getLocation(seed, &problemInput)
			if location < minLocation {
				minLocation = location
			}
			c += 1
		}
	}

	fmt.Println(minLocation)
}

func getLocation(seed int, problemInput *ProblemInput) int {
	x := seed
	x = getValueFromMap(x, &(*problemInput).seedToSoil)
	x = getValueFromMap(x, &(*problemInput).soilToFertilizer)
	x = getValueFromMap(x, &(*problemInput).fertilizerToWater)
	x = getValueFromMap(x, &(*problemInput).waterToLight)
	x = getValueFromMap(x, &(*problemInput).lightToTemperature)
	x = getValueFromMap(x, &(*problemInput).temperatureToHumidiy)
	x = getValueFromMap(x, &(*problemInput).humidityToLocation)
	return x
}

func parseInput() ProblemInput {
	problemInput := ProblemInput{}
	problemInput.seedToSoil = map[int][2]int{}
	problemInput.soilToFertilizer = map[int][2]int{}
	problemInput.fertilizerToWater = map[int][2]int{}
	problemInput.waterToLight = map[int][2]int{}
	problemInput.lightToTemperature = map[int][2]int{}
	problemInput.temperatureToHumidiy = map[int][2]int{}
	problemInput.humidityToLocation = map[int][2]int{}

	categories := []string{
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	}

	// Build maps
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	problemInput.seeds = sliceStringToInt(strings.Split(lines[0][7:], " "))

	category := ""
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if sliceStringContains(categories, line[:len(line)-1]) {
			category = line[:len(line)-1]
			continue
		}

		split1 := sliceStringToInt(strings.Split(line, " "))
		dest := split1[0]
		source := split1[1]
		length := split1[2]

		if category == categories[0] {
			fillMap(dest, source, length, &problemInput.seedToSoil)
		}
		if category == categories[1] {
			fillMap(dest, source, length, &problemInput.soilToFertilizer)
		}
		if category == categories[2] {
			fillMap(dest, source, length, &problemInput.fertilizerToWater)
		}
		if category == categories[3] {
			fillMap(dest, source, length, &problemInput.waterToLight)
		}
		if category == categories[4] {
			fillMap(dest, source, length, &problemInput.lightToTemperature)
		}
		if category == categories[5] {
			fillMap(dest, source, length, &problemInput.temperatureToHumidiy)
		}
		if category == categories[6] {
			fillMap(dest, source, length, &problemInput.humidityToLocation)
		}
	}

	return problemInput
}
