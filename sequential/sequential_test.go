package sequential

import "testing"

func TestSeqList(t *testing.T) {
	l := []int{1, 2, 3, 4}
	s := New(l)

	exp := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}

	for i := 0; i < 12; i++ {
		v := s.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
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

	exp := []interface{}{nil, nil, nil, nil, nil, nil}
	for i := 0; i < 6; i++ {
		v := s.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestResetSeqList(t *testing.T) {
	l := []int{1, 2, 3}
	s := New(l)

	first := s.Next()
	s.Reset()
	second := s.Next()

	if second != l[0] {
		t.Errorf("Reset failed: %v != %v", second, l[0])
	}
	if first != second {
		t.Errorf("Next failed: %v != %v", first, second)
	}
}
