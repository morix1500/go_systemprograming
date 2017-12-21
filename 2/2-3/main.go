package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	writer := gzip.NewWriter(w)
	writer.Header.Name = "test.gz"
	encoder := json.NewEncoder(io.MultiWriter(writer, os.Stdout))
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
	writer.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
