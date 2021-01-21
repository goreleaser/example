package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Println("hello world, gitlab! (", version, runtime.GOOS, runtime.GOARCH, ")")
}
