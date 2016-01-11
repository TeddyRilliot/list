package program

import (
	"math/rand"
	"testing"

	"github.com/TeddyRilliot/list/random"
	"github.com/TeddyRilliot/list/sequential"
)

func TestProgramRules2x(t *testing.T) {
	p := New()
	p.Add(sequential.New([]int{1, 2, 3}), 2)
	p.Add(random.New([]string{"a", "b", "c"}), 2)

	rand.Seed(1)
	exp := []interface{}{1, 2, "a", "b", 3, 1, "c", "a", 2, 3, "b", "c", 1, 2, "c", "b"}

	for i := 0; i < len(exp); i++ {
		v := p.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestProgramRules1x(t *testing.T) {
	p := New()
	p.Add(sequential.New([]int{1, 2, 3}), 1)
	p.Add(random.New([]string{"a", "b", "c"}), 1)

	rand.Seed(1)
	exp := []interface{}{1, "b", 2, "c", 3, "a", 1, "a", 2, "b", 3, "c", 1, "c", 2, "b"}

	for i := 0; i < len(exp); i++ {
		v := p.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestProgram1Rule2x(t *testing.T) {
	p := New()
	p.Add(sequential.New([]int{1, 2, 3}), 2)

	exp := []interface{}{1, 2, 3, 1, 2, 3, 1, 2, 3}

	for i := 0; i < len(exp); i++ {
		v := p.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestProgramRule1x(t *testing.T) {
	p := New()
	p.Add(sequential.New([]int{1, 2, 3}), 1)

	exp := []interface{}{1, 2, 3, 1, 2, 3, 1, 2, 3}

	for i := 0; i < len(exp); i++ {
		v := p.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestEmptyProgram(t *testing.T) {
	p := New()

	exp := []interface{}{nil, nil, nil, nil, nil, nil}

	for i := 0; i < len(exp); i++ {
		v := p.Next()
		t.Logf("Loop %d, got %v", i, v)

		if v != exp[i] {
			t.Errorf("Expecting %v, got %v", exp[i], v)
		}
	}
}

func TestInvalidRuleNb(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected panic, nothing happened")
		}
	}()

	p := New()
	p.Add(sequential.New([]int{1}), 0)
}
