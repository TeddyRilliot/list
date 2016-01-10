package random

import (
	"fmt"
	"math/rand"
	"reflect"
)

type randList struct {
	list   []interface{}
	perm   []int
	offset int
}

// New creates a new random list with items.
// The randomization method is smart, a returned value will not be returned
// again until all values available are returned at least once.
func New(l interface{}) *randList {
	if reflect.TypeOf(l).Kind() != reflect.Slice {
		panic(fmt.Errorf("Param must be a slice"))
	}

	s := reflect.ValueOf(l)
	items := make([]interface{}, s.Len())

	for i := 0; i < len(items); i++ {
		items[i] = s.Index(i).Interface()
	}

	return &randList{
		list:   items,
		perm:   rand.Perm(len(items)),
		offset: 0,
	}
}

// Next returns the following element in the random list.
// When all items are returned once, a call to Reset is made to compute a new
// random order.
func (r *randList) Next() interface{} {
	nb := len(r.perm)

	if nb == 0 {
		return nil
	}
	if r.offset >= nb {
		r.Reset()
	}

	defer func() {
		r.offset++
	}()

	i := r.perm[r.offset]
	return r.list[i]
}

// Reset discards the current random order and computes a new one.
// When calling Random, you cannot ensure returned items to be smart, regarding
// what were returned before.
// New random indexes will be generated no matter what items were returned
// before. Don't be surprised if you see the same item two times in a row if
// you Reset in between !
func (r *randList) Reset() {
	r.offset = 0
	r.perm = rand.Perm(len(r.list))
}
