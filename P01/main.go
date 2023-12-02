package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// process_string_part_2("7pqrstsixteen")
	do()
}

func do() {
	strs := load_strings("input.txt")
	sum := 0
	for _, str := range strs {
		// v := process_string_part_1(str)
		v := process_string_part_2(str)
		sum += v
		fmt.Println(str, v)
	}
	fmt.Println(sum)
}

func load_strings(file string) []string {
	content, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(content), "\n")
	return lines
}

func process_string_part_1(str string) int {
	var num1 string
	var num2 string

	// Check first char
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		_, err := strconv.Atoi(char)
		if err == nil {
			num1 = char
			break
		}
	}

	// Check last char
	for i := len(str) - 1; i >= 0; i-- {
		char := string(str[i])
		_, err := strconv.Atoi(char)
		if err == nil {
			num2 = char
			break
		}
	}

	result, _ := strconv.Atoi(num1 + num2)
	return result
}

func process_string_part_2(str string) int {
	var numbers []int

	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	i := 0
	for {
		if i >= len(str) {
			break
		}

		// Check digit
		num, err := strconv.Atoi(string(str[i]))
		if err == nil {
			numbers = append(numbers, num)
			i++
			continue
		}

		// Check literal
		for literal, digit := range digits {
			if len(literal)+i > len(str) {
				continue
			}

			match := true
			for j := range literal {
				if literal[j] != str[i+j] {
					match = false
					break
				}
			}

			if match {
				numbers = append(numbers, digit)
				i += 1
				continue
			}
		}

		// continue
		i++
	}

	num1 := strconv.Itoa(numbers[0])
	num2 := strconv.Itoa(numbers[len(numbers)-1])
	result, _ := strconv.Atoi(num1 + num2)
	return result
}
