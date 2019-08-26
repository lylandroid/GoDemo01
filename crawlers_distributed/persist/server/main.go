package main

import (
	".."
	"../../config"
	"fmt"
)

func main() {
	persist.StartElasticServer(fmt.Sprintf(":%d", config.AppPort), config.ElasticIndex)
}
