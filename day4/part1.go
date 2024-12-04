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

	fmt.Println("w ", width)
	fmt.Println("h ", height)

	counter := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lines[i][j] == 'X' {
				//fmt.Println(i, j)
				if j+3 < width && lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
					counter++
				}
				if i+3 < height && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					counter++
				}
				if i+3 < height && j+3 < width && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					counter++
				}
				if i+3 < height && j-3 >= 0 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					counter++
				}
				if j-3 >= 0 && lines[i][j-1] == 'M' && lines[i][j-2] == 'A' && lines[i][j-3] == 'S' {
					counter++
				}
				if i-3 >= 0 && j-3 >= 0 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					counter++
				}
				if i-3 >= 0 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					counter++
				}
				if i-3 >= 0 && j+3 < width && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					counter++
				}
			}
		}
	}
	fmt.Println("Result: ", counter)
}
