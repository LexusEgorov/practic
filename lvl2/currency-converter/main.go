package main

import "fmt"

type CurrencyConverter struct {
	Rate float64
}

func (c CurrencyConverter) ConvertToUSD(rubles float64) float64 {
	return rubles / c.Rate
}

func (c CurrencyConverter) ConvertToRUB(dollars float64) float64 {
	return dollars * c.Rate
}

func main() {
	converter := CurrencyConverter{Rate: 75.5}
	fmt.Println(converter.ConvertToUSD(755))
}
