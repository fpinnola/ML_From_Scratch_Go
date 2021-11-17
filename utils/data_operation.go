package utils

func MeanSquaredError(y_true []float32, y_pred []float32) (error float32) {
	var mse float32 = 0
	for i := range y_true {
		mse = mse + (y_true[i]-y_pred[i])*(y_true[i]-y_pred[i])
	}
	mse = mse / float32(len(y_true))
	return mse
}
