package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
	label int
}

type HandList []Hand

func ParseInput() HandList {
	hands := HandList{}

	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		split1 := strings.Split(line, " ")
		cards := split1[0]
		bid, _ := strconv.Atoi(split1[1])

		hand := Hand{
			cards: cards,
			bid:   bid,
		}

		hands = append(hands, hand)
	}

	return hands
}
