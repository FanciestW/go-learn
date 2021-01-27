package main

import "fmt"

func main() {
	var grades = [...]float32{89, 86, 94, 100, 91, 72, 89, 85, 97, 95, 100, 99}
	gradeSlice := grades[:]
	average := getAverage(&gradeSlice)
	fmt.Println(average)
}

func getAverage(nums *[]float32) float32 {
	var sum float32 = 0
	for _, n := range *nums {
		sum += n
	}
	return sum / float32(len(*nums))
}
