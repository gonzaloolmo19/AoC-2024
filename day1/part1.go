package main

import (
	"fmt"
	"os"
	"sort"
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
	array2 := make([]int, 0, 5)

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
		array2 = append(array2, field2)
	}

	sort.Ints(array1)
	sort.Ints(array2)

	// fmt.Println(array1)
	// fmt.Println(array2)

	sum := 0
	for i := 0; i < len(array1); i++ {
		sum += absInt(array1[i] - array2[i])
	}

	fmt.Println("Result: ", sum)
}
