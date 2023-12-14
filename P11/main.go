package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tile struct {
	x        int
	y        int
	isGalaxy bool
}

func main() {
	// do(1) // part 1
	do(1000000) // part 1
}

func do(expandingFactor int) {
	space := parseInput()
	sum := 0
	galaxyPairs := 0

	// expand
	space1 := *expandRows(&space, expandingFactor)

	// count
	galaxies := findGalaxies(&space1)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxy1 := galaxies[i]
			galaxy2 := galaxies[j]
			galaxyPairs++

			if galaxy1.y > galaxy2.y {
				sum += galaxy1.y - galaxy2.y
			} else {
				sum += galaxy2.y - galaxy1.y
			}

		}
	}
	fmt.Println("Galaxy pairs", galaxyPairs)
	fmt.Println("Sum (partial)", sum)

	// transpose and expand again
	space2 := *transposeSpace(&space)
	space2 = *expandRows(&space2, expandingFactor)

	// count
	galaxies = findGalaxies(&space2)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxy1 := galaxies[i]
			galaxy2 := galaxies[j]
			galaxyPairs++

			if galaxy1.y > galaxy2.y {
				sum += galaxy1.y - galaxy2.y
			} else {
				sum += galaxy2.y - galaxy1.y
			}

		}
	}

	// print
	fmt.Println(sum)

	// printSpace(&space)
}

func expandRows(space *[][]Tile, mult int) *[][]Tile {
	finalSize := 0

	for y := range *space {
		row := (*space)[y]
		finalSize++

		rowHasNoGalaxies := true
		for x := range row {
			tile := row[x]
			if tile.isGalaxy {
				rowHasNoGalaxies = false
				break
			}
		}
		if rowHasNoGalaxies {
			finalSize += mult - 1
		}
	}

	expandedSpace := make([][]Tile, finalSize)

	i := 0
	for y := range *space {
		row := (*space)[y]
		expandedSpace[i] = row
		i++

		rowHasNoGalaxies := true
		for x := range row {
			tile := row[x]
			if tile.isGalaxy {
				rowHasNoGalaxies = false
				break
			}
		}

		if rowHasNoGalaxies {
			for m := 1; m < mult; m++ {
				expandedSpace[i] = row
				i++
			}
		}
	}

	return &expandedSpace
}

func transposeSpace(space *[][]Tile) *[][]Tile {
	rowCount := len(*space)
	colCount := len((*space)[0])

	transposed := make([][]Tile, colCount)
	for c := 0; c < colCount; c++ {
		transposed[c] = make([]Tile, rowCount)
		for r := 0; r < rowCount; r++ {
			transposed[c][r] = (*space)[r][c]
		}
	}
	return &transposed
}

func findGalaxies(space *[][]Tile) []Tile {
	galaxies := []Tile{}
	for y := range *space {
		for x := range (*space)[y] {
			tile := (*space)[y][x]
			if tile.isGalaxy {
				tile.x = x
				tile.y = y
				galaxies = append(galaxies, tile)
			}
		}
	}
	return galaxies
}

func printSpace(space *[][]Tile) {
	for y := range *space {
		for x := range (*space)[y] {
			tile := (*space)[y][x]
			if tile.isGalaxy {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseInput() [][]Tile {
	tiles := [][]Tile{}
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for y, line := range lines {
		row := []Tile{}
		for x, letter := range line {
			tile := Tile{x: x, y: y, isGalaxy: string(letter) == "#"}
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}
	return tiles
}
