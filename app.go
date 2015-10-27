package main

import (
	"github.com/Wmaxlees/strips-blocks/state"
)

func main() {
	// Initialize the state structure
	startState := state.NewState()
	state.InitStartState(startState, state.InitialStateSimpleProblem)
	startState.Print()
}
