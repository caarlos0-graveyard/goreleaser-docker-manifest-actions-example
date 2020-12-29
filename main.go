package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Println("example", version, runtime.GOOS, runtime.GOARCH)
}
