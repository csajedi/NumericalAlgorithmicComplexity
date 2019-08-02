nil
type
  symbol* = char

##  The 'symbol' is a char (256)

type
  direction* = enum
    DIR_LEFT, STOP, DIR_RIGHT


##  possible directions
##  struct which contains the result of a transition of the TM

type
  transition_result* {.bycopy.} = object
    control_state*: cint       ##  new state
    write_symbol*: symbol      ##  written symbol


var dir*: direction


##  the state of the TM

type
  turing_machine_state* {.bycopy.} = object
    control_state*: cint       ##  current state
    head_position*: cint       ##  position of the head in the tape
    max_head_position*: cint   ##  number of the most-right visited cell (>=0)
    min_head_position*: cint   ##  number of the most-left visited cell  (<=0)
    tape_r_size*: cint         ##  size of the right tape
    tape_l_size*: cint         ##  size of the left tape
    tape_r*: ptr symbol         ##  right tape (array of symbols)
    tape_l*: ptr symbol         ##  left tape (array of symbols)


##  General data of the TM

type
  turing_machine* {.bycopy.} = object
    number_colors*: cint       ##  number of symbols
    number_states*: cint       ##  number of states
    halting_state*: cint       ##  halting state
    initial_control_state*: cint ##  initial state
    blank_symbol*: symbol      ##  blank symbol
    state*: turing_machine_state
    transition_table*: ptr ptr transition_result ##  REMOVE ????


proc show_data*(m: ptr turing_machine) =
  discard

## ****************************************************************************
##    Executes a step in the given TM
## ***************************************************************************

proc run_step*(m: ptr turing_machine): cint =
  discard

## *****************************************************************
##   Initialization of a Turing Machine
## ****************************************************************

proc init_turing_machine*(states: cint; colors: cint; blank: cint; numberTM: mpz_t): turing_machine =
  discard

proc delete_state*(m: ptr turing_machine) =
  discard

proc outputTM*(m: ptr turing_machine): string =
  discard

proc main*(argn: cint; argv: ptr cstring): cint =
  discard
