package asm

import (
	"fmt"
	"regexp"
	"strings"
)

var asmRe = regexp.MustCompile(`ASM:.*_main:\s*mov\s+eax,\s*(\d+)\s*;\s*add\s+eax,\s*(\d+)\s*;\s*ret`)

// Run parses the assembly and emits an object view with text bytes and symbols.
func Run(assembly string) string {
	m := asmRe.FindStringSubmatch(assembly)
	if len(m) != 3 {
		return "OBJECT: <unrecognized>"
	}
	a := m[1]
	b := m[2]

	// Fake x86-like encodings for illustration:
	// mov eax, imm32 => B8 imm32-le
	// add eax, imm32 => 05 imm32-le
	// ret => C3
	bytes := []string{
		// B8 <imm32 a>
		"B8", le32(a),
		// 05 <imm32 b>
		"05", le32(b),
		// C3
		"C3",
	}

	var bld strings.Builder
	bld.WriteString("OBJECT: add.o\n")
	bld.WriteString("Symbols:\n")
	bld.WriteString("  - _main: .text+0\n")
	bld.WriteString("Sections:\n")
	bld.WriteString("  .text (size: ")
	bld.WriteString(fmt.Sprintf("%d", encodedLen(bytes)))
	bld.WriteString(")\n")
	bld.WriteString("    Bytes: ")
	bld.WriteString(strings.Join(bytes, " "))
	return bld.String()
}

// le32 converts a small decimal string into a little-endian 32-bit hex representation.
// For simplicity we only support small non-negative ints that fit in 1 byte and zero-extend
// meaning this part won't just break on non-int literals,
// it will also break on large int literals
func le32(dec string) string {
	// naive decimal to byte
	val := 0
	for i := 0; i < len(dec); i++ {
		val = val*10 + int(dec[i]-'0')
	}
	if val < 0 {
		val = 0
	}
	if val > 255 {
		val = val % 256
	}
	// little-endian 32-bit: xx 00 00 00
	return fmt.Sprintf("%02X 00 00 00", val)
}

func encodedLen(parts []string) int {
	count := 0
	for _, p := range parts {
		// p may include spaces, split
		fields := strings.Fields(p)
		count += len(fields)
	}
	return count
}
