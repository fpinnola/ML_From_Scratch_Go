package examples

import (
	"fmt"
	"math/rand"
)

func predict(theta []float32, x []float32) (res float32) {
	var sum float32 = 0
	for i, xval := range x {
		sum = sum + theta[i]*xval
	}
	return sum
}

func training_step(theta []float32, X [][]float32, y []float32, alpha float32) (res []float32) {
	new_theta := make([]float32, len(theta))
	copy(new_theta, theta)

	for i, _ := range new_theta {
		var cost float32 = 0
		for j, xvals := range X {
			prediction := predict(theta, xvals)
			cost = cost + (prediction-y[j])*xvals[i]
		}
		new_theta[i] = theta[i] - alpha*(float32(1)/float32(len(X)))*cost
	}

	return new_theta
}

func LinearRegression() {
	fmt.Println("Here is a Linear regression example")
	rand.Seed(1)

	X_train := make([][]float32, 10)
	for i := 0; i < 10; i++ {
		val_0 := float32(1)
		val_1 := rand.Float32() * 10
		val_2 := rand.Float32() * 10
		xval := []float32{val_0, val_1, val_2}
		X_train[i] = xval
	}
	y_train := make([]float32, len(X_train))
	theta := make([]float32, len(X_train[0]))

	for i, _ := range theta {

		theta[i] = rand.Float32()
	}
	for i, v := range X_train {
		y_train[i] = v[0]*2 + v[1]*1
	}

	for i := 0; i < 2000; i++ {
		theta = training_step(theta, X_train, y_train, 0.01)
	}

	test_predict := predict(theta, []float32{1, 2, 3})
	fmt.Println("Test input {1,2,3} prediction: ", test_predict)
	fmt.Println("Expected: ", 1*2+2*1)
}
