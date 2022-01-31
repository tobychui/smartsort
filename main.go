package main

import (
	"fmt"

	"github.com/tobychui/smartsort/sorter"
)

func main() {
	/*
		filenames, _ := filepath.Glob("./test/*.jpg")
		log.Println(filenames)
		sortedFilenames := sorter.SortString(filenames)
		log.Println(sortedFilenames)
	*/

	original := []string{"a1c.o", "A1c.o", "a2016c.o", "a10c.o", "a2c.o", "A2c.o", "a11c.o"}
	fmt.Println("Before: ", original)
	results := sorter.SortString(original)
	fmt.Println("After: ", results)

}
