package compiler

import (
	"farcical/ast"
	"farcical/code"
	"farcical/lexer"
	"farcical/object"
	"farcical/parser"
	"fmt"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testinstructions failed: %s", err)
		}

		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}

func testInstructions(expected []code.Instructions, actual code.Instructions) error {
	concatted := concatInstructions(expected)

	if len(actual) != len(concatted) {
		return fmt.Errorf("wrong instructions length.\nwant=%q\ngot=%q", concatted, actual)
	}

	for i, ins := range concatted {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot=%q", i, concatted, actual)
		}
	}
	return nil
}

// We need concatInstructions because the expectedInstructions is a slice of slices of bytes, not just a slice
// because we used code.Make to generate the expectedInstructions - which produces a []byte
// Basically this flattens the slice of slices into a slice
func concatInstructions(s []code.Instructions) code.Instructions {
	out := code.Instructions{}

	fmt.Println("Instruction:")
	fmt.Println(s)
	// the String method will implicitly be called
	// since string method includes ReadOperands method, which
	// returns number of bytes read, we should see the numbers incrementing like so:
	// 0000, 0003, 0006, etc, depending on num bytes read...

	for _, ins := range s {
		fmt.Println("Instruction element:")
		fmt.Printf("%q\n", ins)
		out = append(out, ins...)
	}

	fmt.Printf("Outputting from concat instructions:\n%s", out)
	return out
}

func testConstants(t *testing.T, expected []interface{}, actual []object.Object) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d", len(actual), len(expected))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed: %s", i, err)
			}
		}
	}
	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not integer, got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value, got=%d, want=%d", result.Value, expected)
	}

	return nil
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
