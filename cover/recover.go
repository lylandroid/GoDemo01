package main

import (
	"fmt"
)

func tryCover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error record: ", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is err"))
	panic(123)
}

func main() {
	tryCover()
}
