package sequential

import (
	"fmt"
	"reflect"
)

type seqList struct {
	list   []interface{}
	offset int
}

// New creates a new sequential list with items.
func New(l interface{}) *seqList {
	if reflect.TypeOf(l).Kind() != reflect.Slice {
		panic(fmt.Errorf("Param must be a slice"))
	}

	s := reflect.ValueOf(l)
	items := make([]interface{}, s.Len())

	for i := 0; i < len(items); i++ {
		items[i] = s.Index(i).Interface()
	}

	return &seqList{
		list:   items,
		offset: 0,
	}
}

// Next returns the following item in the sequential list.
// If the last item of the list were returned previously, the list restarts
// from the first item.
func (s *seqList) Next() interface{} {
	nb := len(s.list)

	if nb == 0 {
		return nil
	}
	if s.offset >= nb {
		s.Reset()
	}

	defer func() {
		s.offset++
	}()
	return s.list[s.offset]
}

// Reset sets the position of the list to its first item.
// The following call to Next() will return the first item of the list.
func (s *seqList) Reset() {
	s.offset = 0
}
