package main

import "fmt"

func consts() {
}

func enums() {
	const (
		java = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Print(1 << 10)
}

func main() {
	enums()
}
