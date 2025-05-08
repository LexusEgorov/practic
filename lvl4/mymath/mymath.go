package mymath

func Average(numbers []float64) float64 {
	var sum float64

	for _, n := range numbers {
		sum += n
	}

	return sum / float64(len(numbers))
}
