package compiler

import (
	"farcical/ast"
	"farcical/code"
	"farcical/object"
	"fmt"
)

// The compiler. Maintains the instructions being generated and the constants pool.
type Compiler struct {
	instructions code.Instructions // The bytecode instructions being generated.
	constants    []object.Object   // Constants pool for storing objects.
}

// Creates a new instance of the Farcical compiler.
func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

// Generates bytecode instructions from the provided AST node.
// It returns an error if the compilation process encounters any issues.
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
	case *ast.InfixExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err2 := c.Compile(node.Right)
		if err2 != nil {
			return err
		}
	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))
	}
	return nil
}

// Add a constant to the *Compiler's constants slice
// Returns its index in the slice (NOT the constant itself)
func (c *Compiler) addConstant(obj object.Object) int {
	fmt.Printf("Compiler: Add value to constants: %+v\n", obj.Inspect())
	c.constants = append(c.constants, obj) // add the new object to constant pool/slice

	fmt.Printf("Compiler: At pos/index: %d\n", len(c.constants)-1)
	return len(c.constants) - 1 // return its index in the constants slice
}

// Makes a bytecode instruction, adds it to the *Compiler's instructions slice,
// and returns the starting position of the just-emitted instruction (allowing us to go back and modify it later?)
func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	// fmt.Println("emit:")
	// fmt.Println(ins)
	pos := c.addInstruction(ins)
	return pos
}

// Adds a new instruction to the *Compiler's instructions slice
// Returns the *position* of the instruction in the slice
func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	// fmt.Println("all intructions:")
	// fmt.Println(c.instructions.String())
	return posNewInstruction
}

// Returns the compiled bytecode produced by the compiler.
// It packages the generated bytecode instructions and constants into a Bytecode structure.
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// Represents the compiled bytecode produced by the compiler.
type Bytecode struct {
	Instructions code.Instructions // The generated bytecode instructions.
	Constants    []object.Object   // Constants used in the bytecode.
}
