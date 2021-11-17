package main

import (
	"fmt"

	"fpinnola.dev/ml-from-scratch/examples"
)

func main() {
	fmt.Println()
	fmt.Println("Here are some Matrix operation examples")
	examples.RunMatrixExamples()
	fmt.Println()
	examples.LinearRegression()
}
