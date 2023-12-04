package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	do2()
}

func do1() {
	parts := map[int]int{}
	partsSum := 0

	// Load input
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	// For each line, find symbols and then parts
	for i, line := range lines {
		for j, char := range line {
			if !isSymbol(string(char)) {
				continue
			}

			var num int
			var err error

			// Look on previous line
			if i > 0 {
				previousLine := lines[i-1]
				num, err = findNumber(previousLine, j)
				if err == nil {
					parts[num]++
					partsSum += num
				} else {
					num, err = findNumber(previousLine, j-1)
					if err == nil {
						parts[num]++
						partsSum += num
					}
					num, err = findNumber(previousLine, j+1)
					if err == nil {
						parts[num]++
						partsSum += num
					}
				}
			}

			// Look on this line
			num, err = findNumber(line, j-1)
			if err == nil {
				parts[num]++
				partsSum += num
			}
			num, err = findNumber(line, j+1)
			if err == nil {
				parts[num]++
				partsSum += num
			}

			// Look on next line
			if i < len(lines)-1 {
				nextLine := lines[i+1]
				num, err = findNumber(nextLine, j)
				if err == nil {
					parts[num]++
					partsSum += num
				} else {
					num, err = findNumber(nextLine, j-1)
					if err == nil {
						parts[num]++
						partsSum += num
					}
					num, err = findNumber(nextLine, j+1)
					if err == nil {
						parts[num]++
						partsSum += num
					}
				}
			}
		}
	}

	fmt.Println(parts)
	fmt.Println()
	fmt.Println(partsSum)
}

func do2() {
	sum := 0

	// Load input
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	// For each line, find symbols and then parts
	for i, line := range lines {
		for j, char := range line {
			if string(char) != "*" {
				continue
			}

			var nums []int
			var num int
			var err error

			// Look on previous line
			if i > 0 {
				previousLine := lines[i-1]
				num, err = findNumber(previousLine, j)
				if err == nil {
					nums = append(nums, num)
				} else {
					num, err = findNumber(previousLine, j-1)
					if err == nil {
						nums = append(nums, num)
					}
					num, err = findNumber(previousLine, j+1)
					if err == nil {
						nums = append(nums, num)
					}
				}
			}

			// Look on this line
			num, err = findNumber(line, j-1)
			if err == nil {
				nums = append(nums, num)
			}
			num, err = findNumber(line, j+1)
			if err == nil {
				nums = append(nums, num)
			}

			// Look on next line
			if i < len(lines)-1 {
				nextLine := lines[i+1]
				num, err = findNumber(nextLine, j)
				if err == nil {
					nums = append(nums, num)
				} else {
					num, err = findNumber(nextLine, j-1)
					if err == nil {
						nums = append(nums, num)
					}
					num, err = findNumber(nextLine, j+1)
					if err == nil {
						nums = append(nums, num)
					}
				}
			}

			if len(nums) > 2 {
				println("WHAT!?")
			}

			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}

	fmt.Println(sum)
}

func isSymbol(char string) bool {
	// Is dot?
	if char == "." {
		return false
	}

	// Is number?
	_, err := strconv.Atoi(char)
	if err == nil {
		return false
	}

	return true
}

func findNumber(str string, position int) (int, error) {
	if position < 0 {
		return -1, errors.New("Position < 0")
	}
	if position >= len(str) {
		return -1, errors.New("Position >= len(str)")
	}

	// fmt.Println(string(str[position]))

	// Find start
	start := position
	for i := position; i >= 0; i-- {
		_, err := strconv.Atoi(string(str[i]))
		if err != nil {
			break
		}
		start = i
	}

	// Find end
	end := position
	for i := position; i < len(str); i++ {
		_, err := strconv.Atoi(string(str[i]))
		if err != nil {
			break
		}
		end = i + 1
	}

	// Return
	if start == end {
		return start, errors.New("No number found")
	}

	num, _ := strconv.Atoi(str[start:end])
	return num, nil
}
