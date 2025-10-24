package core

// Purpose: Minimal pipeline coordinator to run each compilation phase and print results.

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"c_compiler_demo/core/asm"
	"c_compiler_demo/core/ast"
	"c_compiler_demo/core/codegen"
	"c_compiler_demo/core/ir"
	"c_compiler_demo/core/lexer"
	"c_compiler_demo/core/linker"
	"c_compiler_demo/core/parser"
	"c_compiler_demo/core/sema"
)

// Phase represents a single compilation phase with a label and a result string.
type Phase struct {
	Label  string
	Result string
}

// BuildPhases constructs the ordered list of phases for the demo using stubs.
func BuildPhases(source string) []Phase {
	tokens := lexer.Run(source)
	parseTree := parser.Run(tokens)
	astText := ast.Run(parseTree)
	semaText := sema.Run(astText)
	irText := ir.Run(semaText)
	assembly := codegen.Run(irText)
	objectFile := asm.Run(assembly)
	executable := linker.Run(objectFile)

	return []Phase{
		{Label: "the file contents is", Result: source},
		{Label: "after lexing, the result is", Result: tokens},
		{Label: "after parsing, the result is", Result: parseTree},
		{Label: "after AST, the result is", Result: astText},
		{Label: "after semantic analysis, the result is", Result: semaText},
		{Label: "after ir, the result is", Result: irText},
		{Label: "after codegen (assembly), the result is", Result: assembly},
		{Label: "after assembling, the result is", Result: objectFile},
		{Label: "after linking, the result is", Result: executable},
	}
}

// RunPhases prints each phase and waits for a single keypress between phases.
func RunPhases(phases []Phase) error {
	reader := bufio.NewReader(os.Stdin)
	for i, p := range phases {
		// Colored header
		header := fmt.Sprintf("%s%d.%s %s%s%s", colorBoldCyan, i+1, colorReset, colorBold, p.Label, colorReset)
		fmt.Println(header)
		// Indented result body
		fmt.Println(indentLines(p.Result, "    "))
		if i < len(phases)-1 {
			fmt.Println(colorDim + "press any key to view the next step" + colorReset)
			_, _ = reader.ReadByte()
			for reader.Buffered() > 0 {
				_, _ = reader.ReadByte()
			}
			fmt.Println()
		} else {
			fmt.Println(colorBoldGreen + "this is the final step" + colorReset)
		}
	}
	return nil
}

// Formatting helpers
const (
	colorReset     = "\033[0m"
	colorBold      = "\033[1m"
	colorBoldCyan  = "\033[1;36m"
	colorBoldGreen = "\033[1;32m"
	colorDim       = "\033[90m"
)

func indentLines(s string, prefix string) string {
	if s == "" {
		return ""
	}
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = prefix + lines[i]
	}
	return strings.Join(lines, "\n")
}
