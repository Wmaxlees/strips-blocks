package state

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
	"github.com/Wmaxlees/strips-blocks/opcodes"
)

type State []uint16

// TODO: Make this more efficient
func (state *State) Print() {
	blockStacks := make([][]uint16, 0)
	var holding uint16 = 0
	// Get all the blocks on the table and find if there is a block being held by arm
	for _, pred := range *state {
		if (pred&opcodes.OpCodeMask) == opcodes.OnOpCode && (pred&argtwo.ArgTwoMask) == argtwo.Floor {
			blockStacks = append(blockStacks, []uint16{(pred & argone.ArgOneMask) >> 6})
		}

		if (pred & opcodes.OpCodeMask) == opcodes.HoldingOpCode {
			holding = pred & argone.ArgOneMask >> 6
		}
	}

	// Stack all the blocks on top of these
	newBlocks := true
	for newBlocks {
		newBlocks = false
		for i, stack := range blockStacks {
			top := stack[len(stack)-1]
			for _, pred := range *state {
				if pred&argtwo.ArgTwoMask == top {
					blockStacks[i] = append(stack, (pred&argone.ArgOneMask)>>6)
					newBlocks = true
				}
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

	// Print the holding indication
	fmt.Println("___________")
	if holding != 0 {
		fmt.Printf(" %c\n%c%c%c\n[%c]\n\n", 0x2502, 0x250C, 0x2534, 0x2510, holding+64)
	} else {
		fmt.Printf(" %c\n%c%c%c\n\n", 0x2502, 0x250C, 0x2534, 0x2510)
	}

	// Print out the stacks from the top down
	for i := stackHeight; i >= -1; i-- {

		// Print blocks
		for j := 0; j < len(blockStacks); j++ {
			// Print the floor
			if i == -1 {
				fmt.Print("----")
				continue
			}

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
	InitialStateSimpleProblem         StateDefinition = iota
	InitialStateSussmanAnomalyProblem StateDefinition = iota
	InitialStateExtraProblem          StateDefinition = iota
	GoalStateSimpleProblem            StateDefinition = iota
	GoalStateSussmanAnomalyProblem    StateDefinition = iota
	GoalStateExtraProblem             StateDefinition = iota
)

func (state *State) AddPredicate(opCode uint16, argOne uint16, argTwo uint16) {
	*state = append(*state, opCode|argOne|argTwo)
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

	case InitialStateExtraProblem:
		// First stack
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockI, argtwo.Blank)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockI, argtwo.BlockA)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockA, argtwo.BlockC)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockC, argtwo.BlockF)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockF, argtwo.Floor)

		// Second stack
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockJ, argtwo.Blank)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockJ, argtwo.BlockB)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockB, argtwo.Floor)

		// Third stack
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockH, argtwo.Blank)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockH, argtwo.BlockD)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockD, argtwo.Floor)

		// Forth stack
		state.AddPredicate(opcodes.ClearOpCode, argone.BlockL, argtwo.Blank)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockL, argtwo.BlockG)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockG, argtwo.BlockE)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockE, argtwo.BlockK)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockK, argtwo.Floor)

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

	case GoalStateExtraProblem:
		// First stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockA, argtwo.BlockB)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockB, argtwo.BlockC)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockC, argtwo.BlockD)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockD, argtwo.Floor)

		// Second stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockE, argtwo.BlockF)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockF, argtwo.BlockG)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockG, argtwo.BlockH)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockH, argtwo.Floor)

		// Third stack
		state.AddPredicate(opcodes.OnOpCode, argone.BlockI, argtwo.BlockJ)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockJ, argtwo.BlockK)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockK, argtwo.BlockL)
		state.AddPredicate(opcodes.OnOpCode, argone.BlockL, argtwo.Floor)
	default:
	}
}

func (state *State) CheckPredicate(predicate uint16) bool {
	for _, existingPred := range *state {
		if uint16(existingPred) == predicate {
			return true
		}
	}

	return false
}

func (state *State) findWhatBlockIsOn(block uint16) uint16 {
	for _, pred := range *state {
		opCode, argOne, argTwo := opcodes.GetComponents(pred)

		if opCode == opcodes.OnOpCode && argOne == block {
			return argTwo
		}
	}

	return argtwo.Blank
}

func (state *State) findWhatIsOnBlock(block uint16) uint16 {
	block = block >> 6

	for _, pred := range *state {
		opCode, argOne, argTwo := opcodes.GetComponents(pred)

		if opCode == opcodes.OnOpCode && argTwo == block {
			return argOne
		}
	}

	return argtwo.Blank
}

