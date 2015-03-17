package main

import (
	"bball-graph/src/model"
	"bball-graph/src/test"
	"flag"
)

var tester, gen, load bool
var file string
func init() {

	flag.BoolVar(&tester, "test", false, "run tests")
	flag.BoolVar(&gen, "gen", false, "generate dummy data")
	flag.BoolVar(&load, "load", false, "load into model")
	flag.StringVar(&file, "file", "", "file to load")

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
			model.Import(file)
		} else {
			model.Import("data.csv")
		}
	}
}