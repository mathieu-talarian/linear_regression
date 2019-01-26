package common

type Limits struct {
	Min float64 `json:"min" csv:"min"`
	Max float64 `json:"max" csv:"max"`
}

func (l *Limits) Get(datas exploiatableDatas) {
	l.Min = datas.GetMin()
	l.Max = datas.GetMax()
}
