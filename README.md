## C Compiler Demo (Proof of Concept)

Minimal step-by-step visualization of a compilation pipeline for a trivial C program that adds two integers.

### Project Layout

- `cmd/ccdemo`: CLI entrypoint
- `core/`: pipeline logic
- `examples/`: example input files

### Usage

Prerequisites: Go 1.21+ installed.

```bash
go run ./cmd/ccdemo
```

The CLI prints each phase and waits for a key between steps. The input file is `examples/add.c`.

### Example Input

```c
#include <stdio.h>

int main() {
    return (2 + 2);
}
```


