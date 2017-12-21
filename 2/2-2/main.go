package main

import (
	"encoding/csv"
	"os"
)

func main() {
	str := []string{"foo", "bar", "hoge", "fuga"}
	writer := csv.NewWriter(os.Stdout)
	writer.Write(str)
	writer.Flush()
}
