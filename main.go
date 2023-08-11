package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("len: first argument should be non-empty")
		os.Exit(1)
	}
	fmt.Println(len(os.Args[1]))
}
