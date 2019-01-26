package main

import (
	"flag"
	"fmt"
)

type flags struct {
	tmpFileThetas string
	tmpFileMinMax string
}

var global flags

func parseFlags() (err error) {
	var f flags
	flag.StringVar(&f.tmpFileThetas, "t", ".thetas.csv", "The file thetas output")
	flag.StringVar(&f.tmpFileMinMax, "m", ".min_max.csv", "The file min max output")
	flag.Parse()
	global = f
	if len(flag.Args()) != 1 {
		return fmt.Errorf("Not enough arguments, mileage needed")
	}
	return
}
