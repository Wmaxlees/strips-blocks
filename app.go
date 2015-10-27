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
	state.InitState(currentState, state.InitialStateSussmanAnomalyProblem)
	fmt.Println("Start State:")
	currentState.Print()

	goalState := state.NewState()
	state.InitState(goalState, state.GoalStateSussmanAnomalyProblem)
	fmt.Println("\nGoal State:")
	goalState.Print()

	stack := new(stripsstack.StripsStack)
	fullList := make([]uint16, 0)
	topOfStack := make([][]uint16, 0)
	for _, predicate := range goalState.GetPredicates() {
		topOfStack = append(topOfStack, []uint16{predicate})
		fullList = append(fullList, predicate)
	}
	stack.Push(fullList)
	stack.Append(topOfStack)
	fmt.Println("\nStack:")
	stack.Print()
	fmt.Println()

	for stack.GetLength() > 0 {
		nextSlice := stack.Peek()

		var next uint16
		if len(nextSlice) == 1 {
			next = nextSlice[0]
		} else {
			good := true
			for _, pred := range nextSlice {
				if !currentState.CheckPredicate(pred) {
					stack.Push([]uint16{pred})
					good = false
				}
			}

			if good {
				stack.Pop()
			}

			continue
		}

		fmt.Println("\n\n")
		stack.Print()
		fmt.Println()

		if next&uint16(opcodes.NonActionMask) != uint16(opcodes.NonActionMask) {
			fmt.Println("Performing Action")
			currentState.Execute(next)
			stack.Pop()

			fmt.Println("Current State:")
			currentState.Print()
		} else if currentState.CheckPredicate(next) {
			fmt.Println("First item on stack is already in state")
			stack.Pop()
		} else {
			fmt.Println("First item on stack doesn't exist in current state")
			possibilities := currentState.FindApplications(next)

			if len(possibilities) > 0 {
				stack.Pop()

				// TODO: Add in branching
				stack.Push([]uint16{possibilities[0]})

				fullList := make([]uint16, 0)
				topOfStack := make([][]uint16, 0)
				for _, predicate := range currentState.GetPreconditions(possibilities[0]) {
					topOfStack = append(topOfStack, []uint16{predicate})
					fullList = append(fullList, predicate)
				}
				stack.Push(fullList)
				stack.Append(topOfStack)
			} else {
				fmt.Println("No Possible Moves")
				break
			}
		}

		// Sleep
		time.Sleep(3000 * time.Millisecond)
	}
}
