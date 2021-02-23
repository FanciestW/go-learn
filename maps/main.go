package main

import (
	"fmt"
)

func main() {
	// Creates a map where key is a string and value is an int.
	myMap := make(map[int]int)

	var numbers []int = []int{1, 2, 3, 1, 2, 3, 5, 7, 8, 2, 5, 9}

	for i, n := range numbers {
		_, ok := myMap[n]
		if !ok {
			myMap[n] = 1
		} else {
			myMap[n]++
		}
		fmt.Printf("i: %d\tn: %d\n", i, n)
	}
	fmt.Println(myMap)
}
