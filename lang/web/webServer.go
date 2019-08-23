package main

import (
	"./err"
	"./filelisting"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", err.ErrWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
