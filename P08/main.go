package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type NodeMap map[string][2]string // node : [left_node right_node]

func main() {
	do2()
}

func do1() {
	instructions, nodeMap := parseInput()

	node := "AAA"
	i := 0
	steps := 0
	for {
		// are we there ?
		if node == "ZZZ" {
			break
		}

		// need to repeat instructions ?
		if i >= len(instructions) {
			i = 0
		}

		// step
		if string(instructions[i]) == "L" {
			node = nodeMap[node][0]
		} else {
			node = nodeMap[node][1]
		}

		i++
		steps++
	}

	fmt.Println(steps)
}

func do22() {
	instructions, nodeMap := parseInput()

	// Find starting nodes
	nodes := []string{}
	for node := range nodeMap {
		if endsWith(node, "A") {
			nodes = append(nodes, node)
		}
	}

	i := 0
	steps := 0
	for {
		// are we there ?
		areWeThere := true
		for n := range nodes {
			if !endsWith(nodes[n], "Z") {
				areWeThere = false
				break
			}
		}
		if areWeThere {
			break
		}

		// need to repeat instructions ?
		if i >= len(instructions) {
			i = 0
		}

		// step
		for n := range nodes {
			if string(instructions[i]) == "L" {
				nodes[n] = nodeMap[nodes[n]][0]
			} else {
				nodes[n] = nodeMap[nodes[n]][1]
			}
		}

		i++
		steps++
	}

	fmt.Println("DONE", steps)
}

func do2() {
	instructions, nodeMap := parseInput()

	// Find starting nodes
	nodes := []string{}
	for node := range nodeMap {
		if endsWith(node, "A") {
			nodes = append(nodes, node)
		}
	}

	results := []int{}
	for n := range nodes {
		i := 0
		steps := 0

		for {
			// are we there ?
			if endsWith(nodes[n], "Z") {
				break
			}

			// need to repeat instructions ?
			if i >= len(instructions) {
				i = 0
			}

			// step
			if string(instructions[i]) == "L" {
				nodes[n] = nodeMap[nodes[n]][0]
			} else {
				nodes[n] = nodeMap[nodes[n]][1]
			}

			i++
			steps++
		}
		results = append(results, steps)
	}

	least_common_multiple := LCM(1, 1, results...)
	fmt.Println(least_common_multiple)
}

func endsWith(node string, letter string) bool {
	lastLetter := node[len(node)-1:]
	if string(lastLetter) == letter {
		return true
	}
	return false
}

func parseInput() (string, NodeMap) {
	nodeMap := NodeMap{}

	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	instructions := lines[0]

	for _, line := range lines[2:] {
		split1 := strings.Split(line, " = ")
		split2 := strings.Split(split1[1], ", ")

		nodeName := split1[0]
		leftNode := split2[0][1:]
		rightNode := split2[1][:len(split2[1])-1]

		nodeMap[nodeName] = [2]string{leftNode, rightNode}
	}

	return instructions, nodeMap
}
