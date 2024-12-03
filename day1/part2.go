package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
	array1 := make([]int, 0, 5)
	map2 := make(map[int]int)

	for _, line := range lines {
		if line == "" {
			break
		}
		columns := strings.Fields(line)
		field1, err := strconv.Atoi(columns[0])
		if err != nil {
			fmt.Println("Could not convert to int")
			os.Exit(1)
		}
		array1 = append(array1, field1)
		field2, err := strconv.Atoi(columns[1])
		if err != nil {
			fmt.Println("Could not convert to int")
			os.Exit(1)
		}
		map2[field2]++
	}

	// fmt.Println(array1)
	// fmt.Println(map2)

	sum := 0
	for i := 0; i < len(array1); i++ {
		sum += array1[i] * map2[array1[i]]
	}

	fmt.Println("Result: ", sum)
}
