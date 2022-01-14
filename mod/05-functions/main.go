package main

import (
	"fmt"
	common "nst-golang-course/00-common"
)

func main() {
	common.LabelMessage("Functions!")
	data := []float64{5, 6, 4, 7, 3, 8, 2, 9, 1}
	fmt.Println("Data: ", data)
	m := mean(data...)
	fmt.Println("Mean: ", m)
}

func mean(items ...float64) float64 {
	var result float64 = 0
	var iterations float64 = 0
	// with "_" index gets ignored
	for _, item := range items {
		result += item
		iterations += 1
	}

	return result / iterations

	// min09
}
