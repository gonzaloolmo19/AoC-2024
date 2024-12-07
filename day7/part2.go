package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ecuation struct {
	result int
	list   []int
}

type generator struct {
	n       int
	current int
	max     int
}

func newGenerator(n int) *generator {
	return &generator{
		n:       n,
		current: 0,
		max:     int(math.Pow(3, float64(n))),
	}
}

func (gen *generator) Next() []int {
	if gen.current >= gen.max {
		return nil
	}

	combination := make([]int, gen.n)
	acc := gen.current
	for i := 0; i < gen.n; i++ {
		combination[gen.n-i-1] = acc % 3
		acc /= 3
	}
	gen.current++

	return combination
}

func concat(x, y int) int {
	str := strconv.Itoa(x) + strconv.Itoa(y)
	res, _ := strconv.Atoi(str)
	return res
}

// Could have done some backtracking for more efficiency, but
// brute force was fast enough for the input (< 2 seconds)
func validEcuation(e ecuation) bool {
	nOperators := len(e.list) - 1
	gen := newGenerator(nOperators)

	for op := gen.Next(); op != nil; op = gen.Next() {
		partialRes := e.list[0]
		for j := 0; j < len(op); j++ {
			if op[j] == 0 {
				partialRes += e.list[j+1]
			} else if op[j] == 1 {
				partialRes *= e.list[j+1]
			} else if op[j] == 2 {
				partialRes = concat(partialRes, e.list[j+1])
			}
		}
		if partialRes == e.result {
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
		if validEcuation(ec) {
			sum += ec.result
		}
	}

	fmt.Println("Result: ", sum)

}
