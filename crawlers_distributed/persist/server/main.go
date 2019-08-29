package main

import (
	".."
	"../../config"
	"flag"
	"fmt"
)

var port = flag.Int("port", 1234, "port 端口号")

//go run main.go --port=1234 (运行脚本)
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("error must specify a port")
		return
	}
	persist.StartElasticServer(fmt.Sprintf(":%d", *port), config.ElasticIndex)
}
