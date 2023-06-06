package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	fileName    = "ファイル の 名前.txt"
	fileContent = "ファイル の 中身"
	port        = ":8111"
)

func main() {
	// 両方指定
	http.HandleFunc("/file1", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;
filename=%q;
filename*=utf-8''%s`, fileName, url.PathEscape(fileName)))
		io.Copy(w, strings.NewReader(fileContent))
	})
	// filenameだけ指定
	http.HandleFunc("/file2", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;
filename=%q`, fileName))
		io.Copy(w, strings.NewReader(fileContent))
	})
	// filename*=だけ指定
	http.HandleFunc("/file3", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;
filename*=utf-8''%s`, url.PathEscape(fileName)))
		io.Copy(w, strings.NewReader(fileContent))
	})
	fmt.Printf("Listening: %s\n", port)
	http.ListenAndServe(port, nil)
}
