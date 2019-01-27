package common

import "math"

type Limits struct {
	Min float64 `json:"min" csv:"min"`
	Max float64 `json:"max" csv:"max"`
}

func (l *Limits) Get(datas exploiatableDatas) {
	l.Min = datas.GetMin()
	l.Max = datas.GetMax()
}

func (l *Limits) SetZero() {
	l.Min = 0
	l.Max = math.Inf(1)
}
