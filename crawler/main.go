package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	//body := string(b)
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	result := compile.FindAllSubmatch(b, -1)
	fmt.Println(len(result))
	for _, v := range result {
		fmt.Printf("%s \t %s \n", v[1], v[2])
	}

}
