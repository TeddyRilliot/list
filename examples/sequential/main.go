package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/trilliot/list/sequential"
)

type item struct {
	ID   int
	Name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var items = []item{
	item{1, "Title 1"},
	item{2, "Title 2"},
	item{3, "Title 3"},
}

func main() {
	seq := sequential.New(items)

	for l := 0; l < 3; l++ {
		fmt.Println("Loop ", l)
		for i := 1; i <= len(items); i++ {
			fmt.Printf("\t%v\n", seq.Next())
		}
	}
}
