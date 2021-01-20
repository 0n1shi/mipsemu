package main

import (
	"fmt"

	"github.com/0n1shi/mipsemu/pkg/mips"
)

func main() {
	fmt.Println("hello world")
	_, _ = mips.NewEmulator([]byte{})
}
