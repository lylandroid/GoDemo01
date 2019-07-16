package main

import (
	"../mock"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, param map[string]string) string
}

type RetrieverPost interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func session(rp RetrieverPost) string {
	rp.Post(url, map[string]string{"contents": "session post map param"})
	return rp.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{"retriever imp get"}
	r2 := &mock.Retriever{"retriever imp get"}
	fmt.Println(download(r))
	//var r2 RetrieverPost
	//r2 = &mock.Retriever{"retrieverPost"}
	fmt.Println(session(r2))
	/*r = real.Retriever{}
	fmt.Println(download(r))*/

}
