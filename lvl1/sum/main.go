package main

import "fmt"

func main() {
	fmt.Println(SumDigits(123))
	fmt.Println(SumDigits(9875))
}

func SumDigits(n int) int {
	sum := 0

	if n < 0 {
		n *= -1
	}

	for n > 0 {
		sum += n % 10

		n = n / 10
	}

	return sum
}
