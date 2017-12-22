package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	io.Copy(w, f)
}
