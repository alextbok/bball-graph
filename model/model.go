package model

import (
	"fmt"
	"os"
	"bball-graph/src/state"
	"math/rand"
	"time"
	"strconv"
)

const (
	NUM_RECORDS = 48*25*60 //25 records per second
)

var m map[uint32]uint32

func Import(fp string) {
	//for each line in file:
		//read into array A

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
	for i := 0; i < NUM_RECORDS*50; i++ { //write player data for 60 games
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





