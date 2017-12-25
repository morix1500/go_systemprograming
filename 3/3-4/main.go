package main

import (
	"archive/zip"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Dispostion", "attachment; filename=hoge.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	file, err := zipWriter.Create("hoge.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("hogeeeeeeeeeeee"))
	zipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
