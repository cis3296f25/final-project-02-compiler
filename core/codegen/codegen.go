package codegen

import (
	"regexp"
)

var addRe = regexp.MustCompile(`IR:\s*tmp0\s*=\s*add\s+(\d+)\s*,\s*(\d+)\s*;\s*ret\s+tmp0`)

// Run returns an assembly-like textual representation for the trivial program using IR operands.
// this isn't AS bad as ir was, but it's still far from the assembly we'll get later
func Run(ir string) string {
	m := addRe.FindStringSubmatch(ir)
	if len(m) != 3 {
		return "ASM: <unrecognized>"
	}
	a := m[1]
	b := m[2]
	return "ASM: section .text; global _main; _main: mov eax, " + a + "; add eax, " + b + "; ret"
}
