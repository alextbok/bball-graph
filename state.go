package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 * TODO: 
 * 1) fix bounds checking (do something useful)
 *
 */
const (
	NUM_ROWS uint32 = 5
	NUM_COLS uint32 = 5
	WIDTH uint32 = 50
	HEIGHT uint32 = 25
	BIT_NUM uint32 = 6
	BALL_MASK uint32 = uint32(1 << 5)
 )


func main() {

	p1 :=  enPlayerState(false, 0, 0) 		//(0,0)
	p2 :=  enPlayerState(false, 22, 2) 		//(2,0)
	p3 :=  enPlayerState(true, 39, 24) 		//(3,4)
	p4 :=  enPlayerState(false, 33, 7)		//(3,1)
	p5 :=  enPlayerState(false, 45.3, 3.5)	//(4,0)

	s := enCourtState(p1, p2, p3, p4, p5)
	printBinary(s, true)

	fmt.Println( decPlayerState(s, 1) ) //false 0 0
	fmt.Println( decPlayerState(s, 2) ) //false 0 2
	fmt.Println( decPlayerState(s, 3) ) //true 4 3
	fmt.Println( decPlayerState(s, 4) )	//false 1 3
	fmt.Println( decPlayerState(s, 5) ) //false 0 4
	
}

//returns a column given an x position
func col(x float64) uint32 {
	if uint32(x) > WIDTH {
		//TODO: do something else
		panic("x position out of bounds")
	}
	return ( uint32(math.Floor(x)) / (WIDTH/NUM_COLS) )
}

//returns the row given a y position
func row(y float64) uint32 {
	if uint32(y) > HEIGHT {
		//TODO: do something else
		panic("y position out of bounds")
	}
	return ( uint32(math.Floor(y)) / (HEIGHT/NUM_ROWS) )
}

//encodes state of individual player
func enPlayerState(ball bool, x, y float64) uint32 {

	if ball {
		return ( uint32(row(y)*NUM_COLS + col(x)) | BALL_MASK )
	}

	return uint32(row(y)*NUM_COLS + col(x))
}

//encodes a court state with individual player states
func enCourtState(p1, p2, p3, p4, p5 uint32) uint32 {

	return (
		uint32(0)	|
		(p1 << 26)	|
		(p2 << 20)	|
		(p3 << 14)	|
		(p4 << 8)	|
		(p5 << 2))
}

//prints a number in its binary form
//pads with zeros to print 32 bits if b is set
func printBinary(u uint32, b bool) {

	p := strconv.FormatInt( int64(u), 2 )

	if b {
		l := len(p)
		if l < 32 {
			for i := 0; i < (32 - l); i++ {
				p = "0" + p
			}
		}
	}

	fmt.Printf("%s\n", p)
}

//decodes a player's state
//returns ( <player has ball> ? true : false ), row, col
func decPlayerState(cstate uint32, pnum uint32) (ball bool, row uint32, col uint32) {

	//mask off bits and shift back
	pstate := cstate & uint32(0x3F << (BIT_NUM * (5 - pnum) + 2)) >> (BIT_NUM * (5 - pnum) + 2) 
	//mask to get first bit
	ball = ((pstate & BALL_MASK) == BALL_MASK)

	pstate = (pstate & ^BALL_MASK)
	row = ( pstate / NUM_COLS )
	col = ( pstate % NUM_COLS )

	return
}


