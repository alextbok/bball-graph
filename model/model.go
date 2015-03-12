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

/*
type Pair struct {
	state	uint32
	p		float64
}
*/
var m			map[uint32]map[uint32]float64
var alias		map[uint32][]uint32
var prob		map[uint32][]float64

func init() {
	m = make(map[uint32]map[uint32]float64)
	alias = make(map[uint32][]uint32)
	prob = make(map[uint32][]float64)
	rand.Seed(time.Now().UnixNano())
}

func Import(fp string) {

	f, err := os.Open(fp)
	if err != nil {
		fmt.Println("Error:", err)
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
		//fmt.Println(s)
		//state.PrintBinary(s, true)
		if prevState != 0 {
			//technically a valid state but implies that
			//everyone is top left corner and no one has ball
			addState(prevState, s)
		}
		prevState = s
	}
	for k, v := range(m) {
		initAlias(k, v)
	}
	//fmt.Println(len(m))
	//fmt.Println(len(alias))
	//for k, _ := range m {
	//	fmt.Printf("%v: %v\n", k, draw(k))
	//}
	fmt.Println(alias)
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
		m[prevState] = make(map[uint32]float64)
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

	state.PrintBinary(3, true)

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

/*
 * http://www.keithschwarz.com/darts-dice-coins/
 */
func initAlias(key uint32, a map[uint32]float64) {

	K := len(a)
	alias[key] = make([]uint32, K)
	prob[key] = make([]float64, K)
	p := make([]float64, K)

	total := 0.0
	for _, v := range a {
		total += v
	}
	for k, _ := range a {
		a[k] = ( a[k] / total )
	}

	smaller := make([]uint32, K)
	larger := make([]uint32, K)

	i := uint32(0)
	for _, v := range a {
		p[i] = ( v * float64(K) )
		if p[i] < 1.0 {
			smaller = append(smaller, i)
		} else {
			larger = append(larger, i)
		}
	}

	s := len(smaller) - 1
	l := len(larger) - 1
	var small uint32
	var large uint32
	for s > 0 && l > 0 {
		small = smaller[s]
		large = larger[l]
		s = s - 1
		l = l - 1

		prob[key][small] = p[small]
		alias[key][small] = large

		p[large] = p[large] + ( p[small] - 1 )

		if p[large] < 1.0 {
			s += 1
			smaller[s] = large
		} else {
			l += 1
			larger[l] = large
		}

	}

	for i = 0; l > 0; i++ {
		prob[key][larger[i]] = 1
		l = l - 1
	}
	for i = 0; s > 0; i++ {
		prob[key][smaller[i]] = 1
		s = s - 1
	}
}

func draw(key uint32) uint32{
	i := rand.Intn(len(alias[key]))
	h := rand.Float64()
	if h <= prob[key][i] {
		return uint32(i)
	}
	return alias[key][i]
}

























