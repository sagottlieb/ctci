package main

import (
	"fmt"
)

func main() {
	fmt.Println("Run go test and go test -bench=.")
}

type checker interface {
	Check(string) bool
	GetName() string
}
