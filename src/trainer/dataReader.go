package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type datas struct {
	km    float64
	price float64
}

func (d datas) normalize(limits Limits) datas {
	return datas{
		km:    (d.km - limits.Min) / (limits.Max - limits.Min),
		price: d.price,
	}
}

func (d datas) estimatePrice(t *Thetas) float64 {
	return (t.Zero + (t.One * d.km)) - d.price
}

func retData(km, price string) *datas {
	var k, p int
	var err error
	if k, err = strconv.Atoi(km); err != nil {
		return nil
	}
	if p, err = strconv.Atoi(price); err != nil {
		return nil
	}
	return &datas{
		km:    float64(k),
		price: float64(p),
	}
}

func dataReader(dataset string) (ret []datas, err error) {
	var csvFile *os.File
	var line *[]string
	csvFile, err = os.Open(dataset)
	if err != nil {
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	for {
		*line, err = reader.Read()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}
		retDt := retData((*line)[0], (*line)[1])
		if retDt != nil {
			ret = append(ret, *retDt)
		}
	}
	return
}
