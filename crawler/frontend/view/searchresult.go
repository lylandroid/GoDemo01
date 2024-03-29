package view

import (
	"html/template"
	"io"
	"../model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}

}

func (s SearchResultView) Reader(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
