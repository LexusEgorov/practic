package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(n int) string {
	res := ""

	if n%3 == 0 {
		res = "Fizz"
	}

	if n%5 == 0 {
		res += "Buzz"
	}

	if len(res) == 0 {
		return fmt.Sprint(n)
	}

	return res
}
