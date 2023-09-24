package vm

import (
	"farcical/code"
	"farcical/compiler"
	"farcical/object"
	"fmt"
)

const StackSize = 2048

// Farcical Virtual Machine.
// Holds the constants and instructions generated by the compiler.
// Has a stack, with StackSize number of elements.
// Instead of modifying the stack itself, decrement or increment the sp (stackpointer)
// to grow or shrink the stack.
// sp points to the next free slot, so to add a slot it goes to stack[sp]
type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object
	sp    int // stackpointer. always points to the next value - top of the stack is stack[sp-1]. sp itself points to the next FREE slot
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}

// Returns the object on the top of the stack
func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

// The main loop of the FVM, running the fetch-decode-execute cycle
// ip - instruction pointer, fetches the current instruction
// turn the byte into an opcode - NOT using code.Lookup as that is too slow (it
// costs time to move the byte around)
func (vm *VM) Run() error {
	// FETCH
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		// DECODE
		switch op {
		case code.OpConstant:
			// decode the bytes AFTER the opcode (the operands) - not using code.ReadOperands for the same reasons as Lookup
			// gues since its a uint16 (2 bytes) it knows to only read 2?
			constIndex := code.ReadUint16(vm.instructions[ip+1:]) // still not sure what's going on here - we're reading an index? so where is the value?
			ip += 2                                               // careful to increment ip by correct amount - next iteration must be pointing at an opcode not an operand

			// Uses constIndex to get to the constant in vm.constants and push it to the stack
			// I think what happens is: if we get an OpConstant opcode, read the instruction which gives us the index in the bytecod
			// constants slice of where the constant's value is stored. Then push it to the vm constants slice.
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		case code.OpAdd:
			// EXECUTE
			// take the top two elements from the stack, pop them off it, extract their values, add them, push the result to the stack
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			result := leftValue + rightValue
			vm.push(&object.Integer{Value: result})
		case code.OpPop:
			vm.pop()
		}
	}
	return nil
}

// Returns the element that was last popped off the stack (by code.OpPop).
// Since sp points to the next *free* slot, and we only pop elements by decrementing the stack pointer,
// (without setting them to nil), this is where we find the last elements that were previously
// on top of the stack
func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

// Pushes the object to the top of the stack
func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow") // woop woop
	}

	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

// Pop the top element of the stack off.
// Take the element from the top of the stack and put it to the side, then decrement sp.
// This allows the location of the element that was just popped off to be overwritten
// eventually.
func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}
