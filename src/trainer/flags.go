package main

import (
	"flag"
	"fmt"
)

type flags struct {
	dataset       string
	learningRate  float64
	tmpFileThetas string
	tmpFileMinMax string
	dataOutType   int
	graph         bool
}

var global flags

func parseFlags() (f flags, err error) {
	flag.StringVar(&f.dataset, "d", "data.csv", "name of the dataset file")
	flag.Float64Var(&f.learningRate, "l", 0.1, "learning rate")
	flag.StringVar(&f.tmpFileThetas, "t", ".thetas", "The file thetas output")
	flag.StringVar(&f.tmpFileMinMax, "m", ".min_max", "The file min max output")
	flag.IntVar(&f.dataOutType, "o", 0, "File output type\nDefault (0) = csv, \n(1) = json")
	flag.BoolVar(&f.graph, "g", false, "Graph required")
	flag.Parse()
	global = f
	if len(flag.Args()) != 0 {
		return f, fmt.Errorf("Too much arguments")
	}
	return
}
