package view

import (
	"../model"
	"html/template"
	"os"
	"testing"

)

func TestTemplate(t *testing.T) {
	view := CreateSearchResultView("index.html")
	data := model.SearchResult{}
	err := template.Execute(os.Stdout, view.Reader(os.Stdout,data))
	if err != nil {
		panic(err)
	}

}
