package test

import (
	"bball-graph/src/state"
	"fmt"
)

func TestStateCodec() {
	p1 :=  state.EnPlayerState(false, 0, 0) 		//(0,0)
	p2 :=  state.EnPlayerState(false, 22, 2) 		//(2,0)
	p3 :=  state.EnPlayerState(true, 39, 24) 		//(3,4)
	p4 :=  state.EnPlayerState(false, 33, 7)		//(3,1)
	p5 :=  state.EnPlayerState(false, 45.3, 3.5)	//(4,0)

	s := state.EnCourtState(p1, p2, p3, p4, p5)
	state.PrintBinary(s, true)

	fmt.Println( state.DecPlayerState(s, 1) ) //false 0 0
	fmt.Println( state.DecPlayerState(s, 2) ) //false 0 2
	fmt.Println( state.DecPlayerState(s, 3) ) //true 4 3
	fmt.Println( state.DecPlayerState(s, 4) )	//false 1 3
	fmt.Println( state.DecPlayerState(s, 5) ) //false 0 4
}