package main

import (
	"../../config"
	"../../rpcsupport"
	"../../worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc( /*fmt.Sprintf(":%d", config.Worker0) */ config.AppAddress,
		worker.CrawlersService{}))
}
