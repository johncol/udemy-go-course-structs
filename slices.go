package main

import "fmt"

func duplicateValues(numbers []int) {
	for i, number := range numbers {
		numbers[i] = number * 2
	}
}

func slicesSample() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	duplicateValues(numbers)
	fmt.Println(numbers)
}
