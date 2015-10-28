package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/Wmaxlees/strips-blocks/opcodes"
	"github.com/Wmaxlees/strips-blocks/state"
	"github.com/Wmaxlees/strips-blocks/stripsstack"
	"os"
)

var currentState *state.State
var goalState *state.State
var stack *stripsstack.StripsStack
var verbose bool

func parseCLArgs() (int, bool) {
	probFlagPtr := flag.Int("problem", 0, "The problem to start with\n\t0 - Simple Problem\n\t1 - Sussman Anomaly")
	verbFlagPtr := flag.Bool("verbose", false, "Verbose mode")
	flag.Parse()

	return *probFlagPtr, *verbFlagPtr
}

func main() {
	goalState = new(state.State)
	currentState = new(state.State)

	// Get the command line args
	var problemFlag int
	problemFlag, verbose = parseCLArgs()

	// Initialize the start and end states
	if problemFlag == 0 {
		state.InitState(currentState, state.InitialStateSimpleProblem)
		state.InitState(goalState, state.GoalStateSimpleProblem)
	} else if problemFlag == 1 {
		state.InitState(currentState, state.InitialStateSussmanAnomalyProblem)
		state.InitState(goalState, state.GoalStateSussmanAnomalyProblem)
	} else if problemFlag == 2 {
		state.InitState(currentState, state.InitialStateExtraProblem)
		state.InitState(goalState, state.GoalStateExtraProblem)
	}

	// Display the start and goal state
	fmt.Println("Start State:")
	currentState.Print()
	fmt.Println("\nGoal State:")
	goalState.Print()

	// Initialize the stack
	stack = initStack(goalState)

	run()
}

func initStack(goal *state.State) *stripsstack.StripsStack {
	stack := new(stripsstack.StripsStack)

	// Generate the separate portions of the stack
	fullList := make([]uint16, 0)
	topOfStack := make([][]uint16, 0)
	for _, predicate := range *goal {
		topOfStack = append(topOfStack, []uint16{predicate})
		fullList = append(fullList, predicate)
	}

	// Push the portions to the stack
	stack.Push(fullList)
	stack.Append(topOfStack)

	// Print the stack
	if verbose {
		fmt.Println("\nStack:")
		stack.Print()
		fmt.Println()
	}

	return stack
}

func run() {
	// Initialize the reader
	reader := bufio.NewReader(os.Stdin)
	if !verbose {
		fmt.Println("Press enter to continue...")
		reader.ReadString('\n')
	}

	// Loop through the stack
	for stack.GetLength() > 0 {

		if verbose {
			// User hits return before continuing
			fmt.Println("Press enter to continue...")
			reader.ReadString('\n')
		}

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

		if verbose {
			fmt.Println("\n\nStack:")
			stack.Print()
			fmt.Println()
		}

		if next&opcodes.NonActionMask != opcodes.NonActionMask {
			if verbose {
				fmt.Println("Performing Action")
			}

			currentState.Execute(next)
			stack.Pop()

			fmt.Println()
			fmt.Println("Current State:")
			currentState.Print()

			if !verbose {
				fmt.Println("Press enter to continue...")
				reader.ReadString('\n')
			}
		} else if currentState.CheckPredicate(next) {
			if verbose {
				fmt.Println("First item on stack is already in state")
			}
			stack.Pop()
		} else {
			if verbose {
				fmt.Println("First item on stack doesn't exist in current state")
			}

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
	}
}
