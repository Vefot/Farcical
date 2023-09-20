package vm

import (
	"farcical/ast"
	"farcical/compiler"
	"farcical/lexer"
	"farcical/object"
	"farcical/parser"
	"fmt"
	"testing"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 2}, // fix this to handle addition later
	}

	runVmTests(t, tests)
}

// Lex and parse the input, pass it to the compiler (& check for errors), then hands the bytecode instructions
// generated by the compiler to the New function (the new VM instance)
func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err2 := vm.Run()
		if err2 != nil {
			t.Fatalf("vm error: %s", err)
		}

		stackElem := vm.StackTop() // get the object that's left sitting at the top of the VM's stack - check it matches what we expect

		testExpectedObject(t, tt.expected, stackElem)
	}
}

func testExpectedObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
		fmt.Printf("VM tests: expected %d at top of stack - got %s\n\n", expected, actual.Inspect())
	}
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
