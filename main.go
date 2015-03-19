package main

import (
	"bball-graph/src/model"
	"bball-graph/src/naive_model"
	"bball-graph/src/test"
	"flag"
)

var tester, gen, load, naive bool
var file string
func init() {

	flag.BoolVar(&tester, "test", false, "run tests")
	flag.BoolVar(&gen, "gen", false, "generate dummy data")
	flag.BoolVar(&load, "load", false, "load into model")
	flag.StringVar(&file, "file", "", "file to load")
	flag.BoolVar(&naive, "naive", false, "load into naive model")

}

func main() {

	flag.Parse()

	if tester {
		test.TestStateCodec()
	} 
	if gen {
		model.SmartGenDummyData("data.csv")
	}
	if load {
		if file != "" {
			imp(naive, file)
		} else {
			imp(naive, "data.csv")
		}
	}
}

func imp(naive bool, file string) {
	if naive {
		naive_model.Import(file)
	} else {
		model.Import(file)
	}
}