package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

type DirVector struct {
	vert int
	hor int
}

type pathHistory struct {
	stepped bool
	direction Direction
}

const(
	Up Direction = iota
	Right
	Down
	Left
)

var Dir =  [4]DirVector{
	Up:    {vert: -1, hor: 0},  // Up is (-1, 0)
	Right: {vert: 0, hor: 1},   // Right is (0, 1)
	Down:  {vert: 1, hor: 0},   // Down is (1, 0)
	Left:  {vert: 0, hor: -1},  // Left is (0, -1)
}

func tryObstacle(m []string, obsCol, obsRow, startCol, startRow int) bool {
	if obsCol== startCol && obsRow== startRow {
		return false
	}
	height := len(m)
	width := len(m[0])
	path := make([][][]Direction, height)
	for i := range path {
		// Each inner slice represents a row
		path[i] = make([][]Direction, width)
	}

	direction := 0
	col := startCol
	row := startRow
	for !(col < 0 || col >= width || row < 0 || row >= height) {
		// Check if we entered a loop
		for _, dir := range path[row][col] {
			if dir == Direction(direction) {
				fmt.Println("loop if obstacle in ", obsCol, obsRow)
				return true
			}
		}
		path[row][col] = append(path[row][col], Direction(direction))

		if m[row][col] == '#' || (row == obsRow && col== obsCol){
			// fmt.Println("Colision")
			col = col - Dir[direction].hor
			row = row - Dir[direction].vert
			direction = (direction + 1) % 4
			continue
		}

		// Update position
		col += Dir[direction].hor
		row += Dir[direction].vert
		// fmt.Println(col, row)
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	// Read the entire file into memory
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(data)
	lines := strings.Split(text, "\n")
	lines = lines[:len(lines)-1]

	width := len(lines[0])
	height := len(lines)

	fmt.Println("width", width)
	fmt.Println("height", height) 



	var startCol,startRow int
	for i, line := range lines {
		for j, char := range line {
			if char == '^' {
				startCol = j
				startRow = i
			}
		}
	}

	counter := 0
	for i, line := range lines {
		for j, _:= range line {
			if tryObstacle(lines, j, i, startCol, startRow) {
				counter++
			}
		}
	}

	fmt.Println("Result: ", counter)

}
