package main

import "testing"

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
		// Format RXY (20-bit displacement)
		{name: "LG", input: "E31F10008004", want: "LG    R1,8(R15,R1)"},
		// Error: non-hex input
		{name: "non-hex", input: "ZZZZ", wantErr: true},
		// Error: wrong length for opcode class
		{name: "bad-length", input: "18", wantErr: true},
		// Error: opcode not in table
		{name: "unknown-opcode", input: "FFFF", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := disasm(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Errorf("disasm(%q) expected error, got %q", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Errorf("disasm(%q) unexpected error: %v", tc.input, err)
				return
			}
			if got != tc.want {
				t.Errorf("disasm(%q)\n  got:  %q\n  want: %q", tc.input, got, tc.want)
			}
		})
	}
}
