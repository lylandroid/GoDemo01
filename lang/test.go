package main

import (
	"fmt"
	"reflect"
)

type Test2 struct {

}

func main() {
	fmt.Println(reflect.TypeOf(Test2{}).Name())
}
