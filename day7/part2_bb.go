package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ecuation struct {
	result int
	list   []int
}

type pair struct {
	first  int
	second int
}

func concat(x, y int) int {
	str := strconv.Itoa(x) + strconv.Itoa(y)
	res, _ := strconv.Atoi(str)
	return res
}

// now using branch and bound, it is considerably faster (450 ms) than brute force (1.7 s) or backtrackint (1.45 s)
// But running the result is a little diferent to the correct answer. I have a bug somewhere
func verifyEcuation(e ecuation) bool {
	stack := []pair{pair{0, 0}}
	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if state.first == len(e.list) && state.second == e.result {
			return true
		}

		if state.first < len(e.list) && state.second <= e.result {
			stack = append(stack, pair{state.first + 1, state.second + e.list[state.first]})
			stack = append(stack, pair{state.first + 1, state.second * e.list[state.first]})
			stack = append(stack, pair{state.first + 1, concat(state.second, e.list[state.first])})
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
	lines = lines[:len(lines)-1]

	var ecuations []ecuation

	for _, line := range lines {
		aux := strings.Split(line, ":")
		result, _ := strconv.Atoi(aux[0])
		listStr := strings.Fields(aux[1])
		list := make([]int, len(listStr))
		for i, elem := range listStr {
			list[i], _ = strconv.Atoi(elem)
		}
		ecuations = append(ecuations, ecuation{result, list})
	}

	sum := 0
	for _, ec := range ecuations {
		if verifyEcuation(ec) {
			sum += ec.result
		}
	}

	fmt.Println("Result: ", sum)

}
