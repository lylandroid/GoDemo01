package main

import (
	"../../rpcsupport"
	"../../worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 9000, "worker port")

func main() {
	flag.Parse()
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port) /* config.AppAddress*/,
		worker.CrawlersService{}))
}
