package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type WebAppHandler func(writer http.ResponseWriter, request *http.Request) error

type userError string

func (e userError) Error() string {
	return e.Error()
}

func (e userError) Message() string {
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	fmt.Println("path: ", path)
	if strings.Index(path, prefix) == -1 {
		//return errors.New()
		return userError("path must start with " + prefix)
	}
	path = path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		// panic(err)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	bytesContents, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(bytesContents)
	return nil
}
