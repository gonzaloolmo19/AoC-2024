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

func concat(x, y int) int {
	str := strconv.Itoa(x) + strconv.Itoa(y)
	res, _ := strconv.Atoi(str)
	return res
}

// Now I use backtracking. It is a little bit faster than brute force for
// the input, but not too much (1.45 s)
func backtrack(e ecuation, path []int, nSol *int) {
	if *nSol > 0 {
		return
	}
	nOperators := len(e.list) - 1
	calcRes := e.list[0]
	for i, op := range path {
		if op == 0 {
			calcRes+= e.list[i+1]
		} else if op == 1 {
			calcRes *= e.list[i+1]
		} else if op == 2 {
			calcRes = concat(calcRes, e.list[i+1])
		}
	}
	if len(path) == nOperators && calcRes == e.result {
		*nSol++
	}

	if calcRes <= e.result && len(path) < nOperators{
		for i := 0; i < 3; i++ {
			path = append(path, i)
			backtrack(e, path, nSol)
			path = path[:len(path)-1]
		}
	}
}

func verifyEcuation(e ecuation) bool {
	nSol := 0
	backtrack(e, []int{}, &nSol)
	return nSol > 0
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
			sum+= ec.result
		}
	}

	fmt.Println("Result: ", sum)

}
