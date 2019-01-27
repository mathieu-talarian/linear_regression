package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mathmoul/linear_regression/src/common"
	"os"
	"strconv"
)

func checkErr(dt interface{}, err error) interface{} {
	if err != nil {
		log.Fatal("Aïe ", err)
	}
	return dt
}

func thetasData(zero, one string) (*float64, *float64, error) {
	zz, err := strconv.ParseFloat(zero, 1)
	if err != nil {
		return nil, nil, err
	}
	oo, err := strconv.ParseFloat(one, 1)
	if err != nil {
		return nil, nil, err
	}
	return &zz, &oo, nil
}

func limitsData(zero, one string) (*float64, *float64, error) {
	min, err := strconv.ParseFloat(zero, 1)
	if err != nil {
		return nil, nil, err
	}
	max, err := strconv.ParseFloat(one, 1)
	if err != nil {
		return nil, nil, err
	}
	return &min, &max, nil
}

func openThetas(thetas *common.Thetas) (err error) {
	var line []string
	csvFile, err := os.Open(global.tmpFileThetas)
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
		if z, o, err := thetasData(line[0], line[1]); err != nil {
			log.Println(err)
			continue
		} else {
			if z != nil && o != nil {
				thetas.Zero = *z
				thetas.One = *o
			}
		}
	}
	return
}

func openMinMax(limits *common.Limits) (err error) {
	var line []string
	csvFile, err := os.Open(global.tmpFileMinMax)
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
		if min, max, err := limitsData(line[0], line[1]); err != nil {
			log.Println(err)
			continue
		} else {
			if min != nil && max != nil {
				limits.Min = *min
				limits.Max = *max
			}
		}
	}
	return
}

func ask(kilometrage *float64) (err error) {
	var s string
	fmt.Println("Insert kilometrage: ")
	if _, err = fmt.Scan(&s); err != nil {
		return
	}
	if *kilometrage, err = strconv.ParseFloat(s, 1); err != nil {
		return
	}
	return
}

func normalize(kilometrage float64, limits common.Limits) float64 {
	return (kilometrage - limits.Min) / (limits.Max - limits.Min)
}

func estimate(normalized float64, thetas common.Thetas) float64 {
	return thetas.Zero + thetas.One*normalized
}

func main() {
	var thetas common.Thetas
	var limits common.Limits
	var kilometrage float64
	_ = checkErr(nil, parseFlags())
	if err := openThetas(&thetas); err != nil {
		log.Println(err)
		thetas.SetZero()
	}
	fmt.Println("Thetas =>", thetas)
	if err := openMinMax(&limits); err != nil {
		log.Println(err)
		limits.SetZero()
	}
	fmt.Println("Limits =>", limits)
	if err := ask(&kilometrage); err != nil {
		log.Fatal(err)
	}
	fmt.Println("estimated price => ", estimate(normalize(kilometrage, limits), thetas))
}
