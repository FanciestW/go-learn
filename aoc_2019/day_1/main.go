package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() []int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var data []int

	for scanner.Scan() {
		var text string = scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			log.Println("Unable to convert text to int")
			continue
		}
		data = append(data, num)
	}

	file.Close()
	return data
}

func calcFuelReq(mass int) int {
	return (mass / 3) - 2
}

func main() {
	var data []int = readFile()
	var fuelReqs []int
	for _, num := range data {
		fuelReqs = append(fuelReqs, calcFuelReq(num))
	}

	var sum int = 0
	for _, num := range fuelReqs {
		sum += num
	}
	fmt.Println(sum)
}
