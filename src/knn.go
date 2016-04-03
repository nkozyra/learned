package learned

type KNN struct {
	xData []float64
	yData []float64
}

func NewKNN() (k KNN) {
	return k
}

func (k KNN) AddTrain(xTrain []float64, yTrain []float64) {
	k.xData = xTrain
	k.yData = yTrain
}

func (k KNN) Query() {

}
