package controller

import (
	"fmt"
	"github.com/olivere/elastic"
	"net/http"
	"strconv"
	"strings"
	"../view"
	"../../persist"
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
	fmt.Fprintf(w, "q=%s,from=%d", q, from)
	//page := getSea
	//16-10 12分钟
}
