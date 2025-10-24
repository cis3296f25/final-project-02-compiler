package linker

import (
	"fmt"
	"regexp"
	"strings"
)

var textSizeRe = regexp.MustCompile(`(?m)^\s*\.text \(size: (\d+)\)`)
var mainSymRe = regexp.MustCompile(`(?m)^\s*- _main: \.text\+0$`)

// Run parses the object view and emits a mock ELF64 layout with entry and mappings.
func Run(objectFile string) string {
	textSize := parseTextSize(objectFile)
	hasMain := mainSymRe.FindStringSubmatch(objectFile) != nil
	if textSize == 0 || !hasMain {
		return "EXECUTABLE: <unrecognized>"
	}

	// Fake addresses for illustration
	base := uint64(0x400000)
	phoff := uint64(64) // pretend ELF header is 64 bytes
	textVaddr := uint64(0x401000)
	entry := textVaddr // _main at start of .text

	var b strings.Builder
	b.WriteString("EXECUTABLE: a.out\n")
	b.WriteString("ELF64:\n")
	b.WriteString(fmt.Sprintf("  Entry: 0x%X (_main)\n", entry))
	b.WriteString("  ProgramHeaders:\n")
	b.WriteString(fmt.Sprintf("    PHDR[0]: PT_LOAD off=0x%X vaddr=0x%X filesz=%d memsz=%d flags=RX align=0x1000\n", phoff, textVaddr, textSize, textSize))
	b.WriteString("  Sections:\n")
	b.WriteString(fmt.Sprintf("    .text: vaddr=0x%X size=%d\n", textVaddr, textSize))
	b.WriteString("  Symbols:\n")
	b.WriteString("    _start -> _main\n")
	b.WriteString(fmt.Sprintf("  ImageBase: 0x%X\n", base))
	return b.String()
}

func parseTextSize(s string) int {
	m := textSizeRe.FindStringSubmatch(s)
	if len(m) != 2 {
		return 0
	}
	// simple decimal parse
	val := 0
	for i := 0; i < len(m[1]); i++ {
		val = val*10 + int(m[1][i]-'0')
	}
	return val
}
