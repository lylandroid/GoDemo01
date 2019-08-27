package main

import (
	"fmt"
	"reflect"
)

type Test2 struct {

}


func main() {

	var parserFunc ParserFunc
	parserFunc = NilParser{}.Parse
	fmt.Println(reflect.TypeOf(Test2{}).Name())
	fmt.Println(reflect.TypeOf(parserFunc).Name())
}