func (state *State) findWhatIsBeingHeld() uint16 {
	for _, pred := range *state {
		opCode, argOne, _ := opcodes.GetComponents(pred)

		if opCode == opcodes.HoldingOpCode {
			return argOne
		}
	}

	return 0
}

func (state *State) GetPreconditions(cmd uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(cmd)
	switch opCode {
	case opcodes.StackOpCode: // Stack(x,y)
		if argTwo != argtwo.Floor {
			return []uint16{
				opcodes.HoldingOpCode | argOne | argtwo.Blank, // Holding(x)
				opcodes.ClearOpCode | argTwo<<6,               // Clear(y)
			}
		} else {
			return []uint16{
				opcodes.HoldingOpCode | argOne | argtwo.Blank, // Holding(x)
			}
		}
	case opcodes.UnstackOpCode: // Unstack(x,y)
		return []uint16{
			opcodes.OnOpCode | argOne | argTwo,                  // On(x,y)
			opcodes.ClearOpCode | argOne,                        // Clear(x)
			opcodes.HoldingOpCode | argone.Blank | argtwo.Blank, // Holding(nil)
		}
	case opcodes.PickupOpCode: // Pickup(x)
		return []uint16{
			opcodes.ClearOpCode | argOne,                        // Clear(x)
			opcodes.OnOpCode | argOne | argtwo.Floor,            // On(x, Floor)
			opcodes.HoldingOpCode | argone.Blank | argtwo.Blank, // Holding(nil)
		}
	default:
		return nil
	}
}

func (state *State) getDeletes(cmd uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(cmd)

	switch opCode {
	case opcodes.StackOpCode:
		return []uint16{
			opcodes.ClearOpCode | argTwo<<6, // Clear(y)
			opcodes.HoldingOpCode | argOne,  // Holding(x)
		}
	case opcodes.UnstackOpCode:
		return []uint16{
			opcodes.OnOpCode | argOne | argTwo, // On(x,y)
			opcodes.HoldingOpCode,              // Holding(nil)
			opcodes.ClearOpCode | argOne,       // Clear(x)
		}
	case opcodes.PickupOpCode:
		return []uint16{
			opcodes.OnOpCode | argOne | argtwo.Floor, // On(x, floor)
			opcodes.HoldingOpCode,                    // Holding(nil)
		}
	default:
		return nil
	}
}

func (state *State) getApplications(cmd uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(cmd)

	switch opCode {
	case opcodes.StackOpCode:
		return []uint16{
			opcodes.HoldingOpCode,              // Holding(nil)
			opcodes.OnOpCode | argOne | argTwo, // On(x,y)
			opcodes.ClearOpCode | argOne,       // Clear(x)
		}
	case opcodes.UnstackOpCode:
		return []uint16{
			opcodes.HoldingOpCode | argOne,  // Holding(x)
			opcodes.ClearOpCode | argTwo<<6, // Clear(y)
		}
	case opcodes.PickupOpCode:
		return []uint16{
			opcodes.HoldingOpCode | argOne, // Holding(x)
		}
	default:
		return nil
	}
}

func (state *State) Execute(cmd uint16) bool {
	for _, predToDelete := range state.getDeletes(cmd) {
		for i, pred := range *state {
			if pred == predToDelete {
				*state = append((*state)[:i], (*state)[i+1:]...)
				break
			}
		}
	}

	// Apply new predicates
	*state = append(*state, state.getApplications(cmd)...)

	return true
}

func (state *State) FindApplications(pred uint16) []uint16 {
	opCode, argOne, argTwo := opcodes.GetComponents(pred)

	switch opCode {
	case opcodes.OnOpCode:
		return []uint16{
			opcodes.StackOpCode | argOne | argTwo, // Stack(x,y) | Putdown(x)
		}
	case opcodes.HoldingOpCode:
		if argOne == argone.Blank {
			return []uint16{
				opcodes.StackOpCode | state.findWhatIsBeingHeld() | argtwo.Floor, // Putdown(x)
			}
		} else {
			base := state.findWhatBlockIsOn(argOne)
			if base == argtwo.Blank {
				return []uint16{
					opcodes.PickupOpCode | argOne, // Pickup(x)
				}
			} else {
				return []uint16{
					opcodes.UnstackOpCode | argOne | base, // Unstack(x, ?)
				}
			}
		}
	case opcodes.ClearOpCode:
		top := state.findWhatIsOnBlock(argOne)
		return []uint16{
			opcodes.UnstackOpCode | top | argOne>>6, // Unstack(?, x)
		}
	default:
		return nil
	}
}
