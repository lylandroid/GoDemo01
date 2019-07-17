package main

import (
	"./filelisting"
	"net/http"
	"./err"
)

func main() {
	http.HandleFunc("/", err.ErrWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
