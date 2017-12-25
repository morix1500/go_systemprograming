package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer   = strings.NewReader("COMPUTER")
	system     = strings.NewReader("SYSTEM")
	programing = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	// ここに各書く

	io.Copy(os.Stdout, stream)
}
