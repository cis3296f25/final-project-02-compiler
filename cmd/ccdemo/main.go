package main

// Purpose: CLI entrypoint for the compilation-steps demonstration.

import (
	"fmt"
	"os"
	"path/filepath"

	"c_compiler_demo/core"
)

func main() {
	examplePath := filepath.Join("examples", "add.c")
	content, err := os.ReadFile(examplePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", examplePath, err)
		os.Exit(1)
	}

	phases := core.BuildPhases(string(content))
	if err := core.RunPhases(phases); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
