package sequential

import "testing"

func TestSeqList(t *testing.T) {
	l := []int{1, 2, 3, 4}
	s := New(l)

	for i := 0; i < 2; i++ {
		for _, exp := range l {
			v := s.Next()
			if v != exp {
				t.Errorf("Loop %d, %d is not expected (want %d)", i, v, exp)
				return
			}
		}
	}
}

func TestNotSliceSeqList(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected panic, nothing happened")
		}
	}()

	_ = New("not a slice")
}

func TestEmptySeqList(t *testing.T) {
	l := []int{}
	s := New(l)

	for i := 0; i < 2; i++ {
		v := s.Next()
		if v != nil {
			t.Errorf("Loop %d, %d is not expected (want nil)", i, v)
			return
		}
	}
}

func TestResetSeqList(t *testing.T) {
	l := []int{1, 2, 3}
	s := New(l)

	first := s.Next()
	s.Reset()
	second := s.Next()

	if first != second || first != l[0] {
		t.Errorf("%d != %d (reset failed), or %d != %d (next failed)", first, second, first, l[0])
		return
	}
}
