package main

import (
	"fmt"
	"os"
	"strings"
)

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

	counter := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lines[i][j] == 'A' {
				if j+1 < width && j-1 >= 0 && i+1 < height && i-1 >= 0 {
					if lines[i+1][j+1] == 'S' && lines[i-1][j+1] == 'M' && lines[i-1][j-1] == 'M' && lines[i+1][j-1] == 'S' {
						counter++
					}
					if lines[i+1][j+1] == 'S' && lines[i-1][j+1] == 'S' && lines[i-1][j-1] == 'M' && lines[i+1][j-1] == 'M' {
						counter++
					}
					if lines[i+1][j+1] == 'M' && lines[i-1][j+1] == 'S' && lines[i-1][j-1] == 'S' && lines[i+1][j-1] == 'M' {
						counter++
					}
					if lines[i+1][j+1] == 'M' && lines[i-1][j+1] == 'M' && lines[i-1][j-1] == 'S' && lines[i+1][j-1] == 'S' {
						counter++
					}
				}
			}
		}
	}
	fmt.Println("Result: ", counter)
}
