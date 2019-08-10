package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"./controller"
)

func main() {
	execPath()
	fmt.Println("path--------------------------------------")
	fmt.Println(execPath())
	fmt.Println(os.Args)
	fmt.Println(os.Environ())
	startServer()
}

func startServer()  {
	htmlPath := "E:/project_workspace/idea/go/demo1/crawler/frontend/view/index.html"
	http.Handle("/search", controller.CreateSearchResultHandler(htmlPath))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)

}
