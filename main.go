package main

import (
	//"bball-graph/src/state"
	"bball-graph/src/model"
	//"bball-graph/src/test"
)

func main() {

	//test.TestStateCodec()
	//model.GenDummyData("data.csv")
	model.Import("data.csv")
}