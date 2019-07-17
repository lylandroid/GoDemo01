package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"../web/filelisting"
	"../web/err"
	"strings"
	"fmt"
)

func errorPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func TestErrWrapper2(t *testing.T) {
	tests := []struct {
		h       filelisting.WebAppHandler
		code    int
		message string
	}{
		{errorPanic, 500, ""},
	}

	for _, tt := range tests {
		f := err.ErrWrapper(tt.h)
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		response := httptest.NewRecorder()
		f(response, request)

		verifyResponse2(response.Result(), tt.code, tt.message, t)
	}
}

func verifyResponse2(response *http.Response, expendCode int, expendMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expendCode || body != expendMsg {
		t.Errorf("tt.code: %d, tt.message: %s \n"+
			"response.code=%d response.body:%s", expendCode, expendMsg, response.StatusCode, body)
	} else {
		fmt.Println("body: ", body)
	}
}
