package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

type DirVector struct {
	vert int
	hor  int
}

const (
	Up Direction = iota
	Right
	Down
	Left
)

var Dir = [4]DirVector{
	Up:    {vert: -1, hor: 0}, // Up is (-1, 0)
	Right: {vert: 0, hor: 1},  // Right is (0, 1)
	Down:  {vert: 1, hor: 0},  // Down is (1, 0)
	Left:  {vert: 0, hor: -1}, // Left is (0, -1)
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

	var col, row int
	for i, line := range lines {
		for j, char := range line {
			if char == '^' {
				col = j
				row = i
			}
		}
	}

	direction := 0

	path := make([][]int, height)
	for i := range path {
		// Each inner slice represents a row
		path[i] = make([]int, width)
	}

	path[row][col] = 1

	for true {
		nextCol := col + Dir[direction].hor
		nextRow := row + Dir[direction].vert
		if nextCol < 0 || nextCol >= width || nextRow < 0 || nextRow >= height {
			break
		}
		if lines[nextRow][nextCol] == '#' {
			// fmt.Println("Colision")
			direction = (direction + 1) % 4
			nextCol = col + Dir[direction].hor
			nextRow = row + Dir[direction].vert
		}
		col = nextCol
		row = nextRow
		path[row][col] = 1
		// fmt.Println(col, row)
	}

	counter := 0
	for _, row := range path {
		for _, num := range row {
			if num != 0 {
				counter++
			}
		}
	}

	fmt.Println("Result: ", counter)

}
