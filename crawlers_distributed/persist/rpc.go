package persist

import (
	"../../crawler/engine"
	"../../crawler/persist"
	"github.com/olivere/elastic"
)

type ItemSavesService struct {
	Client *elastic.Client
	Index  string
}

func (itemServer *ItemSavesService) Save(item engine.Item, result *string) error {
	err := persist.Save(itemServer.Index, itemServer.Client, item)
	if err == nil {
		*result = "OK"
	}
	return err
}
