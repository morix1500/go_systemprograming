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

	strA := io.NewSectionReader(programing, 5, 1)
	strS := io.NewSectionReader(system, 0, 1)
	strC := io.NewSectionReader(computer, 0, 1)
	strI := io.NewSectionReader(programing, 8, 1)
	strI2 := io.NewSectionReader(programing, 8, 1)
	stream = io.MultiReader(strA, strS, strC, strI, strI2)

	io.Copy(os.Stdout, stream)
}
