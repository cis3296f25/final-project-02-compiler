package ast

import (
	"regexp"
	"strings"
)

var intLitRe = regexp.MustCompile(`IntLiteral\((\d+)\)`) // matches IntLiteral(123) -- will be dynamic eventually

// Run parses the parser's ASCII tree to build a tiny AST for (INT + INT).
// At this stage of compilation, we're working with lexed string literals and a parsed ascii tree
// which is not at all like the AST we'll get later
// as the stages progress, the difference between this and a real compiler will grow
func Run(parseTree string) string {
	a, b, ok := extractTwoInts(parseTree)
	if !ok {
		return "AST: <unrecognized>"
	}
	return asciiAST(a, b)
}

func extractTwoInts(s string) (string, string, bool) {
	matches := intLitRe.FindAllStringSubmatch(s, -1)
	if len(matches) < 2 {
		return "", "", false
	}
	return matches[0][1], matches[1][1], true
}

func asciiAST(a, b string) string {
	var bld strings.Builder
	bld.WriteString("AST:\n")
	bld.WriteString("BinaryExpr(Add)\n")
	bld.WriteString("├─ IntLiteral(" + a + ")\n")
	bld.WriteString("└─ IntLiteral(" + b + ")")
	return bld.String()
}
