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

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	Opconstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]

	fmt.Printf("definition: %v\n", def)

	if !ok {
		return []byte{}
	}

	instructionLen := 1

	fmt.Printf("opcode: %d\n", op)
	fmt.Printf("operands: %v\n", operands)

	for _, w := range def.OperandWidths {
		instructionLen += w
		fmt.Printf("instruction len (inc opcode): %d, width: %v\n", instructionLen, w)
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		fmt.Printf("operand at pos %d: %v\n", i, o)
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}
