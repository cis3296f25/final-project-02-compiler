package parser

import (
	"strings"
)

// Run consumes the lexer token string (format: "TOKENS: ...") and returns an ASCII parse tree
// for the specific pattern: return ( INT + INT ); inside a trivial C function
// Note that a lot of this code is devoted to the ascii tree generation and will therefore be removed
func Run(tokens string) string {
	ts := tokenize(tokens)
	a, b, ok := parseReturnAdd(ts)
	if !ok {
		return "PARSE_TREE: <unrecognized>"
	}
	return asciiTree(a, b)
}

// tokenize splits the incoming token string after the "TOKENS: " prefix.
func tokenize(s string) []string {
	s = strings.TrimSpace(s)
	const prefix = "TOKENS: "
	if strings.HasPrefix(s, prefix) {
		s = strings.TrimSpace(s[len(prefix):])
	}
	if s == "" {
		return nil
	}
	return strings.Fields(s)
}

// parseReturnAdd finds a "KEYWORD(return) LPAREN INT(x) PLUS INT(y) RPAREN SEMICOLON" sequence
// and returns x and y if present.
func parseReturnAdd(ts []string) (string, string, bool) {
	// Find the return keyword first
	i := 0
	for i < len(ts) && ts[i] != "KEYWORD(return)" {
		i++
	}
	if i >= len(ts) {
		return "", "", false
	}
	// expected (hardcoded -- this will become dynamic eventually) sequence after return
	need := []string{"LPAREN", "INT(", "PLUS", "INT(", "RPAREN", "SEMICOLON"}
	i++

	if i >= len(ts) || ts[i] != need[0] {
		return "", "", false
	}
	i++

	if i >= len(ts) || !strings.HasPrefix(ts[i], need[1]) || !strings.HasSuffix(ts[i], ")") {
		return "", "", false
	}
	a := strings.TrimSuffix(strings.TrimPrefix(ts[i], "INT("), ")")
	i++

	if i >= len(ts) || ts[i] != need[2] {
		return "", "", false
	}
	i++

	if i >= len(ts) || !strings.HasPrefix(ts[i], need[3]) || !strings.HasSuffix(ts[i], ")") {
		return "", "", false
	}
	b := strings.TrimSuffix(strings.TrimPrefix(ts[i], "INT("), ")")
	i++

	if i >= len(ts) || ts[i] != need[4] {
		return "", "", false
	}
	i++

	if i >= len(ts) || ts[i] != need[5] {
		return "", "", false
	}
	return a, b, true
}

// This is hard-coded for the specific pattern: return ( INT + INT ); inside a trivial C function
// and will really have to go...
func asciiTree(a, b string) string {
	// Build a simple multi-line ASCII tree for: return ( a + b );
	var bld strings.Builder
	bld.WriteString("PARSE_TREE:\n")
	bld.WriteString("ReturnStmt\n")
	bld.WriteString("├─ LParen\n")
	bld.WriteString("├─ Expr\n")
	bld.WriteString("│  ├─ IntLiteral(" + a + ")\n")
	bld.WriteString("│  ├─ Plus\n")
	bld.WriteString("│  └─ IntLiteral(" + b + ")\n")
	bld.WriteString("├─ RParen\n")
	bld.WriteString("└─ Semicolon")
	return bld.String()
}
