package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeNumAt(slice []int, pos int) []int {
	if pos < 0 || pos >= len(slice) {
		return slice // Return the slice unchanged if position is invalid
	}

	// Create a copy of the slice
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)

	// Remove the element from the copy
	copySlice = append(copySlice[:pos], copySlice[pos+1:]...)

	return copySlice
}

// Not the most efficient way to solve
func verifyReport(line string) bool {
	if len(line) == 1 {
		return true
	}
	fields := strings.Fields(line)
	numbers := make([]int, 0, len(fields))
	for _, char := range fields {
		number, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println("Could not parse number")
			os.Exit(1) 
		}
		numbers = append(numbers,number)
	}

	if verifyReportPartial(numbers) {
		return true
	}

	for i:=0; i<len(line); i++ {
		if verifyReportPartial(removeNumAt(numbers, i)) {
			return true
		}
	}
	return false
}

func verifyReportPartial(numbers []int) bool {
	if (len(numbers)) == 1 {
		return true
	}
	
	isAsc := numbers[0] < numbers[1]
	for i:=0; i< len(numbers)-1; i++ {
		if isAsc {
			if numbers[i] > numbers[i+1] || (1 > numbers[i+1] - numbers[i] || numbers[i+1] - numbers[i] > 3) {
				return false
			}
		} else {
			if numbers[i] < numbers[i+1] || (1 > numbers[i]-numbers[i+1] || numbers[i] - numbers[i+1] > 3) {
				return false
			}
		}
	}

	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	// Read the entire file into memory
	data, err := os.ReadFile(os.Args[1] )
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(data)
	lines := strings.Split(text, "\n")


	cont:= 0

	for _, line := range lines {
		if line != "" {
			if verifyReport(line) {
				cont += 1
			}
		}
	}

	fmt.Println("Result: ", cont) 

// Not the most efficient way to solve
}
