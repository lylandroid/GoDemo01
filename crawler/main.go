package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := string(b)
	fmt.Println(body)

}
