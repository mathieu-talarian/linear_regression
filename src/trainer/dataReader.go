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

func retData(km, price string) *datas {
	var k, pr *int
	var err error
	if _, err = strconv.Atoi(km); err != nil {
		*k, _ = strconv.Atoi(km)
	}
	if _, err = strconv.Atoi(price); err != nil {
		*pr, _ = strconv.Atoi(price)
	}
	if k != nil && pr != nil {
		return &datas{
			km:    float64(*k),
			price: float64(*pr),
		}
	}
	return nil
}

func dataReader(dataset string) (ret []datas, err error) {
	var csvFile *os.File
	var line []string
	csvFile, err = os.Open(dataset)
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
