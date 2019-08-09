package main

import (
	"net/http"
	"./controller"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("view/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
