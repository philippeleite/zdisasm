# zdisasm

A small, dependency-free disassembler for IBM z/Architecture (s390x / mainframe) machine instructions, written in Go.

Give it the hex of one instruction — for example from a dump, a trace, or a listing — and it prints the assembler mnemonic and operands in HLASM style:

```sh
$ zdisasm 4110F010
LA    R1,16(R0,R15)
```

Highlights:

- Covers about 1,200 opcodes, including the vector facility instructions.
- Handles every instruction length: 2, 4, or 6 bytes.
- Works as a command-line tool or as a Go library (`zdisasm.Disasm`).
- Uses only the Go standard library.

## Installation

### CLI

Download a prebuilt binary for Linux, Windows, or macOS from the [Releases](https://github.com/philippeleite/zdisasm/releases) page, or install with Go:

```sh
go install github.com/philippeleite/zdisasm/cmd/zdisasm@latest
```

Or build from source:

```sh
git clone https://github.com/philippeleite/zdisasm.git
cd zdisasm
go build ./cmd/zdisasm
```

### Library

```sh
go get github.com/philippeleite/zdisasm
```

## CLI Usage

```sh
zdisasm <hex-instruction>
```

The hex string must have 4, 8, or 12 hex digits (1, 2, or 3 halfwords), representing a valid z/Architecture instruction.

### Examples

```sh
$ zdisasm 1A34
AR    R3,R4

$ zdisasm 4110F010
LA    R1,16(R0,R15)

$ zdisasm A7F4FFEC
BRC   15,-20

$ zdisasm C0E5FFFFFFEC
BRASL R14,-20

$ zdisasm E31F10080004
LG    R1,8(R15,R1)
```

Invalid input returns an error and exits with code 1:

```sh
$ zdisasm ZZZZ
only hexadecimal digits permitted

$ zdisasm 0800
instruction not found

$ zdisasm 03
invalid instruction length
```

## Library Usage

The `zdisasm` package exports a single function:

```go
func Disasm(hex string) (string, error)
```

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/philippeleite/zdisasm"
)

func main() {
	result, err := zdisasm.Disasm("1A34")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) // AR    R3,R4
}
```

## Running Tests

```sh
go test
```

## License

See [LICENSE](LICENSE) for details.
