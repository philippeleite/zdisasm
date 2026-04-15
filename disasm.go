package zdisasm

import (
	"fmt"
	"strconv"
)

// ── helpers ──────────────────────────────────────────────────────────────────

func isHex(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// signExt12 sign-extends a value from 12 bits to int64.
func signExt12(v int64) int64 {
	if (v>>11)&1 == 1 {
		return v | (-1 << 12)
	}
	return v
}

// signExt16 sign-extends a value from 16 bits to int64.
func signExt16(v int64) int64 {
	if (v>>15)&1 == 1 {
		return v | (-1 << 16)
	}
	return v
}

// signExt20 sign-extends a value from 20 bits to int64.
func signExt20(v int64) int64 {
	if (v>>19)&1 == 1 {
		return v | (-1 << 20)
	}
	return v
}

// signExt24 sign-extends a value from 24 bits to int64.
func signExt24(v int64) int64 {
	if (v>>23)&1 == 1 {
		return v | (-1 << 24)
	}
	return v
}

// signExt32 replicates C's unsigned-to-signed int32 cast.
func signExt32(v uint64) int64 {
	return int64(int32(uint32(v)))
}

// disp20 decodes a 20-bit RSY/RXY/SIY displacement from the hex string.
// High byte is at s[8:10], low 12 bits at s[5:8].
func disp20(s string) int64 {
	raw, _ := strconv.ParseInt(s[8:10]+s[5:8], 16, 64)
	return signExt20(raw)
}

// ── format functions ─────────────────────────────────────────────────────────

func formatE(_ string) string { return "" }

func formatI(s string) string {
	i, _ := strconv.ParseUint(s[2:4], 16, 8)
	return fmt.Sprintf("%d", i)
}

func formatIE(s string) string {
	i1, _ := strconv.ParseUint(s[6:7], 16, 8)
	i2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("%d,%d", i1, i2)
}

func formatMII(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i1, _ := strconv.ParseInt(s[3:6], 16, 16)
	i2, _ := strconv.ParseInt(s[6:12], 16, 32)
	return fmt.Sprintf("%d,%d,%d", m1, signExt12(i1), signExt24(i2))
}

func formatRI1(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	return fmt.Sprintf("R%d,%d", r1, i2)
}

func formatRI2(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	return fmt.Sprintf("%d,%d", m1, i2)
}

func formatRI3(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	return fmt.Sprintf("R%d,%d", r1, signExt16(raw))
}

func formatRI4(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	return fmt.Sprintf("%d,%d", m1, signExt16(raw))
}

func formatRIL1(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i2, _ := strconv.ParseUint(s[4:12], 16, 32)
	return fmt.Sprintf("R%d,%d", r1, i2)
}

func formatRIL2(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i2, _ := strconv.ParseUint(s[4:12], 16, 32)
	return fmt.Sprintf("%d,%d", m1, i2)
}

func formatRIL3(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	raw, _ := strconv.ParseUint(s[4:12], 16, 64)
	return fmt.Sprintf("R%d,%d", r1, signExt32(raw))
}

func formatRIL4(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	raw, _ := strconv.ParseUint(s[4:12], 16, 32)
	return fmt.Sprintf("%d,%d", m1, signExt32(raw))
}

func formatRIE(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	return fmt.Sprintf("R%d,R%d,%d", r1, r3, signExt16(raw))
}

func formatRIEa(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,%d,%d", r1, signExt16(raw), m3)
}

func formatRIEaU(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,%d,%d", r1, i2, m3)
}

func formatRIEb(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d,%d", r1, r2, m3, signExt16(raw))
}

func formatRIEc(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	m3, _ := strconv.ParseUint(s[3:4], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	i2, _ := strconv.ParseUint(s[8:10], 16, 16)
	return fmt.Sprintf("R%d,%d,%d,%d", r1, i2, m3, signExt16(raw))
}

func formatRIEd(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	return fmt.Sprintf("R%d,R%d,%d", r1, r3, i2)
}

func formatRIEf(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	i3, _ := strconv.ParseUint(s[4:6], 16, 16)
	i4, _ := strconv.ParseUint(s[6:8], 16, 16)
	i5, _ := strconv.ParseUint(s[8:10], 16, 16)
	return fmt.Sprintf("R%d,R%d,%d,%d,%d", r1, r2, i3, i4, i5)
}

func formatRIEg(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	m3, _ := strconv.ParseUint(s[3:4], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	return fmt.Sprintf("R%d,%d,%d", r1, signExt16(raw), m3)
}

func formatRIS(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	m3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b4, _ := strconv.ParseUint(s[4:5], 16, 8)
	d4, _ := strconv.ParseUint(s[5:8], 16, 16)
	i2, _ := strconv.ParseUint(s[8:10], 16, 16)
	return fmt.Sprintf("R%d,%d,%d,%d(R%d)", r1, i2, m3, d4, b4)
}

func formatRR1(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	return fmt.Sprintf("R%d,R%d", r1, r2)
}

func formatRR2(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	return fmt.Sprintf("%d,R%d", m1, r2)
}

func formatRR3(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	return fmt.Sprintf("R%d", r1)
}

func formatRRE1(s string) string {
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d", r1, r2)
}

func formatRRE2(s string) string {
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	return fmt.Sprintf("R%d", r1)
}

func formatRRF1(s string) string {
	r1, _ := strconv.ParseUint(s[4:5], 16, 8)
	r3, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,R%d", r3, r1, r2)
}

func formatRRF2(s string) string {
	m3, _ := strconv.ParseUint(s[4:5], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,%d,R%d", r1, m3, r2)
}

func formatRRFa(s string) string {
	r3, _ := strconv.ParseUint(s[4:5], 16, 8)
	m4, _ := strconv.ParseUint(s[5:6], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,R%d,%d", r1, r2, r3, m4)
}

func formatRRF3(s string) string {
	r3, _ := strconv.ParseUint(s[4:5], 16, 8)
	m4, _ := strconv.ParseUint(s[5:6], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,R%d,%d", r1, r3, r2, m4)
}

func formatRRF4(s string) string {
	m3, _ := strconv.ParseUint(s[4:5], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d", r1, r2, m3)
}

func formatRRF5(s string) string {
	r1, _ := strconv.ParseUint(s[4:5], 16, 8)
	r3, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,R%d", r1, r3, r2)
}

func formatRRF6(s string) string {
	m3, _ := strconv.ParseUint(s[4:5], 16, 8)
	m4, _ := strconv.ParseUint(s[5:6], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,%d,R%d,%d", r1, m3, r2, m4)
}

func formatRRF7(s string) string {
	r3, _ := strconv.ParseUint(s[4:5], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,R%d", r1, r2, r3)
}

func formatRRF8(s string) string {
	r3, _ := strconv.ParseUint(s[4:5], 16, 8)
	m4, _ := strconv.ParseUint(s[5:6], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,R%d,%d", r1, r2, r3, m4)
}

func formatRRF9(s string) string {
	m4, _ := strconv.ParseUint(s[5:6], 16, 8)
	r1, _ := strconv.ParseUint(s[6:7], 16, 8)
	r2, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d", r1, r2, m4)
}

func formatRRS(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b4, _ := strconv.ParseUint(s[4:5], 16, 8)
	d4, _ := strconv.ParseUint(s[5:8], 16, 16)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d,%d(R%d)", r1, r2, m3, d4, b4)
}

func formatRS1(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("R%d,R%d,%d(R%d)", r1, r3, d2, b2)
}

func formatRS2(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	m3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("R%d,%d,%d(R%d)", r1, m3, d2, b2)
}

func formatRSI(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	raw, _ := strconv.ParseInt(s[4:8], 16, 64)
	return fmt.Sprintf("R%d,R%d,%d", r1, r3, signExt16(raw))
}

func formatRSL(s string) string {
	l1, _ := strconv.ParseUint(s[2:3], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("%d(%d,R%d)", d1, l1+1, b1)
}

func formatRSLb(s string) string {
	l2, _ := strconv.ParseUint(s[2:4], 16, 16)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	r1, _ := strconv.ParseUint(s[8:9], 16, 8)
	m3, _ := strconv.ParseUint(s[9:10], 16, 8)
	return fmt.Sprintf("R%d,%d(%d,R%d),%d", r1, d2, l2+1, b2, m3)
}

func formatRSY1(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d(R%d)", r1, r3, disp20(s), b2)
}

func formatRSY2(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	m3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("R%d,%d,%d(R%d)", r1, m3, disp20(s), b2)
}

func formatRX1(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("R%d,%d(R%d,R%d)", r1, d2, x2, b2)
}

func formatRX2(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("%d,%d(R%d,R%d)", m1, d2, x2, b2)
}

func formatRXE(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("R%d,%d(R%d,R%d)", r1, d2, x2, b2)
}

func formatRXE2(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,%d(R%d,R%d),%d", r1, d2, x2, b2, m3)
}

func formatRXF(s string) string {
	r3, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	r1, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d(R%d,R%d)", r1, r3, d2, x2, b2)
}

func formatRXG(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("%d,%d(R%d,R%d)", m1, disp20(s), x2, b2)
}

func formatRXY(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("R%d,%d(R%d,R%d)", r1, disp20(s), x2, b2)
}

func formatS(s string) string {
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("%d(R%d)", d2, b2)
}

func formatSI(s string) string {
	i2, _ := strconv.ParseUint(s[2:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	return fmt.Sprintf("%d(R%d),%d", d1, b1, i2)
}

func formatSIL(s string) string {
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	i2, _ := strconv.ParseUint(s[8:12], 16, 32)
	return fmt.Sprintf("%d(R%d),%d", d1, b1, i2)
}

func formatSIY(s string) string {
	i2, _ := strconv.ParseUint(s[2:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("%d(R%d),%d", disp20(s), b1, i2)
}

func formatSMI(s string) string {
	m1, _ := strconv.ParseUint(s[2:3], 16, 8)
	b3, _ := strconv.ParseUint(s[4:5], 16, 8)
	d3, _ := strconv.ParseUint(s[5:8], 16, 16)
	i2, _ := strconv.ParseInt(s[8:12], 16, 64)
	return fmt.Sprintf("%d,%d,%d(R%d)", m1, signExt16(i2), d3, b3)
}

func formatSS1(s string) string {
	l, _ := strconv.ParseUint(s[2:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(%d,R%d),%d(R%d)", d1, l+1, b1, d2, b2)
}

func formatSS2(s string) string {
	l1, _ := strconv.ParseUint(s[2:3], 16, 8)
	l2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(%d,R%d),%d(%d,R%d)", d1, l1+1, b1, d2, l2+1, b2)
}

func formatSS3(s string) string {
	l1, _ := strconv.ParseUint(s[2:3], 16, 8)
	i3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(%d,R%d),%d(R%d),%d", d1, l1+1, b1, d2, b2, i3)
}

func formatSS4(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(R%d,R%d),%d(R%d),R%d", d1, r1, b1, d2, b2, r3)
}

func formatSS5(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	b4, _ := strconv.ParseUint(s[8:9], 16, 8)
	d4, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("R%d,%d(R%d),R%d,%d(R%d)", r1, d2, b2, r3, d4, b4)
}

func formatSS6(s string) string {
	l, _ := strconv.ParseUint(s[2:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(R%d),%d(%d,R%d)", d1, b1, d2, l+1, b2)
}

func formatSSE(s string) string {
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(R%d),%d(R%d)", d1, b1, d2, b2)
}

func formatSSF(s string) string {
	r3, _ := strconv.ParseUint(s[2:3], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(R%d),%d(R%d),R%d", d1, b1, d2, b2, r3)
}

func formatSSG(s string) string {
	r3, _ := strconv.ParseUint(s[2:3], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("R%d,%d(R%d),%d(R%d)", r3, d1, b1, d2, b2)
}

func formatVRIa(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,%d,%d", v1, i2, m3)
}

func formatVRIb(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	i2, _ := strconv.ParseUint(s[4:6], 16, 16)
	i3, _ := strconv.ParseUint(s[6:8], 16, 16)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,%d,%d,%d", v1, i2, i3, m4)
}

func formatVRIc(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,V%d,%d,%d", v1, v3, i2, m4)
}

func formatVRId(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	i4, _ := strconv.ParseUint(s[6:8], 16, 16)
	m5, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,V%d,V%d,%d,%d", v1, v2, v3, i4, m5)
}

func formatVRIe(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	i3, _ := strconv.ParseUint(s[4:7], 16, 16)
	m5, _ := strconv.ParseUint(s[7:8], 16, 8)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,V%d,%d,%d,%d", v1, v2, i3, m4, m5)
}

func formatVRIf(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	m5, _ := strconv.ParseUint(s[6:7], 16, 8)
	i4, _ := strconv.ParseUint(s[7:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,V%d,%d,%d", v1, v2, v3, i4, m5)
}

func formatVRIg(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	i4, _ := strconv.ParseUint(s[4:6], 16, 16)
	m5, _ := strconv.ParseUint(s[6:7], 16, 8)
	i3, _ := strconv.ParseUint(s[7:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,%d,%d,%d", v1, v2, i3, i4, m5)
}

func formatVRIh(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	i2, _ := strconv.ParseUint(s[4:8], 16, 16)
	i3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,%d,%d", v1, i2, i3)
}

func formatVRIi(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	m4, _ := strconv.ParseUint(s[6:7], 16, 8)
	i3, _ := strconv.ParseUint(s[7:9], 16, 16)
	return fmt.Sprintf("V%d,R%d,%d,%d", v1, r2, i3, m4)
}

func formatVRIj(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	m4, _ := strconv.ParseUint(s[6:7], 16, 8)
	i3, _ := strconv.ParseUint(s[7:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,%d,%d", v1, v2, i3, m4)
}

func formatVRIk(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	i5, _ := strconv.ParseUint(s[6:8], 16, 16)
	v4, _ := strconv.ParseUint(s[9:10]+s[8:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,V%d,V%d,%d", v1, v2, v3, v4, i5)
}

func formatVRIl(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	i3, _ := strconv.ParseUint(s[5:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,%d", v1, v2, i3)
}

func formatVRRa(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	m4, _ := strconv.ParseUint(s[6:7], 16, 8)
	m3, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("V%d,V%d,%d,%d", v1, v2, m3, m4)
}

func formatVRRb(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	m5, _ := strconv.ParseUint(s[6:7], 16, 8)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,V%d,V%d,%d,%d", v1, v2, v3, m4, m5)
}

func formatVRRc(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	m5, _ := strconv.ParseUint(s[7:8], 16, 8)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,V%d,V%d,%d,%d", v1, v2, v3, m4, m5)
}

func formatVRRd(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	m5, _ := strconv.ParseUint(s[5:6], 16, 8)
	v4, _ := strconv.ParseUint(s[9:10]+s[8:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,V%d,V%d,%d", v1, v2, v3, v4, m5)
}

func formatVRRe(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	v4, _ := strconv.ParseUint(s[9:10]+s[8:9], 16, 16)
	return fmt.Sprintf("V%d,V%d,V%d,V%d", v1, v2, v3, v4)
}

func formatVRRf(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	r2, _ := strconv.ParseUint(s[3:4], 16, 8)
	r3, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("V%d,R%d,R%d", v1, r2, r3)
}

func formatVRRg(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	i2, _ := strconv.ParseUint(s[5:9], 16, 16)
	return fmt.Sprintf("V%d,%d", v1, i2)
}

func formatVRRh(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	m3, _ := strconv.ParseUint(s[6:7], 16, 8)
	return fmt.Sprintf("V%d,V%d,%d", v1, v2, m3)
}

func formatVRRi(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	m3, _ := strconv.ParseUint(s[6:7], 16, 8)
	m4, _ := strconv.ParseUint(s[7:8], 16, 8)
	return fmt.Sprintf("R%d,V%d,%d,%d", r1, v2, m3, m4)
}

func formatVRRj(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[4:5], 16, 16)
	m4, _ := strconv.ParseUint(s[6:7], 16, 8)
	return fmt.Sprintf("V%d,V%d,V%d,%d", v1, v2, v3, m4)
}

func formatVRRk(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	m3, _ := strconv.ParseUint(s[6:7], 16, 8)
	return fmt.Sprintf("V%d,V%d,%d", v1, v2, m3)
}

func formatVRX(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,%d(R%d,R%d),%d", v1, d2, x2, b2, m3)
}

func formatVRS(s string) string {
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	v1, _ := strconv.ParseUint(s[9:10]+s[8:9], 16, 16)
	return fmt.Sprintf("V%d,R%d,%d(R%d)", v1, r3, d2, b2)
}

func formatVRSa(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v3, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,V%d,%d(R%d),%d", v1, v3, d2, b2, m4)
}

func formatVRSb(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	r3, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,R%d,%d(R%d),%d", v1, r3, d2, b2, m4)
}

func formatVRSc(s string) string {
	r1, _ := strconv.ParseUint(s[2:3], 16, 8)
	v3, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	m4, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,V%d,%d(R%d),%d", r1, v3, d2, b2, m4)
}

func formatVRV(s string) string {
	v1, _ := strconv.ParseUint(s[9:10]+s[2:3], 16, 16)
	v2, _ := strconv.ParseUint(s[9:10]+s[3:4], 16, 16)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	m3, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("V%d,%d(V%d,R%d),%d", v1, d2, v2, b2, m3)
}

func formatVSI(s string) string {
	i3, _ := strconv.ParseUint(s[2:4], 16, 16)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	v1, _ := strconv.ParseUint(s[9:10]+s[8:9], 16, 16)
	return fmt.Sprintf("V%d,%d(R%d),%d", v1, d2, b2, i3)
}

// ── instruction table ─────────────────────────────────────────────────────────

type inst struct {
	mnem string
	fn   func(string) string
}

var itable = map[int]inst{
	0x0101: {"PR    ", formatE},
	0x0102: {"UPT   ", formatE},
	0x0104: {"PTFF  ", formatE},
	0x0107: {"SCKPF ", formatE},
	0x010a: {"PFPO  ", formatE},
	0x010b: {"TAM   ", formatE},
	0x010c: {"SAM24 ", formatE},
	0x010d: {"SAM31 ", formatE},
	0x010e: {"SAM64 ", formatE},
	0x01ff: {"TRAP2 ", formatE},
	0x0400: {"SPM   ", formatRR3},
	0x0500: {"BALR  ", formatRR1},
	0x0600: {"BCTR  ", formatRR1},
	0x0700: {"BCR   ", formatRR2},
	0x0a00: {"SVC   ", formatI},
	0x0b00: {"BSM   ", formatRR1},
	0x0c00: {"BASSM ", formatRR1},
	0x0d00: {"BASR  ", formatRR1},
	0x0e00: {"MVCL  ", formatRR1},
	0x0f00: {"CLCL  ", formatRR1},
	0x1000: {"LPR   ", formatRR1},
	0x1100: {"LNR   ", formatRR1},
	0x1200: {"LTR   ", formatRR1},
	0x1300: {"LCR   ", formatRR1},
	0x1400: {"NR    ", formatRR1},
	0x1500: {"CLR   ", formatRR1},
	0x1600: {"OR    ", formatRR1},
	0x1700: {"XR    ", formatRR1},
	0x1800: {"LR    ", formatRR1},
	0x1900: {"CR    ", formatRR1},
	0x1a00: {"AR    ", formatRR1},
	0x1b00: {"SR    ", formatRR1},
	0x1c00: {"MR    ", formatRR1},
	0x1d00: {"DR    ", formatRR1},
	0x1e00: {"ALR   ", formatRR1},
	0x1f00: {"SLR   ", formatRR1},
	0x2000: {"LPDR  ", formatRR1},
	0x2100: {"LNDR  ", formatRR1},
	0x2200: {"LTDR  ", formatRR1},
	0x2300: {"LCDR  ", formatRR1},
	0x2400: {"HDR   ", formatRR1},
	0x2500: {"LDXR  ", formatRR1},
	0x2600: {"MXR   ", formatRR1},
	0x2700: {"MXDR  ", formatRR1},
	0x2800: {"LDR   ", formatRR1},
	0x2900: {"CDR   ", formatRR1},
	0x2a00: {"ADR   ", formatRR1},
	0x2b00: {"SDR   ", formatRR1},
	0x2c00: {"MDR   ", formatRR1},
	0x2d00: {"DDR   ", formatRR1},
	0x2e00: {"AWR   ", formatRR1},
	0x2f00: {"SWR   ", formatRR1},
	0x3000: {"LPER  ", formatRR1},
	0x3100: {"LNER  ", formatRR1},
	0x3200: {"LTER  ", formatRR1},
	0x3300: {"LCER  ", formatRR1},
	0x3400: {"HER   ", formatRR1},
	0x3500: {"LEDR  ", formatRR1},
	0x3600: {"AXR   ", formatRR1},
	0x3700: {"SXR   ", formatRR1},
	0x3800: {"LER   ", formatRR1},
	0x3900: {"CER   ", formatRR1},
	0x3a00: {"AER   ", formatRR1},
	0x3b00: {"SER   ", formatRR1},
	0x3c00: {"MDER  ", formatRR1},
	0x3d00: {"DER   ", formatRR1},
	0x3e00: {"AUR   ", formatRR1},
	0x3f00: {"SUR   ", formatRR1},
	0x4000: {"STH   ", formatRX1},
	0x4100: {"LA    ", formatRX1},
	0x4200: {"STC   ", formatRX1},
	0x4300: {"IC    ", formatRX1},
	0x4400: {"EX    ", formatRX1},
	0x4500: {"BAL   ", formatRX1},
	0x4600: {"BCT   ", formatRX1},
	0x4700: {"BC    ", formatRX2},
	0x4800: {"LH    ", formatRX1},
	0x4900: {"CH    ", formatRX1},
	0x4a00: {"AH    ", formatRX1},
	0x4b00: {"SH    ", formatRX1},
	0x4c00: {"MH    ", formatRX1},
	0x4d00: {"BAS   ", formatRX1},
	0x4e00: {"CVD   ", formatRX1},
	0x4f00: {"CVB   ", formatRX1},
	0x5000: {"ST    ", formatRX1},
	0x5100: {"LAE   ", formatRX1},
	0x5400: {"N     ", formatRX1},
	0x5500: {"CL    ", formatRX1},
	0x5600: {"O     ", formatRX1},
	0x5700: {"X     ", formatRX1},
	0x5800: {"L     ", formatRX1},
	0x5900: {"C     ", formatRX1},
	0x5a00: {"A     ", formatRX1},
	0x5b00: {"S     ", formatRX1},
	0x5c00: {"M     ", formatRX1},
	0x5d00: {"D     ", formatRX1},
	0x5e00: {"AL    ", formatRX1},
	0x5f00: {"SL    ", formatRX1},
	0x6000: {"STD   ", formatRX1},
	0x6700: {"MXD   ", formatRX1},
	0x6800: {"LD    ", formatRX1},
	0x6900: {"CD    ", formatRX1},
	0x6a00: {"AD    ", formatRX1},
	0x6b00: {"SD    ", formatRX1},
	0x6c00: {"MD    ", formatRX1},
	0x6d00: {"DD    ", formatRX1},
	0x6e00: {"AW    ", formatRX1},
	0x6f00: {"SW    ", formatRX1},
	0x7000: {"STE   ", formatRX1},
	0x7100: {"MS    ", formatRX1},
	0x7800: {"LE    ", formatRX1},
	0x7900: {"CE    ", formatRX1},
	0x7a00: {"AE    ", formatRX1},
	0x7b00: {"SE    ", formatRX1},
	0x7c00: {"MDE   ", formatRX1},
	0x7d00: {"DE    ", formatRX1},
	0x7e00: {"AU    ", formatRX1},
	0x7f00: {"SU    ", formatRX1},
	0x8000: {"SSM   ", formatS},
	0x8200: {"LPSW  ", formatS},
	0x8300: {"DIAG  ", formatE},
	0x8400: {"BRXH  ", formatRSI},
	0x8500: {"BRXLE ", formatRSI},
	0x8600: {"BXH   ", formatRS1},
	0x8700: {"BXLE  ", formatRS1},
	0x8800: {"SRL   ", formatRS1},
	0x8900: {"SLL   ", formatRS1},
	0x8a00: {"SRA   ", formatRS1},
	0x8b00: {"SLA   ", formatRS1},
	0x8c00: {"SRDL  ", formatRS1},
	0x8d00: {"SLDL  ", formatRS1},
	0x8e00: {"SRDA  ", formatRS1},
	0x8f00: {"SLDA  ", formatRS1},
	0x9000: {"STM   ", formatRS1},
	0x9100: {"TM    ", formatSI},
	0x9200: {"MVI   ", formatSI},
	0x9300: {"TS    ", formatS},
	0x9400: {"NI    ", formatSI},
	0x9500: {"CLI   ", formatSI},
	0x9600: {"OI    ", formatSI},
	0x9700: {"XI    ", formatSI},
	0x9800: {"LM    ", formatRS1},
	0x9900: {"TRACE ", formatRS1},
	0x9a00: {"LAM   ", formatRS1},
	0x9b00: {"STAM  ", formatRS1},
	0xa500: {"IIHH  ", formatRI1},
	0xa510: {"IIHL  ", formatRI1},
	0xa520: {"IILH  ", formatRI1},
	0xa530: {"IILL  ", formatRI1},
	0xa540: {"NIHH  ", formatRI1},
	0xa550: {"NIHL  ", formatRI1},
	0xa560: {"NILH  ", formatRI1},
	0xa570: {"NILL  ", formatRI1},
	0xa580: {"OIHH  ", formatRI1},
	0xa590: {"OIHL  ", formatRI1},
	0xa5a0: {"OILH  ", formatRI1},
	0xa5b0: {"OILL  ", formatRI1},
	0xa5c0: {"LLIHH ", formatRI1},
	0xa5d0: {"LLIHL ", formatRI1},
	0xa5e0: {"LLILH ", formatRI1},
	0xa5f0: {"LLILL ", formatRI1},
	0xa700: {"TMLH  ", formatRI1},
	0xa710: {"TMLL  ", formatRI1},
	0xa720: {"TMHH  ", formatRI1},
	0xa730: {"TMHL  ", formatRI1},
	0xa740: {"BRC   ", formatRI4},
	0xa750: {"BRAS  ", formatRI3},
	0xa760: {"BRCT  ", formatRI3},
	0xa770: {"BRCTG ", formatRI3},
	0xa780: {"LHI   ", formatRI3},
	0xa790: {"LGHI  ", formatRI3},
	0xa7a0: {"AHI   ", formatRI3},
	0xa7b0: {"AGHI  ", formatRI3},
	0xa7c0: {"MHI   ", formatRI3},
	0xa7d0: {"MGHI  ", formatRI3},
	0xa7e0: {"CHI   ", formatRI3},
	0xa7f0: {"CGHI  ", formatRI3},
	0xa800: {"MVCLE ", formatRS1},
	0xa900: {"CLCLE ", formatRS1},
	0xac00: {"STNSM ", formatSI},
	0xad00: {"STOSM ", formatSI},
	0xae00: {"SIGP  ", formatRS1},
	0xaf00: {"MC    ", formatSI},
	0xb100: {"LRA   ", formatRX1},
	0xb200: {"LBEAR ", formatS},
	0xb201: {"STBEAR ", formatS},
	0xb202: {"STIDP ", formatS},
	0xb204: {"SCK   ", formatS},
	0xb205: {"STCK  ", formatS},
	0xb206: {"SCKC  ", formatS},
	0xb207: {"STCKC ", formatS},
	0xb208: {"SPT   ", formatS},
	0xb209: {"STPT  ", formatS},
	0xb20a: {"SPKA  ", formatS},
	0xb20b: {"IPK   ", formatS},
	0xb20d: {"PTLB  ", formatS},
	0xb210: {"SPX   ", formatS},
	0xb211: {"STPX  ", formatS},
	0xb212: {"STAP  ", formatS},
	0xb218: {"PC    ", formatS},
	0xb219: {"SAC   ", formatS},
	0xb21a: {"CFC   ", formatS},
	0xb221: {"IPTE  ", formatRRFa},
	0xb222: {"IPM   ", formatRRE2},
	0xb223: {"IVSK  ", formatRRE1},
	0xb224: {"IAC   ", formatRRE2},
	0xb225: {"SSAR  ", formatRRE2},
	0xb226: {"EPAR  ", formatRRE2},
	0xb227: {"ESAR  ", formatRRE2},
	0xb228: {"PT    ", formatRRE1},
	0xb229: {"ISKE  ", formatRRE1},
	0xb22a: {"RRBE  ", formatRRE1},
	0xb22b: {"SSKE  ", formatRRF4},
	0xb22c: {"TB    ", formatRRE1},
	0xb22d: {"DXR   ", formatRRE1},
	0xb22e: {"PGIN  ", formatRRE1},
	0xb22f: {"PGOUT ", formatRRE1},
	0xb230: {"CSCH  ", formatE},
	0xb231: {"HSCH  ", formatE},
	0xb232: {"MSCH  ", formatS},
	0xb233: {"SSCH  ", formatS},
	0xb234: {"STSCH ", formatS},
	0xb235: {"TSCH  ", formatS},
	0xb236: {"TPI   ", formatS},
	0xb237: {"SAL   ", formatE},
	0xb238: {"RSCH  ", formatE},
	0xb239: {"STCRW ", formatS},
	0xb23a: {"STPCS ", formatS},
	0xb23b: {"RCHP  ", formatE},
	0xb23c: {"SCHM  ", formatE},
	0xb240: {"BAKR  ", formatRRE1},
	0xb241: {"CKSM  ", formatRRE1},
	0xb244: {"SQDR  ", formatRRE1},
	0xb245: {"SQER  ", formatRRE1},
	0xb246: {"STURA ", formatRRE1},
	0xb247: {"MSTA  ", formatRRE2},
	0xb248: {"PALB  ", formatE},
	0xb249: {"EREG  ", formatRRE1},
	0xb24a: {"ESTA  ", formatRRE1},
	0xb24b: {"LURA  ", formatRRE1},
	0xb24c: {"TAR   ", formatRRE1},
	0xb24d: {"CPYA  ", formatRRE1},
	0xb24e: {"SAR   ", formatRRE1},
	0xb24f: {"EAR   ", formatRRE1},
	0xb250: {"CSP   ", formatRRE1},
	0xb252: {"MSR   ", formatRRE1},
	0xb254: {"MVPG  ", formatRRE1},
	0xb255: {"MVST  ", formatRRE1},
	0xb257: {"CUSE  ", formatRRE1},
	0xb258: {"BSG   ", formatRRE1},
	0xb25a: {"BSA   ", formatRRE1},
	0xb25d: {"CLST  ", formatRRE1},
	0xb25e: {"SRST  ", formatRRE1},
	0xb263: {"CMPSC ", formatRRE1},
	0xb276: {"XSCH  ", formatE},
	0xb277: {"RP    ", formatS},
	0xb278: {"STCKE ", formatS},
	0xb279: {"SACF  ", formatS},
	0xb27c: {"STCKF ", formatS},
	0xb27d: {"STSI  ", formatS},
	0xb28f: {"QPACI ", formatS},
	0xb299: {"SRNM  ", formatS},
	0xb29c: {"STFPC ", formatS},
	0xb29d: {"LFPC  ", formatS},
	0xb2a5: {"TRE   ", formatRRE1},
	0xb2a6: {"CU21  ", formatRRF4},
	0xb2a7: {"CU12  ", formatRRF4},
	0xb2b0: {"STFLE ", formatS},
	0xb2b1: {"STFL  ", formatS},
	0xb2b2: {"LPSWE ", formatS},
	0xb2b8: {"SRNMB ", formatS},
	0xb2b9: {"SRNMT ", formatS},
	0xb2bd: {"LFAS  ", formatS},
	0xb2e8: {"PPA   ", formatRRF4},
	0xb2ec: {"ETND  ", formatRRE2},
	0xb2f8: {"TEND  ", formatE},
	0xb2fa: {"NIAI  ", formatIE},
	0xb2fc: {"TABORT ", formatS},
	0xb2ff: {"TRAP4 ", formatS},
	0xb300: {"LPEBR ", formatRRE1},
	0xb301: {"LNEBR ", formatRRE1},
	0xb302: {"LTEBR ", formatRRE1},
	0xb303: {"LCEBR ", formatRRE1},
	0xb304: {"LDEBR ", formatRRE1},
	0xb305: {"LXDBR ", formatRRE1},
	0xb306: {"LXEBR ", formatRRE1},
	0xb307: {"MXDBR ", formatRRE1},
	0xb308: {"KEBR  ", formatRRE1},
	0xb309: {"CEBR  ", formatRRE1},
	0xb30a: {"AEBR  ", formatRRE1},
	0xb30b: {"SEBR  ", formatRRE1},
	0xb30c: {"MDEBR ", formatRRE1},
	0xb30d: {"DEBR  ", formatRRE1},
	0xb30e: {"MAEBR ", formatRRF5},
	0xb30f: {"MSEBR ", formatRRF5},
	0xb310: {"LPDBR ", formatRRE1},
	0xb311: {"LNDBR ", formatRRE1},
	0xb312: {"LTDBR ", formatRRE1},
	0xb313: {"LCDBR ", formatRRE1},
	0xb314: {"SQEBR ", formatRRE1},
	0xb315: {"SQDBR ", formatRRE1},
	0xb316: {"SQXBR ", formatRRE1},
	0xb317: {"MEEBR ", formatRRE1},
	0xb318: {"KDBR  ", formatRRE1},
	0xb319: {"CDBR  ", formatRRE1},
	0xb31a: {"ADBR  ", formatRRE1},
	0xb31b: {"SDBR  ", formatRRE1},
	0xb31c: {"MDBR  ", formatRRE1},
	0xb31d: {"DDBR  ", formatRRE1},
	0xb31e: {"MADBR ", formatRRF5},
	0xb31f: {"MSDBR ", formatRRF5},
	0xb324: {"LDER  ", formatRRE1},
	0xb325: {"LXDR  ", formatRRE1},
	0xb326: {"LXER  ", formatRRE1},
	0xb32e: {"MAER  ", formatRRF5},
	0xb32f: {"MSER  ", formatRRF5},
	0xb336: {"SQXR  ", formatRRE1},
	0xb337: {"MEER  ", formatRRE1},
	0xb338: {"MAYLR ", formatRRF5},
	0xb339: {"MYLR  ", formatRRF5},
	0xb33a: {"MAYR  ", formatRRF5},
	0xb33b: {"MYR   ", formatRRF5},
	0xb33c: {"MAYHR ", formatRRF5},
	0xb33d: {"MYHR  ", formatRRF5},
	0xb33e: {"MADR  ", formatRRF5},
	0xb33f: {"MSDR  ", formatRRF5},
	0xb340: {"LPXBR ", formatRRE1},
	0xb341: {"LNXBR ", formatRRE1},
	0xb342: {"LTXBR ", formatRRE1},
	0xb343: {"LCXBR ", formatRRE1},
	0xb344: {"LEDBR ", formatRRE1},
	0xb345: {"LDXBR ", formatRRE1},
	0xb346: {"LEXBR ", formatRRE1},
	0xb347: {"FIXBR ", formatRRF2},
	0xb348: {"KXBR  ", formatRRE1},
	0xb349: {"CXBR  ", formatRRE1},
	0xb34a: {"AXBR  ", formatRRE1},
	0xb34b: {"SXBR  ", formatRRE1},
	0xb34c: {"MXBR  ", formatRRE1},
	0xb34d: {"DXBR  ", formatRRE1},
	0xb350: {"TBEBR ", formatRRF2},
	0xb351: {"TBDR  ", formatRRF2},
	0xb353: {"DIEBR ", formatRRF3},
	0xb357: {"FIEBR ", formatRRF2},
	0xb358: {"THDER ", formatRRE1},
	0xb359: {"THDR  ", formatRRE1},
	0xb35b: {"DIDBR ", formatRRF3},
	0xb35f: {"FIDBR ", formatRRF2},
	0xb360: {"LPXR  ", formatRRE1},
	0xb361: {"LNXR  ", formatRRE1},
	0xb362: {"LTXR  ", formatRRE1},
	0xb363: {"LCXR  ", formatRRE1},
	0xb365: {"LXR   ", formatRRE1},
	0xb366: {"LEXR  ", formatRRE1},
	0xb367: {"FIXR  ", formatRRE1},
	0xb369: {"CXR   ", formatRRE1},
	0xb370: {"LPDFR ", formatRRE1},
	0xb371: {"LNDFR ", formatRRE1},
	0xb372: {"CPSDR ", formatRRF3},
	0xb373: {"LCDFR ", formatRRE1},
	0xb374: {"LZER  ", formatRRE2},
	0xb375: {"LZDR  ", formatRRE2},
	0xb376: {"LZXR  ", formatRRE2},
	0xb377: {"FIER  ", formatRRE1},
	0xb37f: {"FIDR  ", formatRRE1},
	0xb384: {"SFPC  ", formatRRE2},
	0xb385: {"SFASR ", formatRRE2},
	0xb38c: {"EFPC  ", formatRRE2},
	0xb390: {"CELFBR ", formatRRF6},
	0xb391: {"CDLFBR ", formatRRF6},
	0xb392: {"CXLFBR ", formatRRF6},
	0xb394: {"CEFBR ", formatRRE1},
	0xb395: {"CDFBR ", formatRRE1},
	0xb396: {"CXFBR ", formatRRE1},
	0xb398: {"CFEBR ", formatRRF2},
	0xb399: {"CFDBR ", formatRRF2},
	0xb39a: {"CFXBR ", formatRRF2},
	0xb39c: {"CLFEBR ", formatRRF6},
	0xb39d: {"CLFDBR ", formatRRF6},
	0xb39e: {"CLFXBR ", formatRRF6},
	0xb3a0: {"CELGBR ", formatRRF6},
	0xb3a1: {"CDLGBR ", formatRRF6},
	0xb3a2: {"CXLGBR ", formatRRF6},
	0xb3a4: {"CEGBR ", formatRRE1},
	0xb3a5: {"CDGBR ", formatRRE1},
	0xb3a6: {"CXGBR ", formatRRE1},
	0xb3a8: {"CGEBR ", formatRRF2},
	0xb3a9: {"CGDBR ", formatRRF2},
	0xb3aa: {"CGXBR ", formatRRF2},
	0xb3ac: {"CLGEBR ", formatRRF6},
	0xb3ad: {"CLGDBR ", formatRRF6},
	0xb3ae: {"CLGXBR ", formatRRF6},
	0xb3b4: {"CEFR  ", formatRRE1},
	0xb3b5: {"CDFR  ", formatRRE1},
	0xb3b6: {"CXFR  ", formatRRE1},
	0xb3b8: {"CFER  ", formatRRF2},
	0xb3b9: {"CFDR  ", formatRRF2},
	0xb3ba: {"CFXR  ", formatRRF2},
	0xb3c1: {"LDGR  ", formatRRE1},
	0xb3c4: {"CEGR  ", formatRRE1},
	0xb3c5: {"CDGR  ", formatRRE1},
	0xb3c6: {"CXGR  ", formatRRE1},
	0xb3c8: {"CGER  ", formatRRF2},
	0xb3c9: {"CGDR  ", formatRRF2},
	0xb3ca: {"CGXR  ", formatRRF2},
	0xb3cd: {"LGDR  ", formatRRE1},
	0xb3d0: {"MDTR  ", formatRRF8},
	0xb3d1: {"DDTR  ", formatRRF8},
	0xb3d2: {"ADTR  ", formatRRF8},
	0xb3d3: {"SDTR  ", formatRRF8},
	0xb3d4: {"LDETR ", formatRRF9},
	0xb3d5: {"LEDTR ", formatRRF6},
	0xb3d6: {"LTDTR ", formatRRE1},
	0xb3d7: {"FIDTR ", formatRRF6},
	0xb3d8: {"MXTR  ", formatRRF8},
	0xb3d9: {"DXTR  ", formatRRF8},
	0xb3da: {"AXTR  ", formatRRF8},
	0xb3db: {"SXTR  ", formatRRF8},
	0xb3dc: {"LXDTR ", formatRRF9},
	0xb3dd: {"LDXTR ", formatRRF6},
	0xb3de: {"LTXTR ", formatRRE1},
	0xb3df: {"FIXTR ", formatRRF6},
	0xb3e0: {"KDTR  ", formatRRE1},
	0xb3e1: {"CGDTR ", formatRRF6},
	0xb3e2: {"CUDTR ", formatRRE1},
	0xb3e3: {"CSDTR ", formatRRF9},
	0xb3e4: {"CDTR  ", formatRRE1},
	0xb3e5: {"EEDTR ", formatRRE1},
	0xb3e7: {"ESDTR ", formatRRE1},
	0xb3e8: {"KXTR  ", formatRRE1},
	0xb3e9: {"CGXTR ", formatRRF6},
	0xb3ea: {"CUXTR ", formatRRE1},
	0xb3eb: {"CSXTR ", formatRRF9},
	0xb3ec: {"CXTR  ", formatRRE1},
	0xb3ed: {"EEXTR ", formatRRE1},
	0xb3ef: {"ESXTR ", formatRRE1},
	0xb3f1: {"CDGTR ", formatRRE1},
	0xb3f2: {"CDUTR ", formatRRE1},
	0xb3f3: {"CDSTR ", formatRRE1},
	0xb3f4: {"CEDTR ", formatRRE1},
	0xb3f5: {"QADTR ", formatRRF3},
	0xb3f6: {"IEDTR ", formatRRF3},
	0xb3f7: {"RRDTR ", formatRRF3},
	0xb3f9: {"CXGTR ", formatRRE1},
	0xb3fa: {"CXUTR ", formatRRE1},
	0xb3fb: {"CXSTR ", formatRRE1},
	0xb3fc: {"CEXTR ", formatRRE1},
	0xb3fd: {"QAXTR ", formatRRF3},
	0xb3fe: {"IEXTR ", formatRRF3},
	0xb3ff: {"RRXTR ", formatRRF3},
	0xb600: {"STCTL ", formatRS1},
	0xb700: {"LCTL  ", formatRS1},
	0xb900: {"LPGR  ", formatRRE1},
	0xb901: {"LNGR  ", formatRRE1},
	0xb902: {"LTGR  ", formatRRE1},
	0xb903: {"LCGR  ", formatRRE1},
	0xb904: {"LGR   ", formatRRE1},
	0xb905: {"LURAG ", formatRRE1},
	0xb906: {"LGBR  ", formatRRE1},
	0xb907: {"LGHR  ", formatRRE1},
	0xb908: {"AGR   ", formatRRE1},
	0xb909: {"SGR   ", formatRRE1},
	0xb90a: {"ALGR  ", formatRRE1},
	0xb90b: {"SLGR  ", formatRRE1},
	0xb90c: {"MSGR  ", formatRRE1},
	0xb90d: {"DSGR  ", formatRRE1},
	0xb90e: {"EREGG ", formatRRE1},
	0xb90f: {"LRVGR ", formatRRE1},
	0xb910: {"LPGFR ", formatRRE1},
	0xb911: {"LNGFR ", formatRRE1},
	0xb912: {"LTGFR ", formatRRE1},
	0xb913: {"LCGFR ", formatRRE1},
	0xb914: {"LGFR  ", formatRRE1},
	0xb916: {"LLGFR ", formatRRE1},
	0xb917: {"LLGTR ", formatRRE1},
	0xb918: {"AGFR  ", formatRRE1},
	0xb919: {"SGFR  ", formatRRE1},
	0xb91a: {"ALGFR ", formatRRE1},
	0xb91b: {"SLGFR ", formatRRE1},
	0xb91c: {"MSGFR ", formatRRE1},
	0xb91d: {"DSGFR ", formatRRE1},
	0xb91e: {"KMAC  ", formatRRE1},
	0xb91f: {"LRVR  ", formatRRE1},
	0xb920: {"CGR   ", formatRRE1},
	0xb921: {"CLGR  ", formatRRE1},
	0xb925: {"STURG ", formatRRE1},
	0xb926: {"LBR   ", formatRRE1},
	0xb927: {"LHR   ", formatRRE1},
	0xb928: {"PCKMO ", formatRRE1},
	0xb929: {"KMA   ", formatRRF3},
	0xb92a: {"KMF   ", formatRRE1},
	0xb92b: {"KMO   ", formatRRE1},
	0xb92c: {"PCC   ", formatRRE1},
	0xb92d: {"KMCTR ", formatRRF3},
	0xb92e: {"KM    ", formatRRE1},
	0xb92f: {"KMC   ", formatRRE1},
	0xb930: {"CGFR  ", formatRRE1},
	0xb931: {"CLGFR ", formatRRE1},
	0xb938: {"SORTL ", formatRRE1},
	0xb939: {"DFLTCC ", formatRRF7},
	0xb93a: {"KDSA  ", formatRRE1},
	0xb93b: {"NNPA  ", formatE},
	0xb93c: {"PPNO  ", formatRRE1},
	0xb93e: {"KIMD  ", formatRRF4},
	0xb93f: {"KLMD  ", formatRRF4},
	0xb941: {"CFDTR ", formatRRF6},
	0xb942: {"CLGDTR ", formatRRF6},
	0xb943: {"CLFDTR ", formatRRF6},
	0xb946: {"BCTGR ", formatRRE1},
	0xb949: {"CFXTR ", formatRRF6},
	0xb94a: {"CLGXTR ", formatRRF6},
	0xb94b: {"CLFXTR ", formatRRF6},
	0xb951: {"CDFTR ", formatRRF6},
	0xb952: {"CDLGTR ", formatRRF6},
	0xb953: {"CDLFTR ", formatRRF6},
	0xb959: {"CXFTR ", formatRRF6},
	0xb95a: {"CXLGTR ", formatRRF6},
	0xb95b: {"CXLFTR ", formatRRF6},
	0xb960: {"CGRT  ", formatRRF4},
	0xb961: {"CLGRT ", formatRRF4},
	0xb964: {"NNGRK ", formatRRF7},
	0xb965: {"OCGRK ", formatRRF7},
	0xb966: {"NOGRK ", formatRRF7},
	0xb967: {"NXGRK ", formatRRF7},
	0xb968: {"CLZG  ", formatRRE1},
	0xb969: {"CTZG  ", formatRRE1},
	0xb96c: {"BEXTG ", formatRRF7},
	0xb96d: {"BDEPG ", formatRRF7},
	0xb972: {"CRT   ", formatRRF4},
	0xb973: {"CLRT  ", formatRRF4},
	0xb974: {"NNRK  ", formatRRF7},
	0xb975: {"OCRK  ", formatRRF7},
	0xb976: {"NORK  ", formatRRF7},
	0xb977: {"NXRK  ", formatRRF7},
	0xb980: {"NGR   ", formatRRE1},
	0xb981: {"OGR   ", formatRRE1},
	0xb982: {"XGR   ", formatRRE1},
	0xb983: {"FLOGR ", formatRRE1},
	0xb984: {"LLGCR ", formatRRE1},
	0xb985: {"LLGHR ", formatRRE1},
	0xb986: {"MLGR  ", formatRRE1},
	0xb987: {"DLGR  ", formatRRE1},
	0xb988: {"ALCGR ", formatRRE1},
	0xb989: {"SLBGR ", formatRRE1},
	0xb98a: {"CSPG  ", formatRRE1},
	0xb98b: {"RDP   ", formatRRF3},
	0xb98d: {"EPSW  ", formatRRE1},
	0xb98e: {"IDTE  ", formatRRF3},
	0xb98f: {"CRDTE ", formatRRF3},
	0xb990: {"TRTT  ", formatRRF2},
	0xb991: {"TRTO  ", formatRRF2},
	0xb992: {"TROT  ", formatRRF2},
	0xb993: {"TROO  ", formatRRF2},
	0xb994: {"LLCR  ", formatRRE1},
	0xb995: {"LLHR  ", formatRRE1},
	0xb996: {"MLR   ", formatRRE1},
	0xb997: {"DLR   ", formatRRE1},
	0xb998: {"ALCR  ", formatRRE1},
	0xb999: {"SLBR  ", formatRRE1},
	0xb99a: {"EPAIR ", formatRRE2},
	0xb99b: {"ESAIR ", formatRRE2},
	0xb99d: {"ESEA  ", formatRRE2},
	0xb99e: {"PTI   ", formatRRE1},
	0xb99f: {"SSAIR ", formatRRE2},
	0xb9a1: {"TPEI  ", formatRRE1},
	0xb9a2: {"PTF   ", formatRRE2},
	0xb9aa: {"LPTEA ", formatRRF3},
	0xb9ac: {"IRBM  ", formatRRE1},
	0xb9ae: {"RRBM  ", formatRRE1},
	0xb9af: {"PFMF  ", formatRRE1},
	0xb9b0: {"CU14  ", formatRRF4},
	0xb9b1: {"CU24  ", formatRRF4},
	0xb9b2: {"CU41  ", formatRRE1},
	0xb9b3: {"CU42  ", formatRRE1},
	0xb9bd: {"TRTRE ", formatRRF4},
	0xb9be: {"SRSTU ", formatRRE1},
	0xb9bf: {"TRTE  ", formatRRF4},
	0xb9c0: {"SELFHR ", formatRRF8},
	0xb9c8: {"AHHHR ", formatRRF7},
	0xb9c9: {"SHHHR ", formatRRF7},
	0xb9ca: {"ALHHHR ", formatRRF7},
	0xb9cb: {"SLHHHR ", formatRRF7},
	0xb9cd: {"CHHR  ", formatRRE1},
	0xb9cf: {"CLHHR ", formatRRE1},
	0xb9d8: {"AHHLR ", formatRRF7},
	0xb9d9: {"SHHLR ", formatRRF7},
	0xb9da: {"ALHHLR ", formatRRF7},
	0xb9db: {"SLHHLR ", formatRRF7},
	0xb9dd: {"CHLR  ", formatRRE1},
	0xb9df: {"CLHLR ", formatRRE1},
	0xb9e0: {"LOCFHR ", formatRRF4},
	0xb9e1: {"POPCNT ", formatRRF4},
	0xb9e2: {"LOCGR ", formatRRF4},
	0xb9e3: {"SELGR ", formatRRF8},
	0xb9e4: {"NGRK  ", formatRRF7},
	0xb9e5: {"NCGRK ", formatRRF7},
	0xb9e6: {"OGRK  ", formatRRF7},
	0xb9e7: {"XGRK  ", formatRRF7},
	0xb9e8: {"AGRK  ", formatRRF7},
	0xb9e9: {"SGRK  ", formatRRF7},
	0xb9ea: {"ALGRK ", formatRRF7},
	0xb9eb: {"SLGRK ", formatRRF7},
	0xb9ec: {"MGRK  ", formatRRF7},
	0xb9ed: {"MSGRKC ", formatRRF7},
	0xb9f0: {"SELR  ", formatRRF8},
	0xb9f2: {"LOCR  ", formatRRF4},
	0xb9f4: {"NRK   ", formatRRF7},
	0xb9f5: {"NCRK  ", formatRRF7},
	0xb9f6: {"ORK   ", formatRRF7},
	0xb9f7: {"XRK   ", formatRRF7},
	0xb9f8: {"ARK   ", formatRRF7},
	0xb9f9: {"SRK   ", formatRRF7},
	0xb9fa: {"ALRK  ", formatRRF7},
	0xb9fb: {"SLRK  ", formatRRF7},
	0xb9fd: {"MSRKC ", formatRRF7},
	0xba00: {"CS    ", formatRS1},
	0xbb00: {"CDS   ", formatRS1},
	0xbd00: {"CLM   ", formatRS2},
	0xbe00: {"STCM  ", formatRS2},
	0xbf00: {"ICM   ", formatRS2},
	0xc000: {"LARL  ", formatRIL3},
	0xc010: {"LGFI  ", formatRIL3},
	0xc040: {"BRCL  ", formatRIL4},
	0xc050: {"BRASL ", formatRIL3},
	0xc060: {"XIHF  ", formatRIL1},
	0xc070: {"XILF  ", formatRIL1},
	0xc080: {"IIHF  ", formatRIL1},
	0xc090: {"IILF  ", formatRIL1},
	0xc0a0: {"NIHF  ", formatRIL1},
	0xc0b0: {"NILF  ", formatRIL1},
	0xc0c0: {"OIHF  ", formatRIL1},
	0xc0d0: {"OILF  ", formatRIL1},
	0xc0e0: {"LLIHF ", formatRIL1},
	0xc0f0: {"LLILF ", formatRIL1},
	0xc200: {"MSGFI ", formatRIL3},
	0xc210: {"MSFI  ", formatRIL3},
	0xc240: {"SLGFI ", formatRIL1},
	0xc250: {"SLFI  ", formatRIL1},
	0xc280: {"AGFI  ", formatRIL3},
	0xc290: {"AFI   ", formatRIL3},
	0xc2a0: {"ALGFI ", formatRIL1},
	0xc2b0: {"ALFI  ", formatRIL1},
	0xc2c0: {"CGFI  ", formatRIL3},
	0xc2d0: {"CFI   ", formatRIL3},
	0xc2e0: {"CLGFI ", formatRIL1},
	0xc2f0: {"CLFI  ", formatRIL1},
	0xc420: {"LLHRL ", formatRIL3},
	0xc440: {"LGHRL ", formatRIL3},
	0xc450: {"LHRL  ", formatRIL3},
	0xc460: {"LLGHRL ", formatRIL3},
	0xc470: {"STHRL ", formatRIL3},
	0xc480: {"LGRL  ", formatRIL3},
	0xc4b0: {"STGRL ", formatRIL3},
	0xc4c0: {"LGFRL ", formatRIL3},
	0xc4d0: {"LRL   ", formatRIL3},
	0xc4e0: {"LLGFRL ", formatRIL3},
	0xc4f0: {"STRL  ", formatRIL3},
	0xc500: {"BPRP  ", formatMII},
	0xc600: {"EXRL  ", formatRIL3},
	0xc620: {"PFDRL ", formatRIL4},
	0xc640: {"CGHRL ", formatRIL3},
	0xc650: {"CHRL  ", formatRIL3},
	0xc660: {"CLGHRL ", formatRIL3},
	0xc670: {"CLHRL ", formatRIL3},
	0xc680: {"CGRL  ", formatRIL3},
	0xc6a0: {"CLGRL ", formatRIL3},
	0xc6c0: {"CGFRL ", formatRIL3},
	0xc6d0: {"CRL   ", formatRIL3},
	0xc6e0: {"CLGFRL ", formatRIL3},
	0xc6f0: {"CLRL  ", formatRIL3},
	0xc700: {"BPP   ", formatSMI},
	0xc800: {"MVCOS ", formatSSF},
	0xc810: {"ECTG  ", formatSSF},
	0xc820: {"CSST  ", formatSSF},
	0xc840: {"LPD   ", formatSSG},
	0xc850: {"LPDG  ", formatSSG},
	0xc860: {"CAL   ", formatSSG},
	0xc870: {"CALG  ", formatSSG},
	0xc8f0: {"CALGF ", formatSSG},
	0xcc60: {"BRCTH ", formatRIL3},
	0xcc80: {"AIH   ", formatRIL1},
	0xcca0: {"ALSIH ", formatRIL1},
	0xccb0: {"ALSIHN ", formatRIL1},
	0xccd0: {"CIH   ", formatRIL1},
	0xccf0: {"CLIH  ", formatRIL1},
	0xd000: {"TRTR  ", formatSS1},
	0xd100: {"MVN   ", formatSS1},
	0xd200: {"MVC   ", formatSS1},
	0xd300: {"MVZ   ", formatSS1},
	0xd400: {"NC    ", formatSS1},
	0xd500: {"CLC   ", formatSS1},
	0xd600: {"OC    ", formatSS1},
	0xd700: {"XC    ", formatSS1},
	0xd900: {"MVCK  ", formatSS4},
	0xda00: {"MVCP  ", formatSS4},
	0xdb00: {"MVCS  ", formatSS4},
	0xdc00: {"TR    ", formatSS1},
	0xdd00: {"TRT   ", formatSS1},
	0xde00: {"ED    ", formatSS1},
	0xdf00: {"EDMK  ", formatSS1},
	0xe100: {"PKU   ", formatSS6},
	0xe200: {"UNPKU ", formatSS1},
	0xe302: {"LTG   ", formatRXY},
	0xe303: {"LRAG  ", formatRXY},
	0xe304: {"LG    ", formatRXY},
	0xe306: {"CVBY  ", formatRXY},
	0xe308: {"AG    ", formatRXY},
	0xe309: {"SG    ", formatRXY},
	0xe30a: {"ALG   ", formatRXY},
	0xe30b: {"SLG   ", formatRXY},
	0xe30c: {"MSG   ", formatRXY},
	0xe30d: {"DSG   ", formatRXY},
	0xe30e: {"CVBG  ", formatRXY},
	0xe30f: {"LRVG  ", formatRXY},
	0xe312: {"LT    ", formatRXY},
	0xe313: {"LRAY  ", formatRXY},
	0xe314: {"LGF   ", formatRXY},
	0xe315: {"LGH   ", formatRXY},
	0xe316: {"LLGF  ", formatRXY},
	0xe317: {"LLGT  ", formatRXY},
	0xe318: {"AGF   ", formatRXY},
	0xe319: {"SGF   ", formatRXY},
	0xe31a: {"ALGF  ", formatRXY},
	0xe31b: {"SLGF  ", formatRXY},
	0xe31c: {"MSGF  ", formatRXY},
	0xe31d: {"DSGF  ", formatRXY},
	0xe31e: {"LRV   ", formatRXY},
	0xe31f: {"LRVH  ", formatRXY},
	0xe320: {"CG    ", formatRXY},
	0xe321: {"CLG   ", formatRXY},
	0xe324: {"STG   ", formatRXY},
	0xe325: {"NTSTG ", formatRXY},
	0xe326: {"CVDY  ", formatRXY},
	0xe32a: {"LZRG  ", formatRXY},
	0xe32e: {"CVDG  ", formatRXY},
	0xe32f: {"STRVG ", formatRXY},
	0xe330: {"CGF   ", formatRXY},
	0xe331: {"CLGF  ", formatRXY},
	0xe332: {"LTGF  ", formatRXY},
	0xe334: {"CGH   ", formatRXY},
	0xe336: {"PFD   ", formatRXG},
	0xe338: {"AGH   ", formatRXY},
	0xe339: {"SGH   ", formatRXY},
	0xe33a: {"LLZRGF ", formatRXY},
	0xe33b: {"LZRF  ", formatRXY},
	0xe33c: {"MGH   ", formatRXY},
	0xe33e: {"STRV  ", formatRXY},
	0xe33f: {"STRVH ", formatRXY},
	0xe346: {"BCTG  ", formatRXY},
	0xe347: {"BIC   ", formatRXG},
	0xe348: {"LLGFSG ", formatRXY},
	0xe349: {"STGSC ", formatRXY},
	0xe34c: {"LGG   ", formatRXY},
	0xe34d: {"LGSC  ", formatRXY},
	0xe350: {"STY   ", formatRXY},
	0xe351: {"MSY   ", formatRXY},
	0xe353: {"MSC   ", formatRXY},
	0xe354: {"NY    ", formatRXY},
	0xe355: {"CLY   ", formatRXY},
	0xe356: {"OY    ", formatRXY},
	0xe357: {"XY    ", formatRXY},
	0xe358: {"LY    ", formatRXY},
	0xe359: {"CY    ", formatRXY},
	0xe35a: {"AY    ", formatRXY},
	0xe35b: {"SY    ", formatRXY},
	0xe35c: {"MFY   ", formatRXY},
	0xe35e: {"ALY   ", formatRXY},
	0xe35f: {"SLY   ", formatRXY},
	0xe360: {"LXAB  ", formatRXY},
	0xe361: {"LLXAB ", formatRXY},
	0xe362: {"LXAH  ", formatRXY},
	0xe363: {"LLXAH ", formatRXY},
	0xe364: {"LXAF  ", formatRXY},
	0xe365: {"LLXAF ", formatRXY},
	0xe366: {"LXAG  ", formatRXY},
	0xe367: {"LLXAG ", formatRXY},
	0xe368: {"LXAQ  ", formatRXY},
	0xe369: {"LLXAQ ", formatRXY},
	0xe370: {"STHY  ", formatRXY},
	0xe371: {"LAY   ", formatRXY},
	0xe372: {"STCY  ", formatRXY},
	0xe373: {"ICY   ", formatRXY},
	0xe375: {"LAEY  ", formatRXY},
	0xe376: {"LB    ", formatRXY},
	0xe377: {"LGB   ", formatRXY},
	0xe378: {"LHY   ", formatRXY},
	0xe379: {"CHY   ", formatRXY},
	0xe37a: {"AHY   ", formatRXY},
	0xe37b: {"SHY   ", formatRXY},
	0xe37c: {"MHY   ", formatRXY},
	0xe380: {"NG    ", formatRXY},
	0xe381: {"OG    ", formatRXY},
	0xe382: {"XG    ", formatRXY},
	0xe383: {"MSGC  ", formatRXY},
	0xe384: {"MG    ", formatRXY},
	0xe385: {"LGAT  ", formatRXY},
	0xe386: {"MLG   ", formatRXY},
	0xe387: {"DLG   ", formatRXY},
	0xe388: {"ALCG  ", formatRXY},
	0xe389: {"SLBG  ", formatRXY},
	0xe38e: {"STPQ  ", formatRXY},
	0xe38f: {"LPQ   ", formatRXY},
	0xe390: {"LLGC  ", formatRXY},
	0xe391: {"LLGH  ", formatRXY},
	0xe394: {"LLC   ", formatRXY},
	0xe395: {"LLH   ", formatRXY},
	0xe396: {"ML    ", formatRXY},
	0xe397: {"DL    ", formatRXY},
	0xe398: {"ALC   ", formatRXY},
	0xe399: {"SLB   ", formatRXY},
	0xe39c: {"LLGTAT ", formatRXY},
	0xe39d: {"LLGFAT ", formatRXY},
	0xe39f: {"LAT   ", formatRXY},
	0xe3c0: {"LBH   ", formatRXY},
	0xe3c2: {"LLCH  ", formatRXY},
	0xe3c3: {"STCH  ", formatRXY},
	0xe3c4: {"LHH   ", formatRXY},
	0xe3c6: {"LLHH  ", formatRXY},
	0xe3c7: {"STHH  ", formatRXY},
	0xe3c8: {"LFHAT ", formatRXY},
	0xe3ca: {"LFH   ", formatRXY},
	0xe3cb: {"STFH  ", formatRXY},
	0xe3cd: {"CHF   ", formatRXY},
	0xe3cf: {"CLHF  ", formatRXY},
	0xe500: {"LASP  ", formatSSE},
	0xe501: {"TPROT ", formatSSE},
	0xe502: {"STRAG ", formatSSE},
	0xe50a: {"MVCRL ", formatSSE},
	0xe50e: {"MVCSK ", formatSSE},
	0xe50f: {"MVCDK ", formatSSE},
	0xe544: {"MVHHI ", formatSIL},
	0xe548: {"MVGHI ", formatSIL},
	0xe54c: {"MVHI  ", formatSIL},
	0xe554: {"CHHSI ", formatSIL},
	0xe555: {"CLHHSI ", formatSIL},
	0xe558: {"CGHSI ", formatSIL},
	0xe559: {"CLGHSI ", formatSIL},
	0xe55c: {"CHSI  ", formatSIL},
	0xe55d: {"CLFHSI ", formatSIL},
	0xe560: {"TBEGIN ", formatSIL},
	0xe561: {"TBEGINC ", formatSIL},
	0xe601: {"VLEBRH ", formatVRX},
	0xe602: {"VLEBRG ", formatVRX},
	0xe603: {"VLEBRF ", formatVRX},
	0xe604: {"VLLEBRZ ", formatVRX},
	0xe605: {"VLBRREP ", formatVRX},
	0xe606: {"VLBR  ", formatVRX},
	0xe607: {"VLER  ", formatVRX},
	0xe609: {"VSTEBRH ", formatVRX},
	0xe60a: {"VSTEBRG ", formatVRX},
	0xe60b: {"VSTEBRF ", formatVRX},
	0xe60e: {"VSTBR ", formatVRX},
	0xe60f: {"VSTER ", formatVRX},
	0xe634: {"VPKZ  ", formatVSI},
	0xe635: {"VLRL  ", formatVSI},
	0xe637: {"VLRLR ", formatVRS},
	0xe63c: {"VUPKZ ", formatVSI},
	0xe63d: {"VSTRL ", formatVSI},
	0xe63f: {"VSTRLR ", formatVRS},
	0xe649: {"VLIP  ", formatVRIh},
	0xe64a: {"VCVDQ ", formatVRIj},
	0xe64e: {"VCVBQ ", formatVRRk},
	0xe650: {"VCVB  ", formatVRRi},
	0xe651: {"VCLZDP ", formatVRRk},
	0xe652: {"VCVBG ", formatVRRi},
	0xe654: {"VUPKZH ", formatVRRk},
	0xe655: {"VNCF  ", formatVRRa},
	0xe656: {"VCLFNH ", formatVRRa},
	0xe658: {"VCVD  ", formatVRIi},
	0xe659: {"VSRP  ", formatVRIg},
	0xe65a: {"VCVDG ", formatVRIi},
	0xe65b: {"VPSOP ", formatVRIg},
	0xe65c: {"VUPKZL ", formatVRRk},
	0xe65d: {"VCFN  ", formatVRRa},
	0xe65e: {"VCLFNL ", formatVRRa},
	0xe65f: {"VTP   ", formatVRRg},
	0xe670: {"VPKZR ", formatVRIf},
	0xe671: {"VAP   ", formatVRIf},
	0xe672: {"VSRPR ", formatVRIf},
	0xe673: {"VSP   ", formatVRIf},
	0xe674: {"VSCHP ", formatVRRb},
	0xe675: {"VCRNF ", formatVRRb},
	0xe677: {"VCP   ", formatVRRh},
	0xe678: {"VMP   ", formatVRIf},
	0xe679: {"VMSP  ", formatVRIf},
	0xe67a: {"VDP   ", formatVRIf},
	0xe67b: {"VRP   ", formatVRIf},
	0xe67c: {"VSCSHP ", formatVRRb},
	0xe67d: {"VCSPH ", formatVRRj},
	0xe67e: {"VSDP  ", formatVRIf},
	0xe67f: {"VTZ   ", formatVRIl},
	0xe700: {"VLEB  ", formatVRX},
	0xe701: {"VLEH  ", formatVRX},
	0xe702: {"VLEG  ", formatVRX},
	0xe703: {"VLEF  ", formatVRX},
	0xe704: {"VLLEZ ", formatVRX},
	0xe705: {"VLREP ", formatVRX},
	0xe706: {"VL    ", formatVRX},
	0xe707: {"VLBB  ", formatVRX},
	0xe708: {"VSTEB ", formatVRX},
	0xe709: {"VSTEH ", formatVRX},
	0xe70a: {"VSTEG ", formatVRX},
	0xe70b: {"VSTEF ", formatVRX},
	0xe70e: {"VST   ", formatVRX},
	0xe712: {"VGEG  ", formatVRV},
	0xe713: {"VGEF  ", formatVRV},
	0xe71a: {"VSCEG ", formatVRV},
	0xe71b: {"VSCEF ", formatVRV},
	0xe721: {"VLGV  ", formatVRSc},
	0xe722: {"VLVG  ", formatVRSb},
	0xe727: {"LCBB  ", formatRXE2},
	0xe730: {"VESL  ", formatVRSa},
	0xe733: {"VERLL ", formatVRSa},
	0xe736: {"VLM   ", formatVRSa},
	0xe737: {"VLL   ", formatVRSb},
	0xe738: {"VESRL ", formatVRSa},
	0xe73a: {"VESRA ", formatVRSa},
	0xe73e: {"VSTM  ", formatVRSa},
	0xe73f: {"VSTL  ", formatVRSb},
	0xe740: {"VLEIB ", formatVRIa},
	0xe741: {"VLEIH ", formatVRIa},
	0xe742: {"VLEIG ", formatVRIa},
	0xe743: {"VLEIF ", formatVRIa},
	0xe744: {"VGBM  ", formatVRIa},
	0xe745: {"VREPI ", formatVRIa},
	0xe746: {"VGM   ", formatVRIb},
	0xe74a: {"VFTCI ", formatVRIe},
	0xe74d: {"VREP  ", formatVRIc},
	0xe750: {"VPOPCT ", formatVRRa},
	0xe752: {"VCTZ  ", formatVRRa},
	0xe753: {"VCLZ  ", formatVRRa},
	0xe754: {"VGEM  ", formatVRRa},
	0xe756: {"VLR   ", formatVRRa},
	0xe75c: {"VISTR ", formatVRRa},
	0xe75f: {"VSEG  ", formatVRRa},
	0xe760: {"VMRL  ", formatVRRc},
	0xe761: {"VMRH  ", formatVRRc},
	0xe762: {"VLVGP ", formatVRRf},
	0xe764: {"VSUM  ", formatVRRc},
	0xe765: {"VSUMG ", formatVRRc},
	0xe766: {"VCKSM ", formatVRRc},
	0xe767: {"VSUMQ ", formatVRRc},
	0xe768: {"VN    ", formatVRRc},
	0xe769: {"VNC   ", formatVRRc},
	0xe76a: {"VO    ", formatVRRc},
	0xe76b: {"VNO   ", formatVRRc},
	0xe76c: {"VNX   ", formatVRRc},
	0xe76d: {"VX    ", formatVRRc},
	0xe76e: {"VNN   ", formatVRRc},
	0xe76f: {"VOC   ", formatVRRc},
	0xe770: {"VESLV ", formatVRRc},
	0xe772: {"VERIM ", formatVRId},
	0xe773: {"VERLLV ", formatVRRc},
	0xe774: {"VSL   ", formatVRRc},
	0xe775: {"VSLB  ", formatVRRc},
	0xe777: {"VSLDB ", formatVRId},
	0xe778: {"VESRLV ", formatVRRc},
	0xe77a: {"VESRAV ", formatVRRc},
	0xe77c: {"VSRL  ", formatVRRc},
	0xe77d: {"VSRLB ", formatVRRc},
	0xe77e: {"VSRA  ", formatVRRc},
	0xe77f: {"VSRAB ", formatVRRc},
	0xe780: {"VFEE  ", formatVRRb},
	0xe781: {"VFENE ", formatVRRb},
	0xe782: {"VFAE  ", formatVRRb},
	0xe784: {"VPDI  ", formatVRRc},
	0xe785: {"VBPERM ", formatVRRc},
	0xe786: {"VSLD  ", formatVRId},
	0xe787: {"VSRD  ", formatVRId},
	0xe788: {"VEVAL ", formatVRIk},
	0xe789: {"VBLEND ", formatVRRd},
	0xe78a: {"VSTRC ", formatVRRd},
	0xe78b: {"VSTRS ", formatVRRd},
	0xe78c: {"VPERM ", formatVRRe},
	0xe78d: {"VSEL  ", formatVRRe},
	0xe78e: {"VFMS  ", formatVRRe},
	0xe78f: {"VFMA  ", formatVRRe},
	0xe794: {"VPK   ", formatVRRc},
	0xe795: {"VPKLS ", formatVRRb},
	0xe797: {"VPKS  ", formatVRRb},
	0xe79e: {"VFNMS ", formatVRRe},
	0xe79f: {"VFNMA ", formatVRRe},
	0xe7a1: {"VMLH  ", formatVRRc},
	0xe7a2: {"VML   ", formatVRRc},
	0xe7a3: {"VMH   ", formatVRRc},
	0xe7a4: {"VMLE  ", formatVRRc},
	0xe7a5: {"VMLO  ", formatVRRc},
	0xe7a6: {"VME   ", formatVRRc},
	0xe7a7: {"VMO   ", formatVRRc},
	0xe7a9: {"VMALH ", formatVRRd},
	0xe7aa: {"VMAL  ", formatVRRd},
	0xe7ab: {"VMAH  ", formatVRRd},
	0xe7ac: {"VMALE ", formatVRRd},
	0xe7ad: {"VMALO ", formatVRRd},
	0xe7ae: {"VMAE  ", formatVRRd},
	0xe7af: {"VMAO  ", formatVRRd},
	0xe7b0: {"VDL   ", formatVRRc},
	0xe7b1: {"VRL   ", formatVRRc},
	0xe7b2: {"VD    ", formatVRRc},
	0xe7b3: {"VR    ", formatVRRc},
	0xe7b4: {"VGFM  ", formatVRRc},
	0xe7b8: {"VMSL  ", formatVRRd},
	0xe7b9: {"VACCC ", formatVRRd},
	0xe7bb: {"VAC   ", formatVRRd},
	0xe7bc: {"VGFMA ", formatVRRd},
	0xe7bd: {"VSBCBI ", formatVRRd},
	0xe7bf: {"VSBI  ", formatVRRd},
	0xe7c0: {"VCLFP ", formatVRRa},
	0xe7c1: {"VCFPL ", formatVRRa},
	0xe7c2: {"VCSFP ", formatVRRa},
	0xe7c3: {"VCFPS ", formatVRRa},
	0xe7c4: {"VFLL  ", formatVRRa},
	0xe7c5: {"VFLR  ", formatVRRa},
	0xe7c7: {"VFI   ", formatVRRa},
	0xe7ca: {"WFK   ", formatVRRa},
	0xe7cb: {"WFC   ", formatVRRa},
	0xe7cc: {"VFPSO ", formatVRRa},
	0xe7ce: {"VFSQ  ", formatVRRa},
	0xe7d4: {"VUPLL ", formatVRRa},
	0xe7d5: {"VUPLH ", formatVRRa},
	0xe7d6: {"VUPL  ", formatVRRa},
	0xe7d7: {"VUPH  ", formatVRRa},
	0xe7d8: {"VTM   ", formatVRRa},
	0xe7d9: {"VECL  ", formatVRRa},
	0xe7db: {"VEC   ", formatVRRa},
	0xe7de: {"VLC   ", formatVRRa},
	0xe7df: {"VLP   ", formatVRRa},
	0xe7e2: {"VFS   ", formatVRRc},
	0xe7e3: {"VFA   ", formatVRRc},
	0xe7e5: {"VFD   ", formatVRRc},
	0xe7e7: {"VFM   ", formatVRRc},
	0xe7e8: {"VFCE  ", formatVRRc},
	0xe7ea: {"VFCHE ", formatVRRc},
	0xe7eb: {"VFCH  ", formatVRRc},
	0xe7ee: {"VFMIN ", formatVRRc},
	0xe7ef: {"VFMAX ", formatVRRc},
	0xe7f0: {"VAVGL ", formatVRRc},
	0xe7f1: {"VACC  ", formatVRRc},
	0xe7f2: {"VAVG  ", formatVRRc},
	0xe7f3: {"VA    ", formatVRRc},
	0xe7f5: {"VSCBI ", formatVRRc},
	0xe7f7: {"VS    ", formatVRRc},
	0xe7f8: {"VCEQ  ", formatVRRb},
	0xe7f9: {"VCHL  ", formatVRRb},
	0xe7fb: {"VCH   ", formatVRRb},
	0xe7fc: {"VMNL  ", formatVRRc},
	0xe7fd: {"VMXL  ", formatVRRc},
	0xe7fe: {"VMN   ", formatVRRc},
	0xe7ff: {"VMX   ", formatVRRc},
	0xe800: {"MVCIN ", formatSS1},
	0xe900: {"PKA   ", formatSS6},
	0xea00: {"UNPKA ", formatSS1},
	0xeb04: {"LMG   ", formatRSY1},
	0xeb0a: {"SRAG  ", formatRSY1},
	0xeb0b: {"SLAG  ", formatRSY1},
	0xeb0c: {"SRLG  ", formatRSY1},
	0xeb0d: {"SLLG  ", formatRSY1},
	0xeb0f: {"TRACG ", formatRSY1},
	0xeb14: {"CSY   ", formatRSY1},
	0xeb1c: {"RLLG  ", formatRSY1},
	0xeb1d: {"RLL   ", formatRSY1},
	0xeb20: {"CLMH  ", formatRSY2},
	0xeb21: {"CLMY  ", formatRSY2},
	0xeb23: {"CLT   ", formatRSY2},
	0xeb24: {"STMG  ", formatRSY1},
	0xeb25: {"STCTG ", formatRSY1},
	0xeb26: {"STMH  ", formatRSY1},
	0xeb2b: {"CLGT  ", formatRSY2},
	0xeb2c: {"STCMH ", formatRSY2},
	0xeb2d: {"STCMY ", formatRSY2},
	0xeb2f: {"LCTLG ", formatRSY1},
	0xeb30: {"CSG   ", formatRSY1},
	0xeb31: {"CDSY  ", formatRSY1},
	0xeb3e: {"CDSG  ", formatRSY1},
	0xeb44: {"BXHG  ", formatRSY1},
	0xeb45: {"BXLEG ", formatRSY1},
	0xeb4c: {"ECAG  ", formatRSY1},
	0xeb51: {"TMY   ", formatSIY},
	0xeb52: {"MVIY  ", formatSIY},
	0xeb54: {"NIY   ", formatSIY},
	0xeb55: {"CLIY  ", formatSIY},
	0xeb56: {"OIY   ", formatSIY},
	0xeb57: {"XIY   ", formatSIY},
	0xeb6a: {"ASI   ", formatSIY},
	0xeb6e: {"ALSI  ", formatSIY},
	0xeb71: {"LPSWEY ", formatSIY},
	0xeb7a: {"AGSI  ", formatSIY},
	0xeb7e: {"ALGSI ", formatSIY},
	0xeb80: {"ICMH  ", formatRSY2},
	0xeb81: {"ICMY  ", formatRSY2},
	0xeb8e: {"MVCLU ", formatRSY1},
	0xeb8f: {"CLCLU ", formatRSY1},
	0xeb90: {"STMY  ", formatRSY1},
	0xeb96: {"LMH   ", formatRSY1},
	0xeb98: {"LMY   ", formatRSY1},
	0xeb9a: {"LAMY  ", formatRSY1},
	0xeb9b: {"STAMY ", formatRSY1},
	0xebc0: {"TP    ", formatRSL},
	0xebdc: {"SRAK  ", formatRSY1},
	0xebdd: {"SLAK  ", formatRSY1},
	0xebde: {"SRLK  ", formatRSY1},
	0xebdf: {"SLLK  ", formatRSY1},
	0xebe0: {"LOCFH ", formatRSY2},
	0xebe1: {"STOCFH ", formatRSY2},
	0xebe2: {"LOCG  ", formatRSY2},
	0xebe3: {"STOCG ", formatRSY2},
	0xebe4: {"LANG  ", formatRSY1},
	0xebe6: {"LAOG  ", formatRSY1},
	0xebe7: {"LAXG  ", formatRSY1},
	0xebe8: {"LAAG  ", formatRSY1},
	0xebea: {"LAALG ", formatRSY1},
	0xebf2: {"LOC   ", formatRSY2},
	0xebf3: {"STOC  ", formatRSY2},
	0xebf4: {"LAN   ", formatRSY1},
	0xebf6: {"LAO   ", formatRSY1},
	0xebf7: {"LAX   ", formatRSY1},
	0xebf8: {"LAA   ", formatRSY1},
	0xebfa: {"LAAL  ", formatRSY1},
	0xec42: {"LOCHI ", formatRIEg},
	0xec44: {"BRXHG ", formatRIE},
	0xec45: {"BRXLG ", formatRIE},
	0xec46: {"LOCGHI ", formatRIEg},
	0xec4e: {"LOCHHI ", formatRIEg},
	0xec51: {"RISBLG ", formatRIEf},
	0xec54: {"RNSBG ", formatRIEf},
	0xec55: {"RISBG ", formatRIEf},
	0xec56: {"ROSBG ", formatRIEf},
	0xec57: {"RXSBG ", formatRIEf},
	0xec59: {"RISBGN ", formatRIEf},
	0xec5d: {"RISBHG ", formatRIEf},
	0xec64: {"CGRJ  ", formatRIEb},
	0xec65: {"CLGRJ ", formatRIEb},
	0xec70: {"CGIT  ", formatRIEa},
	0xec71: {"CLGIT ", formatRIEaU},
	0xec72: {"CIT   ", formatRIEa},
	0xec73: {"CLFIT ", formatRIEaU},
	0xec76: {"CRJ   ", formatRIEb},
	0xec77: {"CLRJ  ", formatRIEb},
	0xec7c: {"CGIJ  ", formatRIEc},
	0xec7d: {"CLGIJ ", formatRIEc},
	0xec7e: {"CIJ   ", formatRIEc},
	0xec7f: {"CLIJ  ", formatRIEc},
	0xecd8: {"AHIK  ", formatRIEd},
	0xecd9: {"AGHIK ", formatRIEd},
	0xecda: {"ALHSIK ", formatRIEd},
	0xecdb: {"ALGHSIK ", formatRIEd},
	0xece4: {"CGRB  ", formatRRS},
	0xece5: {"CLGRB ", formatRRS},
	0xecf6: {"CRB   ", formatRRS},
	0xecf7: {"CLRB  ", formatRRS},
	0xecfc: {"CGIB  ", formatRIS},
	0xecfd: {"CLGIB ", formatRIS},
	0xecfe: {"CIB   ", formatRIS},
	0xecff: {"CLIB  ", formatRIS},
	0xed04: {"LDEB  ", formatRXE},
	0xed05: {"LXDB  ", formatRXE},
	0xed06: {"LXEB  ", formatRXE},
	0xed07: {"MXDB  ", formatRXE},
	0xed08: {"KEB   ", formatRXE},
	0xed09: {"CEB   ", formatRXE},
	0xed0a: {"AEB   ", formatRXE},
	0xed0b: {"SEB   ", formatRXE},
	0xed0c: {"MDEB  ", formatRXE},
	0xed0d: {"DEB   ", formatRXE},
	0xed0e: {"MAEB  ", formatRXF},
	0xed0f: {"MSEB  ", formatRXF},
	0xed10: {"TCEB  ", formatRXE},
	0xed11: {"TCDB  ", formatRXE},
	0xed12: {"TCXB  ", formatRXE},
	0xed14: {"SQEB  ", formatRXE},
	0xed15: {"SQDB  ", formatRXE},
	0xed17: {"MEEB  ", formatRXE},
	0xed18: {"KDB   ", formatRXE},
	0xed19: {"CDB   ", formatRXE},
	0xed1a: {"ADB   ", formatRXE},
	0xed1b: {"SDB   ", formatRXE},
	0xed1c: {"MDB   ", formatRXE},
	0xed1d: {"DDB   ", formatRXE},
	0xed1e: {"MADB  ", formatRXF},
	0xed1f: {"MSDB  ", formatRXF},
	0xed24: {"LDE   ", formatRXE},
	0xed25: {"LXD   ", formatRXE},
	0xed26: {"LXE   ", formatRXE},
	0xed2e: {"MAE   ", formatRXF},
	0xed2f: {"MSE   ", formatRXF},
	0xed34: {"SQE   ", formatRXE},
	0xed35: {"SQD   ", formatRXE},
	0xed37: {"MEE   ", formatRXE},
	0xed38: {"MAYL  ", formatRXF},
	0xed39: {"MYL   ", formatRXF},
	0xed3a: {"MAY   ", formatRXF},
	0xed3b: {"MY    ", formatRXF},
	0xed3c: {"MAYH  ", formatRXF},
	0xed3d: {"MYH   ", formatRXF},
	0xed3e: {"MAD   ", formatRXF},
	0xed3f: {"MSD   ", formatRXF},
	0xed40: {"SLDT  ", formatRXF},
	0xed41: {"SRDT  ", formatRXF},
	0xed48: {"SLXT  ", formatRXF},
	0xed49: {"SRXT  ", formatRXF},
	0xed50: {"TDCET ", formatRXE},
	0xed51: {"TDGET ", formatRXE},
	0xed54: {"TDCDT ", formatRXE},
	0xed55: {"TDGDT ", formatRXE},
	0xed58: {"TDCXT ", formatRXE},
	0xed59: {"TDGXT ", formatRXE},
	0xed64: {"LEY   ", formatRXY},
	0xed65: {"LDY   ", formatRXY},
	0xed66: {"STEY  ", formatRXY},
	0xed67: {"STDY  ", formatRXY},
	0xeda8: {"CZDT  ", formatRSLb},
	0xeda9: {"CZXT  ", formatRSLb},
	0xedaa: {"CDZT  ", formatRSLb},
	0xedab: {"CXZT  ", formatRSLb},
	0xedac: {"CPDT  ", formatRSLb},
	0xedad: {"CPXT  ", formatRSLb},
	0xedae: {"CDPT  ", formatRSLb},
	0xedaf: {"CXPT  ", formatRSLb},
	0xee00: {"PLO   ", formatSS5},
	0xef00: {"LMD   ", formatSS2},
	0xf000: {"SRP   ", formatSS3},
	0xf100: {"MVO   ", formatSS3},
	0xf200: {"PACK  ", formatSS2},
	0xf300: {"UNPK  ", formatSS2},
	0xf800: {"ZAP   ", formatSS2},
	0xf900: {"CP    ", formatSS2},
	0xfa00: {"AP    ", formatSS2},
	0xfb00: {"SP    ", formatSS2},
	0xfc00: {"MP    ", formatSS2},
	0xfd00: {"DP    ", formatSS2},
}

// ── core logic ────────────────────────────────────────────────────────────────

// buildKey constructs the map lookup key from the hex instruction string.
func buildKey(s string, ilen int) int {
	opcode, _ := strconv.ParseInt(s[:4], 16, 64)
	key := int(opcode) & 0xFFFFFF00

	switch ilen {
	case 2:
		if opcode >= 0x0100 && opcode <= 0x01ff {
			key = int(opcode)
		}
	case 4:
		op3, _ := strconv.ParseInt(s[3:4], 16, 64)
		op3 <<= 4
		if (opcode >= 0xa500 && opcode <= 0xa5ff) || (opcode >= 0xa700 && opcode <= 0xa7ff) {
			key |= int(op3)
		}
		if (opcode >= 0xb200 && opcode <= 0xb3ff) || (opcode >= 0xb900 && opcode <= 0xb9ff) {
			key = int(opcode)
		}
	case 6:
		op3, _ := strconv.ParseInt(s[3:4], 16, 64)
		op3 <<= 4
		op4, _ := strconv.ParseInt(s[10:12], 16, 64)
		if (opcode >= 0xc000 && opcode <= 0xc0ff) ||
			(opcode >= 0xc200 && opcode <= 0xc2ff) ||
			(opcode >= 0xc400 && opcode <= 0xc4ff) ||
			(opcode >= 0xc600 && opcode <= 0xc6ff) ||
			(opcode >= 0xc800 && opcode <= 0xccff) {
			key |= int(op3)
		}
		if (opcode >= 0xe300 && opcode <= 0xe3ff) ||
			(opcode >= 0xe600 && opcode <= 0xe7ff) ||
			(opcode >= 0xeb00 && opcode <= 0xedff) {
			key |= int(op4)
		}
		if opcode >= 0xe500 && opcode <= 0xe5ff {
			key = int(opcode)
		}
	}
	return key
}

// Disasm decodes a z/Architecture instruction given as a hex string.
// Returns "MNEM operands" or an error.
func Disasm(s string) (string, error) {
	if !isHex(s) {
		return "", fmt.Errorf("only hexadecimal digits permitted")
	}
	if len(s)%2 != 0 {
		return "", fmt.Errorf("hex string must have even length")
	}

	ilen := len(s) / 2
	ilc, _ := strconv.ParseInt(s[:1], 16, 64)

	switch {
	case ilc >= 0x0 && ilc <= 0x3 && ilen != 2:
		return "", fmt.Errorf("invalid instruction length")
	case ilc >= 0x4 && ilc <= 0xb && ilen != 4:
		return "", fmt.Errorf("invalid instruction length")
	case ilc >= 0xc && ilc <= 0xf && ilen != 6:
		return "", fmt.Errorf("invalid instruction length")
	}

	key := buildKey(s, ilen)
	entry, ok := itable[key]
	if !ok {
		return "", fmt.Errorf("instruction not found")
	}

	operands := entry.fn(s)
	if operands == "" {
		return entry.mnem, nil
	}
	return entry.mnem + operands, nil
}
