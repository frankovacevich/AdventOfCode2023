package main

import "fmt"

func main() {
	do2()
}

func do2() {
	data := map[int]int{
		49787980: 298118510661181,
	}
	do(data)
}

func do1() {
	data := map[int]int{
		49: 298,
		78: 1185,
		79: 1066,
		80: 1181,
	}
	do(data)
}

func do(data map[int]int) {
	waysOfWinning := []int{}

	for raceTime, recordDistance := range data {
		count := 0
		for v := 0; v <= raceTime; v++ {
			if v*(raceTime-v) > recordDistance {
				count += 1
			}
		}
		waysOfWinning = append(waysOfWinning, count)
	}

	product := 1
	for _, c := range waysOfWinning {
		product = product * c
	}

	fmt.Println(product)
}
