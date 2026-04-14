package main

import (
	"fmt"
	"os"
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
	raw, _ := strconv.ParseUint(s[4:12], 16, 32)
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

func formatRXF(s string) string {
	r3, _ := strconv.ParseUint(s[2:3], 16, 8)
	x2, _ := strconv.ParseUint(s[3:4], 16, 8)
	b2, _ := strconv.ParseUint(s[4:5], 16, 8)
	d2, _ := strconv.ParseUint(s[5:8], 16, 16)
	r1, _ := strconv.ParseUint(s[8:9], 16, 8)
	return fmt.Sprintf("R%d,R%d,%d(R%d,R%d)", r1, r3, d2, x2, b2)
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

func formatSIY(s string) string {
	i2, _ := strconv.ParseUint(s[2:4], 16, 8)
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	return fmt.Sprintf("%d(R%d),%d", disp20(s), b1, i2)
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

func formatSSE(s string) string {
	b1, _ := strconv.ParseUint(s[4:5], 16, 8)
	d1, _ := strconv.ParseUint(s[5:8], 16, 16)
	b2, _ := strconv.ParseUint(s[8:9], 16, 8)
	d2, _ := strconv.ParseUint(s[9:12], 16, 16)
	return fmt.Sprintf("%d(R%d),%d(R%d)", d1, b1, d2, b2)
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
	0x7700: {"PBL   ", formatRX1},
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
	0xb221: {"IPTE  ", formatRRE1},
	0xb222: {"IPM   ", formatRRE2},
	0xb223: {"IVSK  ", formatRRE1},
	0xb224: {"IAC   ", formatRRE2},
	0xb225: {"SSAR  ", formatRRE2},
	0xb226: {"EPAR  ", formatRRE2},
	0xb227: {"ESAR  ", formatRRE2},
	0xb228: {"PT    ", formatRRE1},
	0xb229: {"ISKE  ", formatRRE1},
	0xb22a: {"RRBE  ", formatRRE1},
	0xb22b: {"SSKE  ", formatRRE1},
	0xb22c: {"TB    ", formatRRE1},
	0xb22d: {"DXR   ", formatRRE1},
	0xb22e: {"PGIN  ", formatRRE1},
	0xb22f: {"PGOUT ", formatRRE1},
	0xb230: {"CSCH  ", formatS},
	0xb231: {"HSCH  ", formatS},
	0xb232: {"MSCH  ", formatS},
	0xb233: {"SSCH  ", formatS},
	0xb234: {"STSCH ", formatS},
	0xb235: {"TSCH  ", formatS},
	0xb236: {"TPI   ", formatS},
	0xb237: {"SAL   ", formatS},
	0xb238: {"RSCH  ", formatS},
	0xb239: {"STCRW ", formatS},
	0xb23a: {"STPCS ", formatS},
	0xb23b: {"RCHP  ", formatS},
	0xb23c: {"SCHM  ", formatS},
	0xb240: {"BAKR  ", formatRRE1},
	0xb241: {"CKSM  ", formatRRE1},
	0xb244: {"SQDR  ", formatRRE1},
	0xb245: {"SQER  ", formatRRE1},
	0xb246: {"STURA ", formatRRE1},
	0xb247: {"MSTA  ", formatRRE2},
	0xb248: {"PALB  ", formatRRE1},
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
	0xb276: {"XSCH  ", formatS},
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
	0xb2f8: {"TEND  ", formatS},
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
	0xb3d0: {"MDTR  ", formatRRF7},
	0xb3d8: {"MXTR  ", formatRRF7},
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
	0xb92e: {"KM    ", formatRRE1},
	0xb92f: {"KMC   ", formatRRE1},
	0xb930: {"CGFR  ", formatRRE1},
	0xb931: {"CLGFR ", formatRRE1},
	0xb93e: {"KIMD  ", formatRRE1},
	0xb93f: {"KLMD  ", formatRRE1},
	0xb946: {"BCTGR ", formatRRE1},
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
	0xb98d: {"EPSW  ", formatRRE1},
	0xb98e: {"IDTE  ", formatRRF1},
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
	0xb9aa: {"LPTEA ", formatRRF3},
	0xb9b0: {"CU14  ", formatRRE1},
	0xb9b1: {"CU24  ", formatRRE1},
	0xb9b2: {"CU41  ", formatRRE1},
	0xb9b3: {"CU42  ", formatRRE1},
	0xb9be: {"SRSTU ", formatRRE1},
	0xba00: {"CS    ", formatRS1},
	0xbb00: {"CDS   ", formatRS1},
	0xbd00: {"CLM   ", formatRS2},
	0xbe00: {"STCM  ", formatRS2},
	0xbf00: {"ICM   ", formatRS2},
	0xc000: {"LARL  ", formatRIL1},
	0xc010: {"LGFI  ", formatRIL1},
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
	0xc240: {"SLGFI ", formatRIL1},
	0xc250: {"SLFI  ", formatRIL1},
	0xc280: {"AGFI  ", formatRIL1},
	0xc290: {"AFI   ", formatRIL1},
	0xc2a0: {"ALGFI ", formatRIL1},
	0xc2b0: {"ALFI  ", formatRIL1},
	0xc2c0: {"CGFI  ", formatRIL1},
	0xc2d0: {"CFI   ", formatRIL1},
	0xc2e0: {"CLGFI ", formatRIL1},
	0xc2f0: {"CLFI  ", formatRIL1},
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
	0xe100: {"PKU   ", formatSS1},
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
	0xe326: {"CVDY  ", formatRXY},
	0xe32e: {"CVDG  ", formatRXY},
	0xe32f: {"STRVG ", formatRXY},
	0xe330: {"CGF   ", formatRXY},
	0xe331: {"CLGF  ", formatRXY},
	0xe33e: {"STRV  ", formatRXY},
	0xe33f: {"STRVH ", formatRXY},
	0xe346: {"BCTG  ", formatRXY},
	0xe350: {"STY   ", formatRXY},
	0xe351: {"MSY   ", formatRXY},
	0xe354: {"NY    ", formatRXY},
	0xe355: {"CLY   ", formatRXY},
	0xe356: {"OY    ", formatRXY},
	0xe357: {"XY    ", formatRXY},
	0xe358: {"LY    ", formatRXY},
	0xe359: {"CY    ", formatRXY},
	0xe35a: {"AY    ", formatRXY},
	0xe35b: {"SY    ", formatRXY},
	0xe35e: {"ALY   ", formatRXY},
	0xe35f: {"SLY   ", formatRXY},
	0xe370: {"STHY  ", formatRXY},
	0xe371: {"LAY   ", formatRXY},
	0xe372: {"STCY  ", formatRXY},
	0xe373: {"ICY   ", formatRXY},
	0xe376: {"LB    ", formatRXY},
	0xe377: {"LGB   ", formatRXY},
	0xe378: {"LHY   ", formatRXY},
	0xe379: {"CHY   ", formatRXY},
	0xe37a: {"AHY   ", formatRXY},
	0xe37b: {"SHY   ", formatRXY},
	0xe380: {"NG    ", formatRXY},
	0xe381: {"OG    ", formatRXY},
	0xe382: {"XG    ", formatRXY},
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
	0xe500: {"LASP  ", formatSSE},
	0xe501: {"TPROT ", formatSSE},
	0xe502: {"STRAG ", formatSSE},
	0xe50e: {"MVCSK ", formatSSE},
	0xe50f: {"MVCDK ", formatSSE},
	0xe800: {"MVCIN ", formatSS1},
	0xe900: {"PKA   ", formatSS1},
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
	0xeb24: {"STMG  ", formatRSY1},
	0xeb25: {"STCTG ", formatRSY1},
	0xeb26: {"STMH  ", formatRSY1},
	0xeb2c: {"STCMH ", formatRSY2},
	0xeb2d: {"STCMY ", formatRSY2},
	0xeb2f: {"LCTLG ", formatRSY1},
	0xeb30: {"CSG   ", formatRSY1},
	0xeb31: {"CDSY  ", formatRSY1},
	0xeb3e: {"CDSG  ", formatRSY1},
	0xeb44: {"BXHG  ", formatRSY1},
	0xeb45: {"BXLEG ", formatRSY1},
	0xeb51: {"TMY   ", formatSIY},
	0xeb52: {"MVIY  ", formatSIY},
	0xeb54: {"NIY   ", formatSIY},
	0xeb55: {"CLIY  ", formatSIY},
	0xeb56: {"OIY   ", formatSIY},
	0xeb57: {"XIY   ", formatSIY},
	0xeb80: {"ICMH  ", formatRSY1},
	0xeb81: {"ICMY  ", formatRSY1},
	0xeb8e: {"MVCLU ", formatRSY1},
	0xeb8f: {"CLCLU ", formatRSY1},
	0xeb90: {"STMY  ", formatRSY1},
	0xeb96: {"LMH   ", formatRSY1},
	0xeb98: {"LMY   ", formatRSY1},
	0xeb9a: {"LAMY  ", formatRSY1},
	0xeb9b: {"STAMY ", formatRSY1},
	0xebc0: {"TP    ", formatRSL},
	0xec44: {"BRXHG ", formatRIE},
	0xec45: {"BRXLG ", formatRIE},
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
	0xed64: {"LEY   ", formatRXY},
	0xed65: {"LDY   ", formatRXY},
	0xed66: {"STEY  ", formatRXY},
	0xed67: {"STDY  ", formatRXY},
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
		if (opcode >= 0xc000 && opcode <= 0xc0ff) || (opcode >= 0xc200 && opcode <= 0xc2ff) {
			key |= int(op3)
		}
		if (opcode >= 0xe300 && opcode <= 0xe3ff) || (opcode >= 0xeb00 && opcode <= 0xedff) {
			key |= int(op4)
		}
		if opcode >= 0xe500 && opcode <= 0xe5ff {
			key = int(opcode)
		}
	}
	return key
}

// disasm decodes a z/Architecture instruction given as a hex string.
// Returns "MNEM  operands" or an error.
func disasm(s string) (string, error) {
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

// ── main ──────────────────────────────────────────────────────────────────────

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: zdisasm <hex-instruction>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		os.Exit(1)
	}
	result, err := disasm(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}
