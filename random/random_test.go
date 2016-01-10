package random

import "testing"

func TestRandList(t *testing.T) {
	l := []int{1, 2, 3, 4}
	r := New(l)

	for i := 0; i < 2; i++ {
		for _, exp := range l {
			v := r.Next()

			// Value is not equal to the sequential one : the list is
			// randomized, test succeeds.
			if v != exp {
				return
			}
		}
	}

	t.Errorf("list %v is not randomized", l)
}

func TestNotSliceRndList(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected panic, nothing happened")
		}
	}()

	_ = New("not a slice")
}

func Test1ItemRandList(t *testing.T) {
	l := []int{1}
	r := New(l)

	for i := 0; i < 2; i++ {
		for _, exp := range l {
			v := r.Next()
			if v != exp {
				t.Errorf("Loop %d, %v is not expected (want %d)", i, v, exp)
				return
			}
		}
	}
}

func TestEmptyRndList(t *testing.T) {
	l := []int{}
	r := New(l)

	for i := 0; i < 2; i++ {
		v := r.Next()
		if v != nil {
			t.Errorf("Loop %d, %v is not expected (want nil)", i, v)
			return
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
			t.Errorf("OK")
			return
		}
	}

	t.Errorf("%v and %v are identical", order1, r1)
}
