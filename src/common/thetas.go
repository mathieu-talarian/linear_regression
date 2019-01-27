package common

type Thetas struct {
	Zero float64 `json:"theta0" csv:"theta0"`
	One  float64 `json:"theta1" csv:"theta1"`
}

func (thetas *Thetas) train(learningRate float64, t Thetas, normalized exploiatableDatas) {
	var sum0 float64
	var sum1 float64
	m := normalized.Len()
	sum0, sum1 = normalized.Sums(t)
	thetas.Zero = learningRate * (float64(1) / float64(m)) * sum0
	thetas.One = learningRate * (float64(1) / float64(m)) * sum1
}

func (thetas *Thetas) SetZero() {
	thetas.Zero = 0
	thetas.One = 0
}

func TrainLoop(learningRate float64, normalized exploiatableDatas, loops int) (thetas Thetas) {
	for i := 0; i < loops; i++ {
		var tmp Thetas
		tmp.train(learningRate, thetas, normalized)
		thetas.Zero -= tmp.Zero
		thetas.One -= tmp.One
	}
	return
}
