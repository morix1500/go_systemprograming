package main

import (
	"io"
	"os"
	"strings"
)

// CopyN implements io.CopyN by myself
func CopyN(dest io.Writer, src io.Reader, length int) {
	r := io.LimitReader(src, int64(length))
	io.Copy(dest, r)
}

func main() {
	src := strings.NewReader("123456789")
	CopyN(os.Stdout, src, 2)
}
