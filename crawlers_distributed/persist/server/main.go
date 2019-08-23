package main

import (
	".."
)

func main() {
	persist.StartElasticServer(":1234", "dating_profile")
}
