package engine

import "reflect"

type ParserFunc func(url string, body []byte, bean interface{}) ParseResult

type Parser interface {
	Parse(url string, body []byte, bean interface{}) ParseResult
	Serialize() (FuncName string, args interface{})
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

func (NilParser) Parse(url string, body []byte, bean interface{}) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return reflect.TypeOf(NilParser{}).Name(), nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(url string, body []byte, bean interface{}) ParseResult {
	return f.parser(url, body, bean)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
