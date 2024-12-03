package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type pair struct {
	first int
	second int
}

func main() {

	re := regexp.MustCompile("mul\\([[:digit:]]{1,3},[[:digit:]]{1,3}\\)")
	reNum := regexp.MustCompile("[[:digit:]]{1,3}") 

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
	fmt.Println(re.FindAllString(text, -1))
	mults := re.FindAllString(text, -1)

	operands := make([]pair, len(mults), len(mults))
	for i:= 0; i < len(mults); i++ {
		nums:= reNum.FindAllString(mults[i], -1)
		if len(nums) != 2 {
			fmt.Println("Wrong mult syntax")
			os.Exit(1) 
		}
		operands[i].first, err = strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Could not convert string to int")
			os.Exit(1) 
		}
		operands[i].second, err = strconv.Atoi(nums[1])
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

	fmt.Println("Result: ", sum) 


}
