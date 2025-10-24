package sema

import (
	"regexp"
	"strings"
)

var intLitRe = regexp.MustCompile(`IntLiteral\((\d+)\)`) // matches IntLiteral(123) (again, will be dynamic eventually)

// Run consumes the ASCII AST and emits a simple typed tree.
// again, the sema step wouldn't normall get a parsed ascii tree, but it's what we have
func Run(ast string) string {
	a, b, ok := extractTwoInts(ast)
	if !ok {
		return "SEMA: <unrecognized>"
	}
	return asciiTyped(a, b)
}

func extractTwoInts(s string) (string, string, bool) {
	matches := intLitRe.FindAllStringSubmatch(s, -1)
	if len(matches) < 2 {
		return "", "", false
	}
	return matches[0][1], matches[1][1], true
}

// This is aggressively hardcoded for the specific pattern: return ( INT + INT ); inside a trivial C function
func asciiTyped(a, b string) string {
	var bld strings.Builder
	bld.WriteString("SEMA:\n")
	bld.WriteString("TypedBinaryExpr(Add : Int)\n")
	bld.WriteString("├─ IntLiteral(" + a + ") : Int\n")
	bld.WriteString("└─ IntLiteral(" + b + ") : Int")
	return bld.String()
}
