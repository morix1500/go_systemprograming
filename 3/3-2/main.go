package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	w, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	r := io.LimitReader(rand.Reader, 1024)

	io.Copy(w, r)
}
