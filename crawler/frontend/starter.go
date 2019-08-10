package main

import (
	"fmt"
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

	/*http.Handle("/search", controller.CreateSearchResultHandler("../demo1/crawler/frontend/view/index.html/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}*/
}

func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)

}
