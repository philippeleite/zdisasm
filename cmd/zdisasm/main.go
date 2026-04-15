package main

import (
	"fmt"
	"os"

	"github.com/philippeleite/zdisasm"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: zdisasm <hex-instruction>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		os.Exit(1)
	}
	result, err := zdisasm.Disasm(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}
