package main

import (
	"fmt"
	"os"
	"strconv"
)


func makeDisk(s string) []int {
	isFile := true
	var disk []int
	id := 0


	for i := range s {
		if isFile {
			fileLength, _ := strconv.Atoi(string(s[i]))
			for j := 0; j < fileLength; j++ {
				disk = append(disk, id) 
			}
			id++
			isFile = false
		} else {
			emptyLenght, _ := strconv.Atoi(string(s[i]))
			for j := 0; j < emptyLenght; j++ {
				disk = append(disk, -1)
			}
			isFile = true
		}
	}
	return disk
}

func arrangeDisk(disk []int) {
	moveP := len(disk) - 1
	emptyP := 0
	for emptyP < len(disk) && disk[emptyP] != -1 {
		emptyP++
	}

	for emptyP < len(disk) && moveP >= 0 && emptyP < moveP {
		// swap characters
		disk[emptyP], disk[moveP] = disk[moveP], disk[emptyP]
		for emptyP < len(disk) && disk[emptyP] != -1 {
			emptyP++
		}
		moveP--
	}
}

func checkSum(disk []int) int {
	sum := 0
	for i := 0; i < len(disk) && disk[i] != -1; i++ {
		sum += i * disk[i]
	}
	return sum
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
	input := string(data)

	disk := makeDisk(input)

	//fmt.Println(string(disk))

	arrangeDisk(disk)

	//fmt.Println(string(disk))

	fmt.Println("Result: ", checkSum(disk))
}
