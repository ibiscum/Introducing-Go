package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(averageWithNamedOutput(xs))

	// Multiple outputs
	average, total := averageWithMultiplesOutputs(xs)
	fmt.Println(average, total)
}

// We use camelCase for function names when they are private (unexported)
// We use PascalCase for function names when they are public (exported)
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func averageWithNamedOutput(xs []float64) (average float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	average = total / float64(len(xs))
	return
}

func averageWithMultiplesOutputs(xs []float64) (average float64, total float64) {
	total = 0.0
	for _, v := range xs {
		total += v
	}
	average = total / float64(len(xs))

	//Is not necessary to specify the name of the variables, because they are already defined in the function declaration
	return average, total
}
