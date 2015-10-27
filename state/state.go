package state

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
	"github.com/Wmaxlees/strips-blocks/opcodes"
)

type State struct {
	predicates []uint16
}

func NewState() *State {
	// Initialize the predicates
	predicates := make([]uint16, 0)
	return &State{predicates: predicates}
}

func (state *State) GetPredicates() []uint16 {
	return state.predicates
}

// TODO: Make this more efficient
func (state *State) Print() {
	blockStacks := make([][]uint16, 0)
	var holding uint16 = 0
	// Get all the blocks on the table and find if there is a block behing held
	for _, pred := range state.predicates {
		if (pred&uint16(opcodes.OpCodeMask)) == uint16(opcodes.OnOpCode) && (pred&uint16(argtwo.ArgTwoMask)) == uint16(argtwo.Floor) {
			blockStacks = append(blockStacks, []uint16{(pred & uint16(argone.ArgOneMask)) >> 6})
		}

		if (pred & uint16(opcodes.OpCodeMask)) == uint16(opcodes.HoldingOpCode) {
			holding = pred & uint16(argone.ArgOneMask) >> 6
		}
	}

	// Stack all the blocks on top of these
	for i, stack := range blockStacks {
		top := stack[len(stack)-1]
		for _, pred := range state.predicates {
			if pred&uint16(argtwo.ArgTwoMask) == top {
				blockStacks[i] = append(stack, (pred&uint16(argone.ArgOneMask))>>6)
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

	// Print the holding indication
	if holding != 0 {
		fmt.Printf("Holding: [%c]\n", holding+64)
	}
}

type StateDefinition byte

const (
	InitialStateSimpleProblem         StateDefinition = iota
	InitialStateSussmanAnomalyProblem StateDefinition = iota
	GoalStateSimpleProblem            StateDefinition = iota
	GoalStateSussmanAnomalyProblem    StateDefinition = iota
)

func (state *State) AddPredicate(opCode opcodes.OpCode, argOne argone.ArgOne, argTwo argtwo.ArgTwo) {
	state.predicates = append(state.predicates, uint16(opCode)|uint16(argOne)|uint16(argTwo))
}

func InitState(state *State, initialState StateDefinition) {
	switch initialState {
	case InitialStateSimpleProblem:
		// First stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockA, argtwo.Floor)    // OnFloor(A)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockB, argtwo.BlockA)   // On(B, A)
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockB, argtwo.Blank) // Clear(B)

		// Second stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockC, argtwo.Floor)    // OnFloor(C)
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockC, argtwo.Blank) // Clear(C)

		// Third stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockD, argtwo.Floor)    // OnFloor(D)
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockD, argtwo.Blank) // Clear(D)

		// Holding
		state.AddPredicate(opcodes.HoldingOpCode, argone.Blank, argtwo.Blank) // Holding(nil)

	case InitialStateSussmanAnomalyProblem:
		// First stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockA, argtwo.Floor)    // OnFloor(A)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockC, argtwo.BlockA)   // On(C, A)
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockC, argtwo.Blank) // Clear(C)

		// Second stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockB, argtwo.Floor)    // OnFloor(B)
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockB, argtwo.Blank) // Clear(B)

		// Holding
		state.AddPredicate(opcodes.HoldingOpCode, argone.Blank, argtwo.Blank) // Holding(nil)

	case GoalStateSimpleProblem:
		// First stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockA, argtwo.Floor)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockC, argtwo.BlockA)

		// Second stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockD, argtwo.Floor)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockB, argtwo.BlockD)

	case GoalStateSussmanAnomalyProblem:
		// Stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockC, argtwo.Floor)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockB, argtwo.BlockC)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockA, argtwo.BlockB)

	default:
	}
}

// func (state *State) CheckHolding(block argone.ArgOne) bool {
// 	for pred := range state.predicates {
// 		// Get components
// 		opCode, argOne, _ := opcodes.GetComponents(uint16(pred))

