package lexer

import (
	"strings"
	"unicode"
)

type tokenKind string

const (
	tokKeyword tokenKind = "KEYWORD"
	tokIdent   tokenKind = "IDENT"
	tokInt     tokenKind = "INT"
	tokPlus    tokenKind = "PLUS"
	tokLParen  tokenKind = "LPAREN"
	tokRParen  tokenKind = "RPAREN"
	tokLBrace  tokenKind = "LBRACE"
	tokRBrace  tokenKind = "RBRACE"
	tokSemicol tokenKind = "SEMICOLON"
)

type token struct {
	kind  tokenKind
	value string
}

// Run returns a textual representation of tokens for the input.
// This lexer is intentionally minimal and only supports the basic addition example
func Run(source string) string {
	tokens := lex(source)
	return formatTokens(tokens)
}

func lex(input string) []token {
	var tokens []token
	r := []rune(input)
	i := 0
	atLineStart := true

	for i < len(r) {
		ch := r[i]

		// Skip preprocessor lines -- i didn't use any and they're tricky
		if atLineStart && ch == '#' {
			for i < len(r) && r[i] != '\n' {
				i++
			}
			atLineStart = true
			if i < len(r) {
				i++
			}
			continue
		}

		if ch == '\n' {
			atLineStart = true
			i++
			continue
		}

		if unicode.IsSpace(ch) {
			atLineStart = false
			i++
			continue
		}

		atLineStart = false

		// Identifiers and keywords
		if unicode.IsLetter(ch) || ch == '_' {
			start := i
			i++
			for i < len(r) && (unicode.IsLetter(r[i]) || unicode.IsDigit(r[i]) || r[i] == '_') {
				i++
			}
			lit := string(r[start:i])
			switch lit {
			case "int", "return":
				tokens = append(tokens, token{kind: tokKeyword, value: lit})
			default:
				tokens = append(tokens, token{kind: tokIdent, value: lit})
			}
			continue
		}

		// int literals
		if unicode.IsDigit(ch) {
			start := i
			i++
			for i < len(r) && unicode.IsDigit(r[i]) {
				i++
			}
			lit := string(r[start:i])
			tokens = append(tokens, token{kind: tokInt, value: lit})
			continue
		}

		// Single char tokens
		switch ch {
		case '+':
			tokens = append(tokens, token{kind: tokPlus})
			i++
			continue
		case '(':
			tokens = append(tokens, token{kind: tokLParen})
			i++
			continue
		case ')':
			tokens = append(tokens, token{kind: tokRParen})
			i++
			continue
		case '{':
			tokens = append(tokens, token{kind: tokLBrace})
			i++
			continue
		case '}':
			tokens = append(tokens, token{kind: tokRBrace})
			i++
			continue
		case ';':
			tokens = append(tokens, token{kind: tokSemicol})
			i++
			continue
		default:
			// Ignore other characters
			i++
			continue
		}
	}
	return tokens
}

func formatTokens(tokens []token) string {
	if len(tokens) == 0 {
		return "TOKENS: <empty>"
	}
	parts := make([]string, 0, len(tokens))
	for _, t := range tokens {
		switch t.kind {
		case tokKeyword:
			parts = append(parts, "KEYWORD("+t.value+")")
		case tokIdent:
			parts = append(parts, "IDENT("+t.value+")")
		case tokInt:
			parts = append(parts, "INT("+t.value+")")
		case tokPlus:
			parts = append(parts, "PLUS")
		case tokLParen:
			parts = append(parts, "LPAREN")
		case tokRParen:
			parts = append(parts, "RPAREN")
		case tokLBrace:
			parts = append(parts, "LBRACE")
		case tokRBrace:
			parts = append(parts, "RBRACE")
		case tokSemicol:
			parts = append(parts, "SEMICOLON")
		default:
			parts = append(parts, string(t.kind))
		}
	}
	return "TOKENS: " + strings.Join(parts, " ")
}
