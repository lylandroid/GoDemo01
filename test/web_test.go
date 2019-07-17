package main

import (
	"../web/err"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"../web/filelisting"
)

type testUserError string

func (e testUserError) Error() string {
	return e.Error()
}

func (e testUserError) Message() string {
	return string(e)
}

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

var tests = []struct {
	h       filelisting.WebAppHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "User Error"},
	{errNotFound, 404, "Not Found Error"},
	{errPermission, 403, "Forbidden Error"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "Success"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := err.ErrWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet,
			"http://www.baidu.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
		fmt.Println()
		fmt.Println()
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, v := range tests {
		f := err.ErrWrapper(v.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verifyResponse(response, v.code, v.message, t)
		fmt.Println()
		fmt.Println()
	}
}

func verifyResponse(response *http.Response, expendCode int, expendMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expendCode || body != expendMsg {
		t.Errorf("tt.code: %d, tt.message: %s \n"+
			"response.code=%d response.body:%s", expendCode, expendMsg, response.StatusCode, body)
	} else {
		fmt.Println("body: ", body)
	}
}
