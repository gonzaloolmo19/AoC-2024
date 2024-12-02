package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func verifyReport(line string) bool {
	if (len(line)) == 1 {
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
	fmt.Println(numbers) 

	isAsc := numbers[0] < numbers[1]
	fmt.Println(isAsc) 
	for i:=0; i< len(numbers)-1; i++ {
		if isAsc {
			if numbers[i] > numbers[i+1] || (1 > numbers[i+1] - numbers[i] || numbers[i+1] - numbers[i] > 3) {
				fmt.Println("i = ", i, "saliendo") 
				return false
			}
		} else {
			if numbers[i] < numbers[i+1] || (1 > numbers[i]-numbers[i+1] || numbers[i] - numbers[i+1] > 3) {
				fmt.Println("dif: ", numbers[i]-numbers[i+1], "i = ", i, "saliendo") 
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
	fmt.Println(len(lines))

	cont:= 0

	for _, line := range lines {
		if line != "" {
			fmt.Println("verify:", verifyReport(line) ) 
			if verifyReport(line) {
				cont += 1
			}
		}
	}

	fmt.Println("Result: ", cont) 



}
