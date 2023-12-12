package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type History []int

func main() {
	do2()
}

func do1() {
	histories := parseInput()
	sum := 0
	for h := range histories {
		sum += determineNextValue(histories[h])
	}
	fmt.Println(sum)
}

func do2() {
	histories := parseInput()
	sum := 0
	for h := range histories {
		sum += determinePrevValue(histories[h])
	}
	fmt.Println(sum)
}

func determineNextValue(history History) int {
	sequences := []History{}
	sequences = append(sequences, history)

	// Find sequences
	for {
		lastSequence := sequences[len(sequences)-1]

		// last sequence is all zeros?
		allZeros := true
		for _, num := range lastSequence {
			if num != 0 {
				allZeros = false
				break
			}
		}
		if allZeros {
			break
		}

		newSequence := make(History, len(lastSequence)-1)
		for i := 1; i < len(lastSequence); i++ {
			newSequence[i-1] = lastSequence[i] - lastSequence[i-1]
		}

		sequences = append(sequences, newSequence)
	}

	// Complete sequences
	lastSequence := sequences[len(sequences)-1]
	lastSequence = append(lastSequence, 0)
	for i := len(sequences) - 2; i >= 0; i-- {
		sequence := sequences[i]
		nextSequence := sequences[i+1]
		newValue := nextSequence[len(nextSequence)-1] + sequence[len(sequence)-1]
		sequences[i] = append(sequence, newValue)
	}

	sequence := sequences[0]
	return sequence[len(sequence)-1]
}

func determinePrevValue(history History) int {
	sequences := []History{}
	sequences = append(sequences, history)

	// Find sequences
	for {
		lastSequence := sequences[len(sequences)-1]

		// last sequence is all zeros?
		allZeros := true
		for _, num := range lastSequence {
			if num != 0 {
				allZeros = false
				break
			}
		}
		if allZeros {
			break
		}

		newSequence := make(History, len(lastSequence)-1)
		for i := 1; i < len(lastSequence); i++ {
			newSequence[i-1] = lastSequence[i] - lastSequence[i-1]
		}

		sequences = append(sequences, newSequence)
	}

	// Complete sequences
	lastSequence := sequences[len(sequences)-1]
	lastSequence = append([]int{0}, lastSequence...)
	for i := len(sequences) - 2; i >= 0; i-- {
		sequence := sequences[i]
		nextSequence := sequences[i+1]
		newValue := sequence[0] - nextSequence[0]
		sequences[i] = append([]int{newValue}, sequence...)
	}

	sequence := sequences[0]
	return sequence[0]
}

func parseInput() []History {
	histories := []History{}
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		split := strings.Split(line, " ")
		history := History{}
		for _, str := range split {
			num, _ := strconv.Atoi(str)
			history = append(history, num)
		}
		histories = append(histories, history)
	}

	return histories
}
