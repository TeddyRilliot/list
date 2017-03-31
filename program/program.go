package program

import "github.com/trilliot/list"

type rule struct {
	list.Interface
	Max int
}

// Program is a type of list that contains rules.
// A rule is a list (no matter what kind of list), that may be repeated
// multiple times.
// If a rule have to be repeated, all of its occurrences have to happen
// before the following rule is played.
type Program struct {
	rules      []*rule
	offset     int
	ruleOffset int
}

// New creates an emtpy program.
func New() *Program {
	return new(Program)
}

// Add inserts a list.Interface in the program.
// The nb parameters controls how many times the rule will be played before
// passing to the next one.
func (p *Program) Add(l list.Interface, nb int) {
	if nb < 1 {
		panic("Rule should be played at least once.")
	}
	r := &rule{l, nb}
	p.rules = append(p.rules, r)
}

// Next returns the following element in the program.
// This method takes care of repeating the rule multiple times if needed.
// When all rules are played, the program restarts from the beginning.
func (p *Program) Next() interface{} {
	nb := len(p.rules)

	if nb == 0 {
		return nil
	}
	if p.offset >= nb {
		p.Reset()
	}

	r := p.rules[p.offset]

	if p.ruleOffset >= r.Max {
		p.offset++
		p.ruleOffset = 0
		return p.Next()
	}

	defer func() {
		p.ruleOffset++
	}()
	return r.Next()
}

// Reset sets the position of the program to its first rule.
// The following call to Next() will return the first occurence of the rule
// (if this rule should be repeated multiple times).
func (p *Program) Reset() {
	p.offset = 0
	p.ruleOffset = 0
}
