package controller

import (
	"../../persist"
	"../view"
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"../model"
	"../../engine"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := persist.NewClient()
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}

}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s,from=%d", q, from)
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		err = h.view.Reader(w, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

}

var index = "dating_profile"

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	termQuery := elastic.NewQueryStringQuery(q)
	var result model.SearchResult
	searchResult, err := h.client.Search().
		Index(index). // search in index "twitter"
		Query(termQuery). // specify the query
		//Sort("user", true). // sort by "user" field, ascending
		From(from).Size(10). // take documents 0-9
		//Pretty(true).       // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		return result, err
	}
	result.Hits = searchResult.TotalHits()
	result.Start = from
	result.Items = searchResult.Each(reflect.TypeOf(engine.Item{}))
	/*for _, item := range searchResult.Each(reflect.TypeOf(engine.Item{})) {
		if t, ok := item.(engine.Item); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}*/

	return result, nil

}
