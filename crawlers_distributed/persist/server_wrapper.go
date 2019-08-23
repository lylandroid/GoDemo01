package persist

import (
	"../../crawler/persist"
	"../rpcsupport"
)

func StartElasticServer(host string, index string) {
	client, err := persist.NewElasticClient()
	if err != nil {
		panic(err)
	}
	err = rpcsupport.ServeRpc(host, &ItemSavesService{
		Client: client,
		Index:  index,
	})
	if err != nil {
		panic(err)
	}
}
