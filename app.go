package main

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/opcodes"
	"github.com/Wmaxlees/strips-blocks/state"
	"github.com/Wmaxlees/strips-blocks/stripsstack"
	"time"
)

func main() {
	// Initialize the state structure
	currentState := state.NewState()
	state.InitState(currentState, state.InitialStateSimpleProblem)
	fmt.Println("Start State:")
	currentState.Print()

	goalState := state.NewState()
	state.InitState(goalState, state.GoalStateSimpleProblem)
	fmt.Println("\nGoal State:")
	goalState.Print()

	stack := new(stripsstack.StripsStack)
	stack.Append(goalState.GetPredicates())
	fmt.Println("\nStack:")
	stack.Print()
	fmt.Println()

	for stack.GetLength() > 0 {
		next := stack.Peek()

		if next&uint16(opcodes.NonActionMask) != uint16(opcodes.NonActionMask) {
			fmt.Println("Performing Action")
			currentState.Execute(next)
			stack.Pop()

			fmt.Println("Current State:")
			currentState.Print()
		} else if currentState.CheckPredicate(next) {
			fmt.Println("First item on stack is already in state")
			stack.Pop()

			stack.Print()

			fmt.Println()

		} else {
			fmt.Println("First item on stack doesn't exist in current state")
			possibilities := currentState.FindApplications(next)

			if len(possibilities) > 0 {
				stack.Pop()

				// TODO: Add in branching
				stack.Push(possibilities[0])
				stack.Append(currentState.GetPreconditions(possibilities[0]))
				stack.Print()

				fmt.Println()
			} else {
				fmt.Println("No Possible Moves")
				stack.Print()
				break
			}
		}

		// Sleep
		time.Sleep(3000 * time.Millisecond)
	}
}
