package main

import (
	"./controller"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	execPath()
	fmt.Println("path--------------------------------------")
	fmt.Println(execPath())
	fmt.Println(os.Args)
	fmt.Println(os.Environ())

	startServer()
}

func startServer() {
	htmlPath := "./crawler/frontend/view/index.html"
	//解决html中找不到本地css文件（到根目录下找对应的文件目录）
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))

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
