package main

import (
	"container/list"
	"fmt"
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

	rocks := list.New()

	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		rocks.PushBack(num)
	}
	nBlinks := 75

	for i := 0; i < nBlinks; i++ {
		fmt.Println("iteracion ", i)
		Blink(rocks)
	}

	// for r := rocks.Front(); r != nil; r = r.Next() {
	// 	fmt.Print(r.Value, " ")
	// }

	fmt.Println("Result: ", rocks.Len())

}

func Blink(r *list.List) {
	for e := r.Front(); e != nil; e = e.Next() {
		if e.Value == 0 {
			e.Value = 1
		} else if len(strconv.Itoa(e.Value.(int)))%2 == 0 {
			str := strconv.Itoa(e.Value.(int))
			firstStr := str[:len(str)/2]
			secondStr := str[len(str)/2:]
			first, _ := strconv.Atoi(firstStr)
			second, _ := strconv.Atoi(secondStr)
			r.InsertBefore(first, e)
			aux := r.InsertBefore(second, e)
			r.Remove(e)
			e = aux
		} else {
			e.Value = e.Value.(int) * 2024
		}
		fmt.Println(e.Value)
	}
}
