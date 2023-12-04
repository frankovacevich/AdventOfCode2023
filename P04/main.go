package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Card struct {
	ID             int
	WinningNumbers []string
	PlayedNumbers  []string
}

func contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	do2()
}

func do1() {
	cards := parseInput()
	points := 0

	for _, card := range cards {
		cardPoints := 0
		for _, number := range card.PlayedNumbers {
			if contains(card.WinningNumbers, number) {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints = 2 * cardPoints
				}
			}
		}

		points += cardPoints
	}

	fmt.Println(points)
}

func do2() {
	cards := parseInput()
	cardsCount := map[int]int{}

	// Initialize cardsCount with 1 for each card
	for c := range cards {
		cardsCount[c] = 1
	}

	// Update count
	for c, card := range cards {

		// Get matches
		matches := 0
		for _, number := range card.PlayedNumbers {
			if contains(card.WinningNumbers, number) {
				matches++
			}
		}

		// Increase cards count
		for j := 0; j < matches; j++ {
			cardsCount[c+j+1] += cardsCount[c]
		}
	}

	// Add count
	sum := 0
	for _, count := range cardsCount {
		sum += count
	}
	fmt.Println(sum)
}

func parseInput() []Card {
	cards := []Card{}

	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	id := 1
	for _, line := range lines {
		line = strings.Replace(line, "  1", " 01", -1)
		line = strings.Replace(line, "  2", " 02", -1)
		line = strings.Replace(line, "  3", " 03", -1)
		line = strings.Replace(line, "  4", " 04", -1)
		line = strings.Replace(line, "  5", " 05", -1)
		line = strings.Replace(line, "  6", " 06", -1)
		line = strings.Replace(line, "  7", " 07", -1)
		line = strings.Replace(line, "  8", " 08", -1)
		line = strings.Replace(line, "  9", " 09", -1)
		line = line[10:]

		split1 := strings.Split(line, " | ")
		winningNumbers := strings.Split(split1[0], " ")
		playedNumbers := strings.Split(split1[1], " ")

		newCard := Card{ID: id, WinningNumbers: winningNumbers, PlayedNumbers: playedNumbers}
		cards = append(cards, newCard)
		id++
	}

	return cards
}
