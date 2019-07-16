package main

import (
	"../adder"
	"bufio"
	"fmt"
	"os"
)

func writerFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := adder.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func myDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}

func main() {
	//myDefer()
	writerFile("defer.txt")
}
