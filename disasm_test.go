package zdisasm_test

import (
	"testing"

	"github.com/philippeleite/zdisasm"
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
		// Bug #1: CGIT (RIE-a) must show R1,I2,M3 (not R1,R3,RI2)
		{name: "CGIT", input: "EC50FFFB8070", want: "CGIT  R5,-5,8"},
		// Bug #1: CLGIT (RIE-a unsigned) must show R1,I2,M3
		{name: "CLGIT", input: "EC50000A8071", want: "CLGIT R5,10,8"},
		// Bug #1: CIT (RIE-a signed)
		{name: "CIT", input: "EC30FFEC4072", want: "CIT   R3,-20,4"},
		// Bug #1: CLFIT (RIE-a unsigned)
		{name: "CLFIT", input: "EC3000144073", want: "CLFIT R3,20,4"},
		// Bug #2: formatMII sign extension for BPRP (12-bit and 24-bit)
		{name: "BPRP-neg", input: "C57FFE800020", want: "BPRP  7,-2,-8388576"},
		{name: "BPRP-pos", input: "C50002000010", want: "BPRP  0,2,16"},
		// Bug #3: formatRRS displacement > 255
		{name: "CRB", input: "ECF12FFF60F6", want: "CRB   R15,R1,6,4095(R2)"},
		// Bug #4: formatVRS (VRS-d) reads V1 from bits 32-35
		{name: "VLRLR", input: "E60310002037", want: "VLRLR V2,R3,0(R1)"},
		{name: "VLRLR-rxb", input: "E60310002137", want: "VLRLR V18,R3,0(R1)"},
		// RXB extends vector register fields to V16-V31
		{name: "VLR-rxb0", input: "E71200000856", want: "VLR   V17,V2,0,0"},
		{name: "VLR-rxb1", input: "E71200000456", want: "VLR   V1,V18,0,0"},
		{name: "VLR-rxb01", input: "E7FF00000C56", want: "VLR   V31,V31,0,0"},
		// Bug #5: formatSMI with large RI2 for BPP
		{name: "BPP", input: "C7F01000FFEC", want: "BPP   15,-20,0(R1)"},
		// Error: odd number of hex digits
		{name: "odd-length", input: "EB1F200000080", wantErr: true},
		// Bug #7: PBL opcode 77 is invalid
		{name: "PBL-invalid", input: "77001000", wantErr: true},
		// Bug #8: LGFI shows signed immediate
		{name: "LGFI", input: "C011FFFFFFFE", want: "LGFI  R1,-2"},
		// Bug #8: AFI shows signed immediate
		{name: "AFI", input: "C299FFFFFFFE", want: "AFI   R9,-2"},
		// Bug #6: ICMH with RSY2 format (mask, not register)
		{name: "ICMH", input: "EB1F20000080", want: "ICMH  R1,15,0(R2)"},
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
