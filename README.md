# zdisasm

Disassembler for IBM z/Architecture (s390x) machine instructions, written in Go.

Takes a hexadecimal string representing a z/Architecture instruction and decodes it into its mnemonic and operands.

## Installation

### CLI

Build from source:

```sh
git clone https://github.com/philippeleite/zdisasm.git
cd zdisasm
go build ./cmd/zdisasm
```

### Library

```sh
go get zdisasm
```

## CLI Usage

```sh
zdisasm <hex-instruction>
```

The hex string must have 2, 4, or 6 hex digits (1, 2, or 3 halfwords), representing a valid z/Architecture instruction.

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
