// Code generated by go-enum
// DO NOT EDIT!

package api

import (
	"fmt"
)

const (
	// StateWaiting is a State of type Waiting
	StateWaiting State = iota + 1
	// StateRunning is a State of type Running
	StateRunning
	// StateFailed is a State of type Failed
	StateFailed
	// StateSucceeded is a State of type Succeeded
	StateSucceeded
)

const _StateName = "WaitingRunningFailedSucceeded"

var _StateMap = map[State]string{
	1: _StateName[0:7],
	2: _StateName[7:14],
	3: _StateName[14:20],
	4: _StateName[20:29],
}

// String implements the Stringer interface.
func (x State) String() string {
	if str, ok := _StateMap[x]; ok {
		return str
	}
	return fmt.Sprintf("State(%d)", x)
}

var _StateValue = map[string]State{
	_StateName[0:7]:   1,
	_StateName[7:14]:  2,
	_StateName[14:20]: 3,
	_StateName[20:29]: 4,
}

// ParseState attempts to convert a string to a State
func ParseState(name string) (State, error) {
	if x, ok := _StateValue[name]; ok {
		return x, nil
	}
	return State(0), fmt.Errorf("%s is not a valid State", name)
}
