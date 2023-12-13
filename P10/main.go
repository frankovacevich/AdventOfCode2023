package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const PIPE_VERTICAL = "|"
const PIPE_HORIZONTAL = "-"
const PIPE_N_E = "L"
const PIPE_N_W = "J"
const PIPE_S_W = "7"
const PIPE_S_E = "F"
const PIPE_START = "S"
const GROUND = "."

type Tile struct {
	x        int
	y        int
	value    string
	inLoop   bool
	enclosed bool
}

func getTile(tiles *[][]Tile, x int, y int) *Tile {
	if x < 0 || x >= len((*tiles)[0]) {
		return nil
	}
	if y < 0 || y >= len((*tiles)) {
		return nil
	}
	return &(*tiles)[y][x]
}

func tileSetEnclosed(tiles *[][]Tile, x int, y int) {
	tile := getTile(tiles, x, y)
	if tile != nil && !tile.inLoop {
		tile.enclosed = true
	}
}

func isSameTile(tile1 *Tile, tile2 *Tile) bool {
	return tile1.x == tile2.x && tile1.y == tile2.y
}

func main() {
	do2()
}

func do1() {
	tiles, tileStart := parseInput()

	loop := *findLoop(&tiles, &tileStart)
	for i := range loop {
		getTile(&tiles, loop[i].x, loop[i].y).inLoop = true
	}

	fmt.Println(len(loop) / 2)
	// printTiles(&tiles, false)
}

func do2() {
	tiles, tileStart := parseInput()

	loop := *findLoop(&tiles, &tileStart)
	for i := range loop {
		getTile(&tiles, loop[i].x, loop[i].y).inLoop = true
	}

	iPrev := 0
	for i := range loop {
		if i == 0 {
			continue
		}

		x := loop[i].x
		y := loop[i].y
		deltaX := x - loop[iPrev].x
		deltaY := y - loop[iPrev].y
		iPrev = i

		tileSetEnclosed(&tiles, x+deltaY, y-deltaX)
		tileSetEnclosed(&tiles, x+deltaY-deltaX, y-deltaX-deltaY)
	}

	for y := range tiles {
		shouldBeEnclosed := false
		for x := range tiles[y] {
			tile := getTile(&tiles, x, y)
			if tile.inLoop {
				shouldBeEnclosed = false
			} else if tile.enclosed {
				shouldBeEnclosed = true
			}

			if shouldBeEnclosed {
				tile.enclosed = true
			}
		}
	}

	sum := 0
	for y := range tiles {
		for x := range tiles[y] {
			tile := getTile(&tiles, x, y)
			if tile.enclosed {
				sum++
			}
		}
	}

	fmt.Println(sum)

	//printTiles(&tiles, true)

}

func printTiles(tiles *[][]Tile, showEnclosed bool) {
	fmt.Println()
	fmt.Println()

	for y := range *tiles {
		for x := range (*tiles)[y] {
			tile := *getTile(tiles, x, y)
			if tile.inLoop {
				fmt.Print(tile.value)
			} else if showEnclosed && tile.enclosed {
				fmt.Print("⭕")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func findLoop(tiles *[][]Tile, tileStart *Tile) *[]Tile {
	tileNorth := getTile(tiles, tileStart.x, tileStart.y-1)
	tileWest := getTile(tiles, tileStart.x-1, tileStart.y)
	tileSouth := getTile(tiles, tileStart.x, tileStart.y+1)
	tileEast := getTile(tiles, tileStart.x+1, tileStart.y)

	tilesToCheck := []*Tile{}
	if tileNorth.value == PIPE_VERTICAL || tileNorth.value == PIPE_S_E || tileNorth.value == PIPE_S_W {
		tilesToCheck = append(tilesToCheck, tileNorth)
	}
	if tileWest.value == PIPE_HORIZONTAL || tileWest.value == PIPE_S_E || tileWest.value == PIPE_N_E {
		tilesToCheck = append(tilesToCheck, tileWest)
	}
	if tileSouth.value == PIPE_VERTICAL || tileSouth.value == PIPE_N_E || tileSouth.value == PIPE_N_W {
		tilesToCheck = append(tilesToCheck, tileSouth)
	}
	if tileEast.value == PIPE_HORIZONTAL || tileEast.value == PIPE_S_W || tileEast.value == PIPE_N_W {
		tilesToCheck = append(tilesToCheck, tileEast)
	}

	for i := range tilesToCheck {
		loop := []Tile{*tileStart}
		prevTile := tileStart
		tile := tilesToCheck[i]

		for {
			if tile == nil {
				break
			} else if tile.value == GROUND {
				break
			} else if tile.value == PIPE_START {
				return &loop
			}

			loop = append(loop, *tile)
			prevTile, tile = tile, findNextTileInLoop(tiles, tile, prevTile)
		}
	}

	return nil
}

func findNextTileInLoop(tiles *[][]Tile, tile *Tile, prevTile *Tile) *Tile {
	possibleNextTiles := [2]*Tile{}
	if tile.value == PIPE_HORIZONTAL {
		possibleNextTiles[0] = getTile(tiles, tile.x+1, tile.y)
		possibleNextTiles[1] = getTile(tiles, tile.x-1, tile.y)
	} else if tile.value == PIPE_VERTICAL {
		possibleNextTiles[0] = getTile(tiles, tile.x, tile.y+1)
		possibleNextTiles[1] = getTile(tiles, tile.x, tile.y-1)
	} else if tile.value == PIPE_N_W {
		possibleNextTiles[0] = getTile(tiles, tile.x, tile.y-1)
		possibleNextTiles[1] = getTile(tiles, tile.x-1, tile.y)
	} else if tile.value == PIPE_N_E {
		possibleNextTiles[0] = getTile(tiles, tile.x, tile.y-1)
		possibleNextTiles[1] = getTile(tiles, tile.x+1, tile.y)
	} else if tile.value == PIPE_S_W {
		possibleNextTiles[0] = getTile(tiles, tile.x, tile.y+1)
		possibleNextTiles[1] = getTile(tiles, tile.x-1, tile.y)
	} else if tile.value == PIPE_S_E {
		possibleNextTiles[0] = getTile(tiles, tile.x, tile.y+1)
		possibleNextTiles[1] = getTile(tiles, tile.x+1, tile.y)
	}

	var nextTile *Tile = nil
	if possibleNextTiles[0] != nil && !isSameTile(possibleNextTiles[0], prevTile) {
		nextTile = possibleNextTiles[0]
	} else if possibleNextTiles[1] != nil && !isSameTile(possibleNextTiles[1], prevTile) {
		nextTile = possibleNextTiles[1]
	}

	return nextTile
}

func parseInput() (tiles [][]Tile, startingTile Tile) {
	tiles = [][]Tile{}
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for y, line := range lines {
		row := []Tile{}
		for x, letter := range line {
			tile := Tile{x: x, y: y, value: string(letter)}
			row = append(row, tile)
			if tile.value == PIPE_START {
				startingTile = tile
			}
		}
		tiles = append(tiles, row)
	}

	return tiles, startingTile
}
