package main

import "log"

func checkErr(dt interface{}, err error) interface{} {
	if err != nil {
		log.Fatal("Aïe ", err)
	}
	return dt
}

func openThetas() {

}

func main() {
	_ = checkErr(nil, parseFlags())
	// openThetas()
	// openMinMax()
}
