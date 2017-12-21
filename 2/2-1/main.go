package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "int: %d string: %s float: %f", 1, "hoge", 2.1)
}
