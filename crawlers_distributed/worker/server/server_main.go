package main

import (
	"../../rpcsupport"
	"fmt"
	"log"
	"../../worker"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.Worker0),
		worker.CrawlersService{}))
}
