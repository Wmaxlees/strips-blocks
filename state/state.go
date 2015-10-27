package state

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/bitcodes"
)

type State struct {
	predicates []uint16
}

func NewState() *State {
	// Initialize the predicates
	predicates := make([]uint16, 0)
	return &State{predicates: predicates}
}

// TODO: Make this more efficient
func (state *State) Print() {
	blockStacks := make([][]uint16, 0)
	// Get all the blocks on the table
	for _, pred := range state.predicates {
		if (pred&bitcodes.OpCodeMask) == bitcodes.OnCode && (pred&bitcodes.Arg2Mask) == bitcodes.FloorArg2 {
			blockStacks = append(blockStacks, []uint16{(pred & bitcodes.Arg1Mask) >> 6})
		}
	}

	// Stack all the blocks on top of these
	for i, stack := range blockStacks {
		top := stack[len(stack)-1]
		for _, pred := range state.predicates {
			if pred&bitcodes.Arg2Mask == top {
				blockStacks[i] = append(stack, (pred&bitcodes.Arg1Mask)>>6)
			}
		}
	}

	// Get the highest stack
	stackHeight := 0
	for _, stack := range blockStacks {
		if stackHeight < len(stack)-1 {
			stackHeight = len(stack) - 1
		}
	}

	// Print out the stacks from the top down
	for i := stackHeight; i >= 0; i-- {
		for j := 0; j < len(blockStacks); j++ {
			if len(blockStacks[j])-1 < i {
				fmt.Print("    ")
			} else {
				fmt.Printf("[%c] ", blockStacks[j][i]+64)
			}
		}
		fmt.Println()
	}
}

type StateDefinition byte

const (
	InitialStateSimpleProblem StateDefinition = iota
	SussmanAnomalyProblem     StateDefinition = iota
)

func InitStartState(state *State, initialState StateDefinition) {
	switch initialState {
	case InitialStateSimpleProblem:
		// First stack
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block01Arg1|bitcodes.FloorArg2)
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block02Arg1|bitcodes.Block01Arg2)

		// Second stack
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block03Arg1|bitcodes.FloorArg2)

		// Third stack
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block04Arg1|bitcodes.FloorArg2)
	case SussmanAnomalyProblem:
		// First stack
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block01Arg1|bitcodes.FloorArg2)
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block03Arg1|bitcodes.Block01Arg2)

		// Second stack
		state.predicates = append(state.predicates, bitcodes.OnCode|bitcodes.Block02Arg1|bitcodes.FloorArg2)
	default:
	}
}