// 		// Check if we have a Holding predicate
// 		if opCode == uint16(opcodes.HoldingOpCode) {
// 			if argOne != uint16(block) { // Block being held is not our block
// 				return false
// 			} else { // Block behing held is our block
// 				return true
// 			}
// 		}
// 	}

// 	// No holding predicate exists
// 	return false
// }

func (state *State) CheckPredicate(predicate uint16) bool {
	for _, existingPred := range state.predicates {
		if uint16(existingPred) == predicate {
			return true
		}
	}

	return false
}

// func (state *State) CheckOn(x argone.ArgOne, y argtwo.ArgTwo) bool {
// 	// Generate the predicate
// 	var check uint16 = uint16(opcodes.OnOpCode) & uint16(x) & uint16(y)

// 	// Check if predicate exists
// 	for pred := range state.predicates {
// 		if uint16(pred) == check {
// 			return true
// 		}
// 	}

// 	// Didn't find the predicate
// 	return false
// }

// func (state *State) CheckClear(block argone.ArgOne) bool {
// 	for pred := range state.predicates {
// 		// Get components
// 		opCode, argOne, argTwo := opcodes.GetComponents(uint16(pred))

// 		// Check if there is a Clear predicate
// 		if opCode == uint16(opcodes.ClearOpCode) && argOne == uint16(block) {
// 			return true
// 		}

// 		// Check if there is an on(x,y) predicate
// 		if opCode == uint16(opcodes.OnOpCode) && argTwo == uint16(block)>>6 {
// 			return false
// 		}
// 	}

// 	// If we have reached here there is no On predicate that corresponds to block
// 	return true
// }

func (state *State) findWhatBlockIsOn(block uint16) uint16 {
	for _, pred := range state.predicates {
		opCode, argOne, argTwo := opcodes.GetComponents(pred)

		if opCode == uint16(opcodes.OnOpCode) && argOne == block {
			return argTwo
		}
	}

	return uint16(argtwo.Blank)
}

func (state *State) findWhatIsOnBlock(block uint16) uint16 {
	block = block >> 6

	for _, pred := range state.predicates {
		opCode, argOne, argTwo := opcodes.GetComponents(pred)

		if opCode == uint16(opcodes.OnOpCode) && argTwo == block {
			return argOne
		}
	}

	return uint16(argtwo.Blank)
}

func (state *State) findWhatIsBeingHeld() uint16 {
	for _, pred := range state.predicates {
		opCode, argOne, _ := opcodes.GetComponents(pred)

		if opCode == uint16(opcodes.HoldingOpCode) {
			return argOne
		}
	}

	return 0
}

func (state *State) GetPreconditions(cmd uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(cmd)
	switch opCode {
	case uint16(opcodes.StackOpCode): // Stack(x,y)
		if argTwo != uint16(argtwo.Floor) {
			return []uint16{
				uint16(opcodes.HoldingOpCode) | argOne | uint16(argtwo.Blank), // Holding(x)
				uint16(opcodes.ClearOpCode) | argTwo<<6,                       // Clear(y)
			}
		} else {
			return []uint16{
				uint16(opcodes.HoldingOpCode) | argOne | uint16(argtwo.Blank), // Holding(x)
			}
		}
	case uint16(opcodes.UnstackOpCode): // Unstack(x,y)
		return []uint16{
			uint16(opcodes.OnOpCode) | argOne | argTwo,                                  // On(x,y)
			uint16(opcodes.ClearOpCode) | argOne,                                        // Clear(x)
			uint16(opcodes.HoldingOpCode) | uint16(argone.Blank) | uint16(argtwo.Blank), // Holding(nil)
		}
	case uint16(opcodes.PickupOpCode): // Pickup(x)
		return []uint16{
			uint16(opcodes.ClearOpCode) | argOne,                                        // Clear(x)
			uint16(opcodes.OnOpCode) | argOne | uint16(argtwo.Floor),                    // On(x, Floor)
			uint16(opcodes.HoldingOpCode) | uint16(argone.Blank) | uint16(argtwo.Blank), // Holding(nil)
		}
	default:
		return nil
	}
}

