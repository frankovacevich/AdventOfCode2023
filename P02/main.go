package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Draw struct {
	R int
	G int
	B int
}

type Game struct {
	ID    int
	Draws []Draw
}

func main() {
	do2()
}

func do1() {
	games := parseInput()
	sum := 0

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	for _, game := range games {
		validGame := true
		for _, draw := range game.Draws {
			if draw.R > maxRed || draw.G > maxGreen || draw.B > maxBlue {
				validGame = false
				break
			}
		}

		if validGame {
			sum += game.ID
		}
	}

	fmt.Println(sum)
}

func do2() {
	games := parseInput()
	sumPowers := 0

	for _, game := range games {
		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, draw := range game.Draws {
			if draw.R > minRed {
				minRed = draw.R
			}
			if draw.G > minGreen {
				minGreen = draw.G
			}
			if draw.B > minBlue {
				minBlue = draw.B
			}
		}

		sumPowers += minRed * minGreen * minBlue
	}

	fmt.Println(sumPowers)
}

func parseInput() []Game {
	games := []Game{}

	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		newGame := Game{}
		split1 := strings.Split(string(line), ": ")
		split2 := strings.Split(string(split1[1]), "; ")

		// Get games
		gameID := strings.Replace(split1[0], "Game ", "", 1)
		newGame.ID, _ = strconv.Atoi(gameID)

		// Get draw
		for _, draw := range split2 {
			newDraw := Draw{}

			split3 := strings.Split(string(draw), ", ")
			for _, die := range split3 {
				split4 := strings.Split(string(die), " ")
				amount, _ := strconv.Atoi(split4[0])
				color := split4[1]

				switch color {
				case "red":
					newDraw.R = amount
				case "green":
					newDraw.G = amount
				case "blue":
					newDraw.B = amount
				}
			}

			newGame.Draws = append(newGame.Draws, newDraw)
		}

		games = append(games, newGame)
	}

	return games
}
