package worker

import (
	"../../crawler/engine"
	"../config"
	"../../crawler/parser/zhenai"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	FuncName string
	Args     interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type SerializeParseResult struct {
	Items              []engine.Item
	SerializedRequests []Request
}

func SerializeRequest(req engine.Request) Request {
	funcName, args := req.Parser.Serialize()
	return Request{
		Url: req.Url,
		Parser: SerializedParser{
			FuncName: funcName,
			Args:     args,
		},
	}
}

func SerializeResult(parseRequest engine.ParseResult) SerializeParseResult {
	result := SerializeParseResult{
		Items: parseRequest.Items,
	}
	for _, req := range parseRequest.Requests {
		result.SerializedRequests = append(result.SerializedRequests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(req Request) (engine.Request, error) {
	parser, err := DeserializeParser(req.Parser)
	return engine.Request{
		Url:    req.Url,
		Parser: parser,
	}, err
}

func DeserializeResult(r SerializeParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.SerializedRequests {
		deserializeParser, err := DeserializeParser(req.Parser)
		if err != nil {
			log.Printf("error deserializeing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:    req.Url,
			Parser: deserializeParser,
		})
	}
	return result
}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FuncName {
	case config.FuncParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.FuncParseCityList), nil
	case config.FuncParseProfileList:
		return engine.NewFuncParser(parser.ParseProfileList, config.FuncParseProfileList), nil
	case config.FuncParseProfile:
		return engine.NewFuncParser(parser.ParseProfile, config.FuncParseProfile), nil
	case config.FuncNilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("nu known parse name")
	}
}

