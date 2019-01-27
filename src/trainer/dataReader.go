package main

import (
	"encoding/csv"
	"io"
	"math"
	"mathmoul/linear_regression/src/common"
	"os"
	"strconv"
)

type data struct {
	km    float64
	price float64
}

type datas []data

func (d datas) GetMax() (max float64) {
	max = math.SmallestNonzeroFloat64
	for _, v := range d {
		if v.km > max {
			max = v.km
		}
	}
	return
}

func (d datas) GetMin() (min float64) {
	min = math.MaxFloat64
	for _, v := range d {
		if v.km < min {
			min = v.km
		}
	}
	return
}

func (d datas) Len() int {
	return len(d)
}

func (d datas) Sums(t common.Thetas) (sum0 float64, sum1 float64) {
	for _, v := range d {
		dt := v.estimatePrice(t)
		sum0 += dt
		sum1 += (dt * v.km)
	}
	return
}

func (d data) normalize(limits common.Limits) data {
	return data{
		km:    (d.km - limits.Min) / (limits.Max - limits.Min),
		price: d.price,
	}
}

func (d data) estimatePrice(t common.Thetas) float64 {
	return (t.Zero + (t.One * d.km)) - d.price
}

func retData(km, price string) *data {
	var k, p int
	var err error
	if k, err = strconv.Atoi(km); err != nil {
		return nil
	}
	if p, err = strconv.Atoi(price); err != nil {
		return nil
	}
	return &data{
		km:    float64(k),
		price: float64(p),
	}
}

func dataReader(file string) (ret datas, err error) {
	var csvFile *os.File
	var line []string
	csvFile, err = os.Open(file)
	if err != nil {
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	for {
		line, err = reader.Read()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}
		retDt := retData(line[0], line[1])
		if retDt != nil {
			ret = append(ret, *retDt)
		}
	}
	return
}
