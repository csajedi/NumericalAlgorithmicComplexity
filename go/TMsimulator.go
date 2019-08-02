package main

/** Rule Components **/
type Symbol byte

type Motion byte

const minSz = 16

type Rule struct {
	State
	Symbol
	Write Symbol
	Motion
	Next State
}

const (
	Left  Motion = 'L'
	Right Motion = 'R'
	Stay  Motion = 'S'
)

func (i *Rule) key() key       { return key{i.State, i.Symbol} }
func (i *Rule) action() action { return action{i.Write, i.Motion, i.Next} }

type key struct {
	State
	Symbol
}

type action struct {
	write Symbol
	Motion
	next State
}

/** Tape Components **/

type Tape struct {
	data      []Symbol
	pos, left int
	blank     Symbol
}
type State string

/** Machine Components **/

type Machine struct {
	tape         *Tape
	start, state State
	transition   map[key]action
	l            func(string, ...interface{}) // XXX
}

// NewTape returns a new tape filled with 'data' and position set to 'start'.
// 'start' does not need to be range, the tape will be extended if required.
func NewTape(blank Symbol, start int, data []Symbol) *Tape {
	t := &Tape{
		data:  data,
		blank: blank,
	}
	if start < 0 {
		t.Left(-start)
	}
	t.Right(start)
	return t
}
func NewMachine(rules []Rule) *Machine {
	m := &Machine{transition: make(map[key]action, len(rules))}
	if len(rules) > 0 {
		m.start = rules[0].State
	}
	for _, r := range rules {
		m.transition[r.key()] = r.action()
	}
	return m
}

//corresponds to init_turing_machine in the TMSimulator.c
func enumerateTM(int states, int colors, mpz_t numberTM) []Rule {
	var table []Rule

}
