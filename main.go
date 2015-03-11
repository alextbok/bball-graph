package main

import (
	"bball-graph/src/state"
	"bball-graph/src/model"
	//"bball-graph/src/test"
)

func main() {

	state.PrintBinary(10, true)
	//test.TestStateCodec()
	model.GenDummyData("data.csv")
}