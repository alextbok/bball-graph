package model

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strconv"
	"bufio"
	"strings"
	"bball-graph/src/state"
)

const (
	NUM_RECORDS = 48*25*60 //25 records per second
)

type Pair struct {
	state	uint32
	p		float64
}

var m map[uint32][]uint32

func Import(fp string) {

	f, err := os.Open(fp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	var prevState uint32
	m = make(map[uint32][]uint32)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines) 
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), ",")
		p1 :=  state.EnPlayerState(false, toFloat64(a[0]), toFloat64(a[1]))
		p2 :=  state.EnPlayerState(false, toFloat64(a[2]), toFloat64(a[3]))
		p3 :=  state.EnPlayerState(false, toFloat64(a[4]), toFloat64(a[5]))
		p4 :=  state.EnPlayerState(false, toFloat64(a[6]), toFloat64(a[7]))
		p5 :=  state.EnPlayerState(false, toFloat64(a[8]), toFloat64(a[9]))
		s := state.EnCourtState(p1, p2, p3, p4, p5)
		//fmt.Println(s)
		//state.PrintBinary(s, true)
		fmt.Println(prevState)
		break
		if prevState != 0 {
			//technically a valid state but implies that
			//everyone is top left corner and no one has ball
			addState(prevState, s)
		}
		prevState = s
	}

}

func addState(prevState, currState uint32) {

	if val, ok := m[prevState]; ok {
		//TODO:something useful
		fmt.Println(val)
		return

	} else {
		//TODO:something useful
		return
	}

}

func toFloat64(s string) float64 {
	u, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("error converting %s to float64.")
		return 0
	}
	return u
}

func GenDummyData(fp string) {

	state.PrintBinary(3, true)

	f, err := os.Create(fp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUM_RECORDS*25; i++ { //write player data for 25 games
		s := ""
		for j := 0; j < 5; j++ {
			x := rand.Float64()*float64(state.WIDTH)
			y := rand.Float64()*float64(state.HEIGHT)
			s += strconv.FormatInt(int64(x), 10) + "," + strconv.FormatInt(int64(y), 10) + ","
		}
		s = s[:len(s) - 1] + "\n"
		f.WriteString(s)
	}
}





