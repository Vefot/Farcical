package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	Opconstant Opcode = iota
)

// Definition is a data structure representing the characteristics of an opcode,
// including its name and operand widths.
type Definition struct {
	Name          string
	OperandWidths []int
}

// The Key is the opcode, and the Value is the Definition.
// They key/opcode is a number represented as a byte, defined as a constant.
var definitions = map[Opcode]*Definition{
	Opconstant: {"OpConstant", []int{2}},
}

// Retrieves the definition of an opcode based on its byte representation.
// If the opcode is defined, it returns the associated Definition;
// otherwise, it returns an error indicating that the opcode is undefined.
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

// Make creates a binary instruction for the virtual machine based on the given opcode and operands.
// It looks up the opcode's definition to determine the instruction format and operand widths, then constructs
// the instruction by encoding the opcode and operands accordingly.
//
// Parameters:
//   - op: The opcode representing the operation to be executed.
//   - operands: A variadic parameter allowing a variable number of integer operands to be passed.
//
// Returns:
//   - []byte: The binary instruction as a byte slice.
func Make(op Opcode, operands ...int) []byte {
	// Attempt to find the opcode definition
	def, ok := definitions[op]

	// If the opcode is not defined, return an empty instruction
	if !ok {
		return []byte{}
	}

	// Initialize the instruction length with 1 (for the opcode itself)
	instructionLen := 1

	// Calculate the total instruction length by adding the widths of all operands
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	// Create a byte slice to store the instruction
	instruction := make([]byte, instructionLen)

	// Set the first byte of the instruction to the opcode
	instruction[0] = byte(op)

	// Initialize the offset to 1, as the opcode is at position 0
	offset := 1

	// Iterate through operands and encode them into the instruction
	for i, o := range operands {
		// Determine the width of the current operand
		width := def.OperandWidths[i]

		// Encode the operand into the instruction based on its width
		switch width {
		case 2:
			// For 2-byte operands, use BigEndian encoding
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		// Update the offset to the next position for the next operand
		offset += width
	}

	// Return the constructed instruction
	return instruction
}
