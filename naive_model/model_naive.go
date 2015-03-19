package naive_model

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strconv"
	"bufio"
	"strings"
	"bball-graph/src/state"
	"math"
)

const (
	NUM_RECORDS = 48*25*60 //25 records per second
)

var m	map[uint32]map[uint32]uint32

func init() {
	m = make(map[uint32]map[uint32]uint32)
	rand.Seed(time.Now().UnixNano())
}

func Import(fp string) {

	f, err := os.Open(fp)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer f.Close()

	var prevState uint32

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

		if prevState != 0 {
			//technically a valid state but implies that
			//everyone is top left corner and no one has ball
			addState(prevState, s)
		}
		prevState = s
	}

	KEY := uint32(739431240)
	fmt.Println(m[KEY])
	outcomes := make(map[uint32]float64)
	ITER := 10000000.0
	for i := 0.0; i < ITER; i++ {
		o := choose(KEY)
		if _, ok := outcomes[o]; ok {
			outcomes[o] += 1.0
		} else {
			outcomes[o] = 1.0
		}
	}
	for k, _ := range outcomes {
		outcomes[k] = outcomes[k] / ITER
	}
	fmt.Println(outcomes)
}	


func addState(prevState, currState uint32) {

	if _, ok := m[prevState]; ok {
		//a list of next states exist
		if _, ok1 := m[prevState][currState]; ok1 {
			//currState is already in the list
			m[prevState][currState] += 1
		} else {
			//currState is not yet in the list
			m[prevState][currState] = 1
		}

	} else {
		//the previous state has not been indexed yet
		m[prevState] = make(map[uint32]uint32)
		m[prevState][currState] = 1
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

	f, err := os.Create(fp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

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


//returns a column given an x position
func col(x float64) uint32 {
	if uint32(x) > state.WIDTH {
		//TODO: do something else
		panic("x position out of bounds")
	}
	return ( uint32(math.Floor(x)) / (state.WIDTH/state.NUM_COLS) )
}

//returns the row given a y position
func row(y float64) uint32 {
	if uint32(y) > state.HEIGHT {
		//TODO: do something else
		panic("y position out of bounds")
	}
	return ( uint32(math.Floor(y)) / (state.HEIGHT/state.NUM_ROWS) )
}

func nextRow(r uint32) uint32 {

	d := rand.Intn(3)
	if d == 0 {
		return r
	}
	if d == 1 {
		if (r + 1) >= state.NUM_ROWS {
			return nextRow(r)
		}
		return r + 1
	}
	if r == 0 {
		return nextRow(r)
	}
	return r - 1

}

func nextCol(c uint32) uint32 {
	
	d := rand.Intn(3)
	if d == 0 {
		return c
	}
	if d == 1 {
		if (c + 1) >= state.NUM_COLS {
			return nextCol(c)
		}
		return c + 1
	}
	if c == 0 {
		return nextCol(c)
	}
	return c - 1

}

//col-> valid x
func x(c uint32) int64 {
	return int64(c*(state.WIDTH/state.NUM_COLS))
}
//row -> valid y
func y(r uint32) int64 {
	return int64(r*(state.HEIGHT/state.NUM_ROWS))
}

//generates random but valid state transitions
//dumb and extremely inefficient
func SmartGenDummyData(fp string) {

	f, err := os.Create(fp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	var p1_r, p1_c, p2_r, p2_c, p3_r, p3_c, p4_r, p4_c, p5_r, p5_c uint32

	for i := 0; i < NUM_RECORDS*25; i++ { //write player data for 25 games

		if i % 30 == 0 {
			p1_r = row(rand.Float64()*float64(state.HEIGHT))
			p1_c = col(rand.Float64()*float64(state.WIDTH))
			p2_r = row(rand.Float64()*float64(state.HEIGHT))
			p2_c = col(rand.Float64()*float64(state.WIDTH))
			p3_r = row(rand.Float64()*float64(state.HEIGHT))
			p3_c = col(rand.Float64()*float64(state.WIDTH))
			p4_r = row(rand.Float64()*float64(state.HEIGHT))
			p4_c = col(rand.Float64()*float64(state.WIDTH))
			p5_r = row(rand.Float64()*float64(state.HEIGHT))
			p5_c = col(rand.Float64()*float64(state.WIDTH))
		} else {
			p1_r = nextRow(p1_r)
			p1_c = nextCol(p1_c)
			p2_r = nextRow(p2_r)
			p2_c = nextCol(p2_c)
			p3_r = nextRow(p3_r)
			p3_c = nextCol(p3_c)
			p4_r = nextRow(p4_r)
			p4_c = nextCol(p4_c)
			p5_r = nextRow(p5_r)
			p5_c = nextCol(p5_c)
		}



		s := ""
		s += strconv.FormatInt(x(p1_c), 10) + "," + strconv.FormatInt(y(p1_r), 10) + ","
		s += strconv.FormatInt(x(p2_c), 10) + "," + strconv.FormatInt(y(p2_r), 10) + ","
		s += strconv.FormatInt(x(p3_c), 10) + "," + strconv.FormatInt(y(p3_r), 10) + ","
		s += strconv.FormatInt(x(p4_c), 10) + "," + strconv.FormatInt(y(p4_r), 10) + ","
		s += strconv.FormatInt(x(p5_c), 10) + "," + strconv.FormatInt(y(p5_r), 10)

		s = s + "\n"
		f.WriteString(s)
	}
}

//performs selection over probability distribution in O(n)
//may not be so bad if state arrays end up small
//may even be better if above condition is met, 
//since there is no initial generation cost
func choose(key uint32) uint32 {

	total := uint32(0)
	a := make([]uint32, len(m[key]))
	a_k := make([]uint32, len(m[key]))
	for k, v := range m[key] {
		total += v
		a = append(a, total)
		a_k = append(a_k, k)
	}
	r := rand.Intn(int(total))
	for i := 0; i < len(a); i++ {
		if r < int(a[i]) {
			return a_k[i]
		}
	}
	return a_k[0]
}
