package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

type Poster struct {
}

func (r Retriever) Get(url string) string {
	fmt.Println("-------------get start----------------------")
	resp, err := http.Get(url)
	fmt.Println("-------------get http end----------------------")
	if err != nil {
		panic(err)
	}
	resultBytes, err := httputil.DumpResponse(resp, true)
	fmt.Println("-------------get read http end----------------------")
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("-------------get return end----------------------")
	return string(resultBytes)

}
