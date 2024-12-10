package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	row    int
	col    int
	heigth int
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

	m := make([][]int, len(lines))
	for j, line := range lines {
		row := make([]int, len(lines[0]))
		for i := range line {
			row[i], _ = strconv.Atoi(string(line[i]))
		}
		m[j] = row
	}

	sum := 0
	for i, row := range m {
		for j, num := range row {
			if num == 0 {
				sum += trailHeadScore(i, j, m)
			}
		}
	}

	fmt.Println("Result: ", sum)

}

func trailHeadScore(row, col int, m [][]int) int {
	result := make(map[position]bool)
	path := make([]position, 0)

	backtrack(position{row, col, 0}, m, path, result)

	return len(result)

}

func backtrack(currentPos position, m [][]int, path []position, result map[position]bool) {
	// if solution append path to result
	if len(path) > 0 && (path[len(path)-1].heigth == 9) {
		// fmt.Println("Solucion encontrada")
		// fmt.Println(path[len(path)-1])
		result[path[len(path)-1]] = true
		// fmt.Println("Longitud local", len(result))
	}

	directions := []struct {
		rowDelta, colDelta int
	}{
		{0, 1},  // Right
		{0, -1}, // Left
		{-1, 0}, // Up
		{1, 0},  // Down
	}

	for _, dir := range directions {
		newRow, newCol := currentPos.row+dir.rowDelta, currentPos.col+dir.colDelta
		if newRow >= 0 && newRow < len(m) && newCol >= 0 && newCol < len(m[0]) &&
			m[newRow][newCol] == currentPos.heigth+1 {
			nextPos := position{newRow, newCol, currentPos.heigth + 1}
			path = append(path, nextPos)
			backtrack(nextPos, m, path, result)
			path = path[:len(path)-1]
		}
	}
}
