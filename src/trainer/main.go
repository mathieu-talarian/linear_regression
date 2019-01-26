package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mathmoul/linear_regression/src/common"
	"os"
)

func checkErr(dt interface{}, err error) interface{} {
	if err != nil {
		log.Fatal("Aïe ", err)
	}
	return dt
}

func normalizeDatas(dataTab datas, limits common.Limits) (normalized datas) {
	for _, v := range dataTab {
		normalized = append(normalized, v.normalize(limits))
	}
	return
}

func main() {
	var limits common.Limits
	f := checkErr(parseFlags()).(flags)
	datas := checkErr(dataReader(f.dataset)).(datas)
	limits.Get(datas)
	normalized := normalizeDatas(datas, limits)
	thetas := common.TrainLoop(f.learningRate, normalized, global.loops)
	if f.dataOutType == 1 {
		thetasRet, _ := json.Marshal(thetas)
		limitsRet, _ := json.Marshal(limits)

		if err := ioutil.WriteFile(f.tmpFileThetas+".json", thetasRet, 0644); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(f.tmpFileMinMax+".json", limitsRet, 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		file := checkErr(os.OpenFile(f.tmpFileThetas+".csv", os.O_RDWR|os.O_CREATE, os.ModePerm)).(*os.File)
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.WriteAll([][]string{{"theta0", "theta1"}, {fmt.Sprint(thetas.Zero), fmt.Sprint(thetas.One)}})

		file2 := checkErr(os.OpenFile(f.tmpFileMinMax+".csv", os.O_RDWR|os.O_CREATE, os.ModePerm)).(*os.File)
		defer file2.Close()
		writer2 := csv.NewWriter(file2)
		defer writer2.Flush()
		writer2.WriteAll([][]string{{"min", "max"}, {fmt.Sprint(limits.Min), fmt.Sprint(limits.Max)}})
	}

	if f.graph {
		graph(datas, limits, thetas)
	}
}
