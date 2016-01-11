package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TeddyRilliot/list/program"
	"github.com/TeddyRilliot/list/random"
	"github.com/TeddyRilliot/list/sequential"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	seq := sequential.New([]int{1, 2, 3})
	rnd := random.New([]string{"a", "b", "c"})

	prog := program.New()
	prog.Add(seq, 2)
	prog.Add(rnd, 2)

	for i := 0; i < 12; i++ {
		fmt.Printf("%v ", prog.Next())
	}
}
