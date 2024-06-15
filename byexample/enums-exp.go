package main

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (serverState ServerState) String() string {
	return stateName[serverState]
}

func enumsExp() {
	fmt.Printf("transition from %s to %s\n", ServerState(StateIdle), stateTransition(StateIdle))
	fmt.Printf("transition from %s to %s\n", ServerState(StateConnected), stateTransition(StateConnected))
	fmt.Printf("transition from %s to %s\n", ServerState(StateError), stateTransition(StateError))
}

func stateTransition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
