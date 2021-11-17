package examples

import (
	"fmt"

	"fpinnola.dev/ml-from-scratch/utils"
)

func RunMatrixExamples() {
	// Examples of the Matrix Operations
	m1 := utils.Matrix32{Values: [][]float32{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}}
	m3 := utils.Matrix32{Values: [][]float32{{1, 2, 3}, {0, 1, 4}, {5, 6, 0}}}
	m2 := utils.Matrix32{Values: [][]float32{{2, 3}, {5, 6}, {8, 9}}}
	fmt.Println(utils.Dot(m1, m2))
	fmt.Println(utils.Scale(m1, 3))
	fmt.Println(utils.Transpose(m1))
	fmt.Println(utils.Transpose(m2))
	fmt.Println(utils.Inverse(m3))
}
