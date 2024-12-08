package main

import (
	"fmt"
	"os"
	"strings"
)

type vector struct {
	row int
	col int
}

func (p1 vector) Add(p2 vector) vector {
	return vector{p1.row + p2.row, p1.col + p2.col}
}

func (p1 vector) Sub(p2 vector) vector {
	return vector{p1.row - p2.row, p1.col - p2.col}
}

func ValidPos(pos vector, width, height int) bool {
	return pos.row >= 0 && pos.row < height && pos.col >= 0 && pos.col < width
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
	height := len(lines)
	width := len(lines[0])

	m := make(map[rune][]vector)

	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				m[char] = append(m[char], vector{i, j})
			}
		}
	}

	a := make(map[vector]bool)
	for key := range m {
		antenas := m[key]
		for i := 0; i < len(antenas)-1; i++ {
			for j := i + 1; j < len(antenas); j++ {
				vector := antenas[i].Sub(antenas[j])
				antinode := antenas[i]
				for ValidPos(antinode, width, height) {
					a[antinode] = true
					antinode = antinode.Add(vector)
				}
				vector = antenas[j].Sub(antenas[i])
				antinode = antenas[j]
				for ValidPos(antinode, width, height) {
					a[antinode] = true
					antinode = antinode.Add(vector)
				}
			}
		}
	}

	fmt.Println("Antinodes: ")
	for key := range a {
		fmt.Println(key)
	}
	fmt.Println("Result: ", len(a))

}
