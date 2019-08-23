package main

import (
	"fmt"
	"regexp"
)

const text = `my is email jsdlfjlds@163.com@aaa
my is email zhagnsan@163.com@aaa
my is email lisi@163.com@aaa
my is email wangwu@163.com.cn

`

func main() {
	compile := regexp.MustCompile(`([a-zA-z0-9]+)@([a-zA-z0-9.]+)(\.[a-zA-z]+)`)
	s := compile.FindAllStringSubmatch(text, -1)
	fmt.Println(s)

}
