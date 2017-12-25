package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	w, err := os.Create("test.gz")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}
	r := strings.NewReader("hoge")

	io.Copy(writer, r)
	zipWriter.Flush()
}
