package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type pair struct {
	first  int
	second int
}

var re *regexp.Regexp = regexp.MustCompile("mul\\([[:digit:]]{1,3},[[:digit:]]{1,3}\\)")
var reNum *regexp.Regexp = regexp.MustCompile("[[:digit:]]{1,3}")

func getMults(text string) int {
	var err error
	mults := re.FindAllString(text, -1)
	fmt.Println("Mults: ", mults)
	operands := make([]pair, len(mults), len(mults))
	for i := 0; i < len(mults); i++ {
		nums := reNum.FindAllString(mults[i], -1)
		if len(nums) != 2 {
			fmt.Println("Wrong mult syntax")
			os.Exit(1)
		}
		operands[i].first, err = strconv.Atoi(nums[0])
		fmt.Println("first: ", operands[i].first)
		if err != nil {
			fmt.Println("Could not convert string to int")
			os.Exit(1)
		}
		operands[i].second, err = strconv.Atoi(nums[1])
		fmt.Println("second: ", operands[i].second)
		if err != nil {
			fmt.Println("Could not convert string to int")
			os.Exit(1)
		}
	}
	fmt.Println(operands)

	sum := 0
	for _, operand := range operands {
		sum += operand.first * operand.second
	}
	return sum
}

func main() {

	reDo := regexp.MustCompile("do\\(\\)")
	reDont := regexp.MustCompile("don't\\(\\)")

	enabled := true

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
	fmt.Println(reDont.FindAllString(text, -1))
	fmt.Println(reDo.FindAllString(text, -1))

	sum := 0
	cont := true
	index := 0
	remaining := text

	for cont {
		if enabled {
			match := reDont.FindStringIndex(remaining)
			if match == nil {
				cont = false
				fmt.Println("processing", remaining)
				sum += getMults(remaining)
			} else {
				enabled = false
				index = match[1]
				fmt.Println("processing", remaining[:index])
				sum += getMults(remaining[:index])
				remaining = remaining[index:]
				fmt.Println("Remaining: ", remaining)
			}
		} else {
			match := reDo.FindStringIndex(remaining)
			if match == nil {
				cont = false
			} else {
				enabled = true
				index = match[1]
				remaining = remaining[index:]
				fmt.Println("Remaining: ", remaining)

			}
		}
	}

	fmt.Println("Result: ", sum)

}
