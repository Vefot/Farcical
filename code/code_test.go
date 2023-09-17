package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		// check the order of bytes returned - picked 65534 so we could test for big endianism
		// first byte is the opcode, then the encoded operand (so 255 and 254 make 65534 - we can test big endianism since they're different numbers)
		{Opconstant, []int{65534}, []byte{byte(Opconstant), 255, 254}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d", len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d got=%d", i, b, instruction[i])
			}
		}
	}
}
