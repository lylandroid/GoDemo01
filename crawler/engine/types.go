package engine

import "reflect"

type ParserFunc func(body []byte, url string) ParseResult

type Parser interface {
	Parse(body []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}

type NilParser struct {
}

func (NilParser) Parse(body []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return reflect.TypeOf(NilParser{}).Name(), nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(body []byte, url string) ParseResult {
	return f.Parse(body, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func CreateFuncParser(p ParserFunc, name string) FuncParser {
	return FuncParser{
		parser: p,
		name:   name,
	}
}