func (state *State) getDeletes(cmd uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(cmd)

	switch opCode {
	case uint16(opcodes.StackOpCode):
		return []uint16{
			uint16(opcodes.ClearOpCode) | argTwo<<6, // Clear(y)
			uint16(opcodes.HoldingOpCode) | argOne,  // Holding(x)
		}
	case uint16(opcodes.UnstackOpCode):
		return []uint16{
			uint16(opcodes.OnOpCode) | argOne | argTwo, // On(x,y)
			uint16(opcodes.HoldingOpCode),              // Holding(nil)
			uint16(opcodes.ClearOpCode) | argOne,       // Clear(x)
		}
	case uint16(opcodes.PickupOpCode):
		return []uint16{
			uint16(opcodes.OnOpCode) | argOne | uint16(argtwo.Floor), // On(x, floor)
			uint16(opcodes.HoldingOpCode),                            // Holding(nil)
		}
	default:
		return nil
	}
}

func (state *State) getApplications(cmd uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(cmd)

	switch opCode {
	case uint16(opcodes.StackOpCode):
		return []uint16{
			uint16(opcodes.HoldingOpCode),              // Holding(nil)
			uint16(opcodes.OnOpCode) | argOne | argTwo, // On(x,y)
			uint16(opcodes.ClearOpCode) | argOne,       // Clear(x)
		}
	case uint16(opcodes.UnstackOpCode):
		return []uint16{
			uint16(opcodes.HoldingOpCode) | argOne,  // Holding(x)
			uint16(opcodes.ClearOpCode) | argTwo<<6, // Clear(y)
		}
	case uint16(opcodes.PickupOpCode):
		return []uint16{
			uint16(opcodes.HoldingOpCode) | argOne, // Holding(x)
		}
	default:
		return nil
	}
}

func (state *State) Execute(cmd uint16) bool {
	// Delete the delete items first
	// fmt.Println("\n\n\n\n\n\nBefore: ")
	// for _, del := range state.predicates {
	// 	opcodes.PrintPredicate(del)
	// }

	for _, predToDelete := range state.getDeletes(cmd) {
		for i, pred := range state.predicates {
			if pred == predToDelete {
				state.predicates = append(state.predicates[:i], state.predicates[i+1:]...)
				break
			}
		}
	}

	// Apply new predicates
	state.predicates = append(state.predicates, state.getApplications(cmd)...)

	// fmt.Println("\n\n\n\n\n\nAfter: ")
	// for _, del := range state.predicates {
	// 	opcodes.PrintPredicate(del)
	// }

	return true
}

func (state *State) FindApplications(pred uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(pred)

	switch opCode {
	case uint16(opcodes.OnOpCode):
		return []uint16{
			uint16(opcodes.StackOpCode) | argOne | argTwo, // Stack(x,y) | Putdown(x)
		}
	case uint16(opcodes.HoldingOpCode):
		if argOne == uint16(argone.Blank) {
			return []uint16{
				uint16(opcodes.StackOpCode) | state.findWhatIsBeingHeld() | uint16(argtwo.Floor), // Putdown(x)
			}
		} else {
			base := state.findWhatBlockIsOn(argOne)
			if base == uint16(argtwo.Blank) {
				return []uint16{
					uint16(opcodes.PickupOpCode) | argOne, // Pickup(x)
				}
			} else {
				return []uint16{
					uint16(opcodes.UnstackOpCode) | argOne | base, // Unstack(x, ?)
				}
			}
		}
	case uint16(opcodes.ClearOpCode):
		top := state.findWhatIsOnBlock(argOne)
		return []uint16{
			uint16(opcodes.UnstackOpCode) | top | argOne>>6, // Unstack(?, x)
		}
	default:
		return nil
	}
}
