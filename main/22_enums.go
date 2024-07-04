package main

import "fmt"

type ServiceState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

const (
	A = "a"
	B = "b"
)

var stateNames = map[ServiceState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (ss ServiceState) String() string {
	return stateNames[ss]
}

func transition(ss ServiceState) ServiceState {
	switch ss {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic("invalid service state")
	}
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns = transition(ns)
	fmt.Println(ns)
}
