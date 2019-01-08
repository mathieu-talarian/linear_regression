package main

import (
	"fmt"
	"log"
)

func checkErr(dt interface{}, err error) interface{} {
	if err != nil {
		log.Fatal("Aïe ", err)
	}
	return dt
}

func main() {
	f := checkErr(parseFlags()).(flags)
	datas := checkErr(dataReader(f.dataset)).([]datas)
	fmt.Println(datas)
}
