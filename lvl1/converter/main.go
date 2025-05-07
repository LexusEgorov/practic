package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("not enought args")
	}

	cTemp, err := strconv.ParseFloat(args[1], 64)

	if err != nil {
		panic(err)
	}

	fTemp := cTemp*1.8 + 32

	fmt.Printf("%.2f°C = %.2f°F\n", cTemp, fTemp)
}
