package random

import (
	"math/rand"
	"testing"
)

func TestRandList(t *testing.T) {
	l := []int{1, 2, 3, 4}
	r := New(l)

	rand.Seed(1)
	exp := []int{1, 2, 3, 4, 4, 3, 1, 2, 2, 1, 3, 4}

	for i := 0; i < len(exp); i++ {
		v := r.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestNotSliceRndList(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected panic, nothing happened")
		}
	}()

	_ = New("not a slice")
}

func TestEmptyRndList(t *testing.T) {
	l := []int{}
	r := New(l)

	exp := []interface{}{nil, nil, nil, nil, nil, nil}

	for i := 0; i < len(exp); i++ {
		v := r.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestResetRndList(t *testing.T) {
	l := []int{1, 2, 3, 4}
	r1 := New(l)

	order1 := [4]interface{}{}
	for i := 0; i < len(l); i++ {
		order1[i] = r1.Next()
	}

	r1.Reset()

	for i := 0; i < len(l); i++ {
		if order1[i] != r1.Next() {
			return
		}
	}

	t.Errorf("%v and %v are identical", order1, r1)
}
