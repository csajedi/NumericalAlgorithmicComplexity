//taken from https://rosettacode.org/wiki/Universal_Turing_machine#Go
package turing
 
type Symbol byte
 
type Motion byte
 
const (
        Left  Motion = 'L'
        Right Motion = 'R'
        Stay  Motion = 'N'
)
 
type Tape struct {
        data      []Symbol
        pos, left int
        blank     Symbol
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
 
func (t *Tape) Stay()          {}
func (t *Tape) Data() []Symbol { return t.data[t.left:] }
func (t *Tape) Read() Symbol   { return t.data[t.pos] }
func (t *Tape) Write(s Symbol) { t.data[t.pos] = s }
 
func (t *Tape) Dup() *Tape {
        t2 := &Tape{
                data:  make([]Symbol, len(t.Data())),
                blank: t.blank,
        }
        copy(t2.data, t.Data())
        t2.pos = t.pos - t.left
        return t2
}
 
func (t *Tape) String() string {
        s := ""
        for i := t.left; i < len(t.data); i++ {
                b := t.data[i]
                if i == t.pos {
                        s += "[" + string(b) + "]"
                } else {
                        s += " " + string(b) + " "
                }
        }
        return s
}
 
func (t *Tape) Move(a Motion) {
        switch a {
        case Left:
                t.Left(1)
        case Right:
                t.Right(1)
        case Stay:
                t.Stay()
        }
}
 
const minSz = 16
 
func (t *Tape) Left(n int) {
        t.pos -= n
        if t.pos < 0 {
                // Extend left
                var sz int
                for sz = minSz; cap(t.data[t.left:])-t.pos >= sz; sz <<= 1 {
                }
                newd := make([]Symbol, sz)
                newl := len(newd) - cap(t.data[t.left:])
                n := copy(newd[newl:], t.data[t.left:])
                t.data = newd[:newl+n]
                t.pos += newl - t.left
                t.left = newl
        }
        if t.pos < t.left {
                if t.blank != 0 {
                        for i := t.pos; i < t.left; i++ {
                                t.data[i] = t.blank
                        }
                }
                t.left = t.pos
        }
}
 
func (t *Tape) Right(n int) {
        t.pos += n
        if t.pos >= cap(t.data) {
                // Extend right
                var sz int
                for sz = minSz; t.pos >= sz; sz <<= 1 {
                }
                newd := make([]Symbol, sz)
                n := copy(newd[t.left:], t.data[t.left:])
                t.data = newd[:t.left+n]
        }
        if i := len(t.data); t.pos >= i {
                t.data = t.data[:t.pos+1]
                if t.blank != 0 {
                        for ; i < len(t.data); i++ {
                                t.data[i] = t.blank
                        }
                }
        }
}
 
type State string
 
type Rule struct {
        State
        Symbol
        Write Symbol
        Motion
        Next State
}
 
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
 
type Machine struct {
        tape         *Tape
        start, state State
        transition   map[key]action
        l            func(string, ...interface{}) // XXX
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
 
func (m *Machine) Run(input *Tape) (int, *Tape) {
        m.tape = input.Dup()
        m.state = m.start
        for cnt := 0; ; cnt++ {
                if m.l != nil {
                        m.l("%3d %4s: %v\n", cnt, m.state, m.tape)
                }
                sym := m.tape.Read()
                act, ok := m.transition[key{m.state, sym}]
                if !ok {
                        return cnt, m.tape
                }
                m.tape.Write(act.write)
                m.tape.Move(act.Motion)
                m.state = act.next
        }
}
An example program using the above package:

package main
 
import (
        ".." // XXX path to above turing package
        "fmt"
)
 
func main() {
        var incrementer = turing.NewMachine([]turing.Rule{
                {"q0", '1', '1', turing.Right, "q0"},
                {"q0", 'B', '1', turing.Stay, "qf"},
        })
        input := turing.NewTape('B', 0, []turing.Symbol{'1', '1', '1'})
        cnt, output := incrementer.Run(input)
        fmt.Println("Turing machine halts after", cnt, "operations")
        fmt.Println("Resulting tape:", output)
 
        var beaver = turing.NewMachine([]turing.Rule{
                {"a", '0', '1', turing.Right, "b"},
                {"a", '1', '1', turing.Left, "c"},
                {"b", '0', '1', turing.Left, "a"},
                {"b", '1', '1', turing.Right, "b"},
                {"c", '0', '1', turing.Left, "b"},
                {"c", '1', '1', turing.Stay, "halt"},
        })
        cnt, output = beaver.Run(turing.NewTape('0', 0, nil))
        fmt.Println("Turing machine halts after", cnt, "operations")
        fmt.Println("Resulting tape:", output)
 
        beaver = turing.NewMachine([]turing.Rule{
                {"A", '0', '1', turing.Right, "B"},
                {"A", '1', '1', turing.Left, "C"},
                {"B", '0', '1', turing.Right, "C"},
                {"B", '1', '1', turing.Right, "B"},
                {"C", '0', '1', turing.Right, "D"},
                {"C", '1', '0', turing.Left, "E"},
                {"D", '0', '1', turing.Left, "A"},
                {"D", '1', '1', turing.Left, "D"},
                {"E", '0', '1', turing.Stay, "H"},
                {"E", '1', '0', turing.Left, "A"},
        })
        cnt, output = beaver.Run(turing.NewTape('0', 0, nil))
        fmt.Println("Turing machine halts after", cnt, "operations")
        fmt.Println("Resulting tape has", len(output.Data()), "cells")
 
        var sort = turing.NewMachine([]turing.Rule{
                // Moving right, first b→B;s1
                {"s0", 'a', 'a', turing.Right, "s0"},
                {"s0", 'b', 'B', turing.Right, "s1"},
                {"s0", ' ', ' ', turing.Left, "se"},
                // Conintue right to end of tape → s2
                {"s1", 'a', 'a', turing.Right, "s1"},
                {"s1", 'b', 'b', turing.Right, "s1"},
                {"s1", ' ', ' ', turing.Left, "s2"},
                // Continue left over b.  a→b;s3, B→b;se
                {"s2", 'a', 'b', turing.Left, "s3"},
                {"s2", 'b', 'b', turing.Left, "s2"},
                {"s2", 'B', 'b', turing.Left, "se"},
                // Continue left until B→a;s0
                {"s3", 'a', 'a', turing.Left, "s3"},
                {"s3", 'b', 'b', turing.Left, "s3"},
                {"s3", 'B', 'a', turing.Right, "s0"},
                // Move to tape start → halt
                {"se", 'a', 'a', turing.Left, "se"},
                {"se", ' ', ' ', turing.Right, "see"},
        })
        input = turing.NewTape(' ', 0, []turing.Symbol("abbabbabababab"))
        cnt, output = sort.Run(input)
        fmt.Println("Turing machine halts after", cnt, "operations")
        fmt.Println("Resulting tape:", output)
}