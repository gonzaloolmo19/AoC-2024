package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type File struct {
	id     int
	length int
}

func makeDisk(s string) *list.List {
	isFile := true
	disk := list.New()
	id := 0

	for i := range s {
		if isFile {
			fileLength, _ := strconv.Atoi(string(s[i]))
			if fileLength > 0 {
				disk.PushBack(File{id, fileLength})
			}
			id++
			isFile = false
		} else {
			emptyLenght, _ := strconv.Atoi(string(s[i]))
			if emptyLenght > 0 {
				disk.PushBack(File{-1, emptyLenght})
			}
			isFile = true
		}
	}
	return disk
}

// It is a very ugly function but I don't have time to think it better :(
// It is very important the order of insertions and removal in the list,
// if not, there could be some edge cases
func arrangeDisk(disk *list.List) {
	maxId := 0

	for e := disk.Back(); e != nil; e = e.Prev() {
		maxId = e.Value.(File).id
		if maxId != -1 {
			break
		}
	}

	for true {
		file := disk.Back()
		for e := disk.Back(); e != nil; e = e.Prev() {
			id := e.Value.(File).id
			if maxId == id {
				file = e
				break
			}
		}
		fileValue, _ := file.Value.(File)
		fileLength := fileValue.length
		//fmt.Println("Moving", fileValue.id)

		if fileValue.id == 0 {
			break
		}
		for e := disk.Front(); e != file; e = e.Next() {

			value, _ := e.Value.(File)
			left := value.length - fileLength

			next := e.Next()
			// if it is empty space and the file fits in there
			if value.id == -1 && left >= 0 {
				disk.InsertBefore(fileValue, e)
				disk.Remove(e)
				newEmptyLength := fileValue.length
				if file.Prev() != nil && file.Prev().Value.(File).id == -1 {
					newEmptyLength += file.Prev().Value.(File).length
					disk.Remove(file.Prev())
				}
				if file.Next() != nil && file.Next().Value.(File).id == -1 {
					newEmptyLength += file.Next().Value.(File).length
					disk.Remove(file.Next())
				}
				disk.InsertBefore(File{-1, newEmptyLength}, file)
				disk.Remove(file)

				if left > 0 {
					disk.InsertBefore(File{-1, left}, next)
				}
				e = next.Prev()
				break
			}
		}
		maxId--
	}

}

func checkSum(disk *list.List) int {
	sum := 0
	pos := 0
	for e := disk.Front(); e != nil; e = e.Next() {
		value := e.Value.(File)
		for i := 0; i < value.length; i++ {
			if value.id != -1 {
				sum += pos * value.id
			}
			pos++
		}
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

	arrangeDisk(disk)
	// for e:= disk.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	fmt.Println("Result: ", checkSum(disk))
}
