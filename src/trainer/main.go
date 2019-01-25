package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func checkErr(dt interface{}, err error) interface{} {
	if err != nil {
		log.Fatal("Aïe ", err)
	}
	return dt
}

func normalizeDatas(dataTab []datas, limits Limits) (normalized []datas) {
	for _, v := range dataTab {
		normalized = append(normalized, v.normalize(limits))
	}
	return
}

type Thetas struct {
	Zero float64 `json:"theta0", csv:"theta0"`
	One  float64 `json:"theta1", csv:"theta1"`
}

func (tmp *Thetas) train(learningRate float64, t *Thetas, normalized []datas) {
	var sum0 float64
	var sum1 float64
	m := len(normalized)
	for _, v := range normalized {
		d := v.estimatePrice(t)
		sum0 += d
		sum1 += (d * v.km)
	}
	tmp.Zero = learningRate * (float64(1) / float64(m)) * sum0
	tmp.One = learningRate * (float64(1) / float64(m)) * sum1
}

func (t *Thetas) trainLoop(learningRate float64, normalized []datas) {
	for i := 0; i < 500000; i++ {
		var tmp Thetas
		tmp.train(learningRate, t, normalized)
		t.Zero -= tmp.Zero
		t.One -= tmp.One
	}
}

type Limits struct {
	Min float64 `json:"min", csv:"min"`
	Max float64 `json:"max", csv:"max"`
}

func getMax(datas []datas) (max float64) {
	max = math.SmallestNonzeroFloat64
	for _, v := range datas {
		if v.km > max {
			max = v.km
		}
	}
	return
}

func getMin(datas []datas) (min float64) {
	min = math.MaxFloat64
	for _, v := range datas {
		if v.km < min {
			min = v.km
		}
	}
	return
}

func (l *Limits) get(datas []datas) {
	l.Min = getMin(datas)
	l.Max = getMax(datas)
}

func main() {
	var thetas Thetas
	var limits Limits
	f := checkErr(parseFlags()).(flags)
	datas := checkErr(dataReader(f.dataset)).([]datas)
	limits.get(datas)
	normalized := normalizeDatas(datas, limits)
	thetas.trainLoop(f.learningRate, normalized)
	thetasRet, _ := json.Marshal(thetas)
	limitsRet, _ := json.Marshal(limits)

	if err := ioutil.WriteFile(f.tmpFileThetas, thetasRet, 0644); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(f.tmpFileMinMax, limitsRet, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(thetas)
}
