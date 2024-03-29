package controller

import (
	"../../engine"
	"../../persist"
	"../model"
	"../view"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := persist.NewElasticClient()
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
	q = reWriteQueryString(q)
	fmt.Println("re...: ", q)
	termQuery := elastic.NewQueryStringQuery(q)
	var result model.SearchResult
	result.Query = q
	searchResult, err := h.client.Search(index).
		//Index(index). // search in index "twitter"
		Type("zhenai").
		Query(termQuery). // specify the query
		//Sort("user", true). // sort by "user" field, ascending
		From(from).Size(10). // take documents 0-9
		//Pretty(true). // pretty print request and response JSON
		//Aggregation("rest_total_hits_as_int",true).
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
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func reWriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1")
}
