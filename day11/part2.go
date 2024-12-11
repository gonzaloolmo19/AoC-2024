package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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
	fields := strings.Fields(text)

	rocks := make(map[int]int)

	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		rocks[num]++
	}

	nBlinks := 75
	fmt.Println("Result: ", Blink(rocks, nBlinks))

}

func Blink(rocks map[int]int, iterations int) int {
	for i := 0; i < iterations; i++ {
		newRocks := make(map[int]int)
		for stone := range rocks {
			if stone == 0 {
				newRocks[1] += rocks[stone]
			} else {
				nDigits := int(math.Log10(float64(stone))) + 1
				if nDigits%2 == 0 {
					div := int(math.Pow10(nDigits / 2))
					newRocks[stone/div] += rocks[stone]
					newRocks[stone%div] += rocks[stone]
				} else {
					newRocks[stone*2024] += rocks[stone]
				}

			}
		}
		rocks = newRocks
	}

	sum := 0
	for _, value := range rocks {
		sum += value
	}
	return sum
}
