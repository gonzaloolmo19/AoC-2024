package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contained(num int, list []int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
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

	var rules, updatesStr []string

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			fmt.Println(i)
			rules = lines[:i]
			updatesStr = lines[i+1:]
			break
		}
	}

	updatesStr = updatesStr[:len(updatesStr)-1]
	fmt.Println(rules)
	fmt.Println(updatesStr)

	ruleMap := make(map[int][]int)

	for _, rule := range rules {
		numbers := strings.Split(rule, "|")
		if len(numbers) != 2 {
			fmt.Println("Bad split in rules")
			os.Exit(1)
		}
		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Could not parse int from rules")
			os.Exit(1)
		}
		second, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Could not parse int from rules")
			os.Exit(1)
		}
		ruleMap[first] = append(ruleMap[first], second)
	}

	fmt.Println(ruleMap)

	updates := make([][]int, 0)
	for _, line := range updatesStr {
		splited := strings.Split(line, ",")
		fmt.Println(splited)
		update := make([]int, 0)
		for _, numStr := range splited {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Could not parse int from updates")
				os.Exit(1)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	fmt.Println(updates)

	sum := 0
	for _, update := range updates {
		validUpdate := true
		for i := 0; i < len(update) && validUpdate; i++ {
			for j := i + 1; j < len(update) && validUpdate; j++ {
				// fmt.Println(update[i])
				// fmt.Println(ruleMap[update[j]])
				if contained(update[i], ruleMap[update[j]]) {
					validUpdate = false
				}
			}
		}
		if validUpdate {
			sum += update[len(update)/2]
			fmt.Println("sumando ", update[len(update)/2])
		}
	}

	fmt.Println("Result: ", sum)

}
