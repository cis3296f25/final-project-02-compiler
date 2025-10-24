package ir

import (
	"regexp"
)

var intLitRe = regexp.MustCompile(`IntLiteral\((\d+)\)\s*:\s*Int`)

// Run parses the SEMA ASCII and emits a small IR program using the actual ints.
// at this stage, because of the several layers of ascii deep we now are, there's nearly no similarity to the IR we'll get later
func Run(sema string) string {
	ints := intLitRe.FindAllStringSubmatch(sema, -1)
	if len(ints) < 2 {
		return "IR: <unrecognized>"
	}
	a := ints[0][1]
	b := ints[1][1]
	return "IR: tmp0 = add " + a + ", " + b + "; ret tmp0"
}
