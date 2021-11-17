package utils

type Matrix32 struct {
	Values [][]float32
}

func create_2dslice(cols int, rows int) (res [][]float32) {
	solution := make([][]float32, rows)
	for x := range solution {
		solution[x] = make([]float32, cols)
	}
	return solution
}

func GetValues() []float32 {
	m := Matrix32{make([][]float32, 2)}
	m.Values[0] = []float32{1, 2, 3, 4}
	m.Values[1] = []float32{5, 6, 7, 8}
	out := []float32{}
	for _, v := range m.Values {
		for _, vx := range v {
			out = append(out, vx)
		}
	}
	return out
}

func Scale(m1 Matrix32, s float32) (prod Matrix32) {
	/*
		Multiply each element in the matrix by a scalar value
	*/
	solution := make([][]float32, len(m1.Values))
	for x := range solution {
		solution[x] = make([]float32, len(m1.Values[x]))
	}

	for i := 0; i < len(m1.Values); i++ {
		for j := 0; j < len(m1.Values[0]); j++ {
			solution[i][j] = m1.Values[i][j] * s
		}
	}
	m_prod := Matrix32{solution}
	return m_prod
}

func Dot(m1 Matrix32, m2 Matrix32) (prod Matrix32) {
	/*
		Matrix Multiplication is possible when the number of columns
		of the first matrix is equal to the number of rows
		of the second matrix
		len(m1[0]) == len(m2)
	*/

	if len(m1.Values[0]) != len(m2.Values) {
		panic("Matrices must be of dimensions (m,n) x (n,o)")
	}
	cols := len(m2.Values[0])
	rows := len(m1.Values)
	solution := make([][]float32, rows)
	for x := range solution {
		solution[x] = make([]float32, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < len(m1.Values[0]); k++ {
				solution[i][j] = solution[i][j] + (m1.Values[i][k] * m2.Values[k][j])
			}
		}
	}
	m_product := Matrix32{solution}
	return m_product

}

func Transpose(m1 Matrix32) (inv Matrix32) {
	/*
		Returns the Transpose of a given matrix
	*/
	var solution [][]float32
	if len(m1.Values) == len(m1.Values[0]) {
		//Square Matrix
		solution = create_2dslice(len(m1.Values), len(m1.Values[0]))

		for i, v := range m1.Values {
			for j := range v {
				solution[i][j] = m1.Values[j][i]
			}
		}
	} else {
		//Rectangular Matrix
		solution = create_2dslice(len(m1.Values), len(m1.Values[0]))
		for i := 0; i < len(solution); i++ {
			for j := 0; j < len(solution[i]); j++ {
				solution[i][j] = m1.Values[j][i]
			}
		}
	}
	return Matrix32{Values: solution}
}

func Inverse(m1 Matrix32) (prod Matrix32) {
	/*
		Gauss - Jordan Elementary Row Operation
		Returns the Inverse of a given Matrix
	*/
	if len(m1.Values) != len(m1.Values[0]) {
		panic("Matrix must be square")
	}

	order := len(m1.Values)
	// Create matrix of 2*n by n size and initializing to 0
	identity := create_2dslice(order*2, len(m1.Values))
	// Add in matrix m1 Values
	for i := 0; i < len(m1.Values); i++ {
		for j := 0; j < len(m1.Values[i]); j++ {
			identity[i][j] = m1.Values[i][j]
		}
	}
	// Add identity matrix to end of matrix m1
	for i := 0; i < order; i++ {
		for j := 0; j < 2*order; j++ {
			if j == (i + order) {
				// fmt.Println(i, j)
				identity[i][j] = 1
			}
		}
	}
	for i := 0; i < order; i++ {
		if identity[i][i] == 0 {
			panic("Matrix is not invertible")
		}

		for j := 0; j < order; j++ {
			if i != j {
				ratio := identity[j][i] / identity[i][i]

				for k := 0; k < order*2; k++ {
					identity[j][k] = identity[j][k] - ratio*identity[i][k]
				}
			}
		}
	}

	// Row operation
	for i := 0; i < order; i++ {
		divisor := identity[i][i]
		for j := 0; j < order*2; j++ {
			identity[i][j] = identity[i][j] / divisor
		}
	}
	solution := make([][]float32, order)
	for i := range solution {
		solution[i] = identity[i][order:]
	}

	return Matrix32{Values: solution}
}

func Test() {

}
