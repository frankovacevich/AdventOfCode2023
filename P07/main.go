package main

import (
	"fmt"
	"sort"
)

const Letters = "AKQJT98765432"

const FIVE_OF_A_KIND = 6
const FOUR_OF_A_KIND = 5
const FULL_HOUSE = 4
const THREE_OF_A_KIND = 3
const TWO_PAIR = 2
const ONE_PAIR = 1
const HIGH_CARD = 0

// sort
func (hands HandList) Len() int               { return len(hands) }
func (hands HandList) Swap(i int, j int)      { hands[i], hands[j] = hands[j], hands[i] }
func (hands HandList) Less(i int, j int) bool { return compareHands(&hands[i], &hands[j]) < 0 }

//

func main() {
	do1()
}

func do1() {
	hands := ParseInput()
	for i := range hands {
		labelHand(&hands[i])
	}
	sort.Sort(hands)

	// calculate winnings
	sum := 0
	for i, hand := range hands {
		rank := i + 1
		sum += rank * hand.bid
	}

	fmt.Println(sum)
}

func compareHands(hand1 *Hand, hand2 *Hand) int {
	// Returns 1, 0 or -1 if hand1 is greater, equal or less than hand2
	if hand1.label > hand2.label {
		return 1
	}
	if hand2.label > hand1.label {
		return -1
	}

	for i := range hand1.cards {
		p1 := letterPrecedence(hand1.cards[i])
		p2 := letterPrecedence(hand2.cards[i])
		if p1 < p2 {
			return 1
		}
		if p2 < p1 {
			return -1
		}
	}

	return 0
}

func letterPrecedence(letter byte) int {
	// the lower the better
	for i := range Letters {
		if letter == Letters[i] {
			return i
		}
	}
	return 99
}

func labelHand(hand *Hand) {
	letterCount := []int{}
	for _, letter := range Letters {
		count := 0
		for _, card := range hand.cards {
			if card == letter {
				count++
			}
		}
		if count > 1 {
			letterCount = append(letterCount, count)
		}
	}

	if len(letterCount) == 0 {
		(*hand).label = HIGH_CARD
	} else if len(letterCount) == 1 {
		if letterCount[0] == 2 {
			(*hand).label = ONE_PAIR
		} else if letterCount[0] == 3 {
			(*hand).label = THREE_OF_A_KIND
		} else if letterCount[0] == 4 {
			(*hand).label = FOUR_OF_A_KIND
		} else if letterCount[0] == 5 {
			(*hand).label = FIVE_OF_A_KIND
		}
	} else if len(letterCount) == 2 {
		if letterCount[0] == 2 && letterCount[1] == 3 {
			(*hand).label = FULL_HOUSE
		} else if letterCount[0] == 3 && letterCount[1] == 2 {
			(*hand).label = FULL_HOUSE
		} else if letterCount[0] == 2 && letterCount[1] == 2 {
			(*hand).label = TWO_PAIR
		}
	}
}
