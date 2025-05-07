package main

import "fmt"

func main() {
	fmt.Println(FilterEven(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func FilterEven(numbers ...int) []int {
	filtered := make([]int, 0)

	for _, num := range numbers {
		if num%2 == 0 {
			filtered = append(filtered, num)
		}
	}

	return filtered
}
