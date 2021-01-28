package main

import (
	"fmt"
	"sort"
)

func main() {
	var grades = [...]float32{89, 86, 94, 100, 91, 72, 89, 85, 97, 95, 100, 99}
	gradeSlice := grades[:]
	average := getAverage(&gradeSlice)
	median := getMedian(&gradeSlice)
	fmt.Printf("Average: %.2f%%\nMedian: %.2f%%\n", average, median)
}

func getAverage(nums *[]float32) float32 {
	var sum float32 = 0
	for _, n := range *nums {
		sum += n
	}
	return sum / float32(len(*nums))
}

func getMedian(nums *[]float32) float32 {
	var sortedNums []float64
	for _, n := range *nums {
		sortedNums = append(sortedNums, float64(n))
	}
	sort.Float64s(sortedNums)
	var midIndex int = len(sortedNums) / 2
	if len(sortedNums)%2 == 0 {
		return float32((sortedNums[midIndex] + sortedNums[midIndex-1]) / 2)
	}
	return float32(sortedNums[midIndex])
}
