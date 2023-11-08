package math

//Use go doc introducingGo/src/chapter8/math Average to see the documentation

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}
