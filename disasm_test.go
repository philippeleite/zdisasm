package zdisasm_test

import (
	"testing"

	"zdisasm"
)

func TestDisasm(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		// Format RR1
		{name: "AR", input: "1A34", want: "AR    R3,R4"},
		// Format RX1
		{name: "LA", input: "4110F010", want: "LA    R1,16(R0,R15)"},
		// Format RI4 (signed 16-bit immediate)
		{name: "BRC", input: "A7F4FFEC", want: "BRC   15,-20"},
		// Format RIL3 (32-bit immediate via unsigned-to-signed cast)
		{name: "BRASL", input: "C0E5FFFFFFEC", want: "BRASL R14,-20"},
		// Format RXY (20-bit displacement): E3=op1, 1=R1, F=X2(R15), 1=B2(R1), 008=DL2, 00=DH2, 04=op2
		{name: "LG", input: "E31F10080004", want: "LG    R1,8(R15,R1)"},
		// Error: non-hex input
		{name: "non-hex", input: "ZZZZ", wantErr: true},
		// Error: wrong length for opcode class
		{name: "bad-length", input: "18", wantErr: true},
		// Error: opcode not in table (valid length 6-byte F-prefix, unknown opcode)
		{name: "unknown-opcode", input: "FFFFFFFFFFFF", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := zdisasm.Disasm(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Errorf("Disasm(%q) expected error, got %q", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Errorf("Disasm(%q) unexpected error: %v", tc.input, err)
				return
			}
			if got != tc.want {
				t.Errorf("Disasm(%q)\n  got:  %q\n  want: %q", tc.input, got, tc.want)
			}
		})
	}
}
