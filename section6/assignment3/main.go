package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fname := os.Args[1]
	file, err := os.Open(fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
