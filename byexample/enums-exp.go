package main

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
  StateIdle: "idle",
  StateConnected: "connected",
  StateError: "error",
  StateRetrying: "retrying",
}

func (serverState ServerState) String() string {
  return stateName[serverState]
}

func enumsExp() {
}
