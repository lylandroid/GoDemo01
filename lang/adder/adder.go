package adder

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(i int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func Fibonacci() intFunc {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intFunc func() int

func (f intFunc) Read(p []byte) (n int, err error) {
	next := f()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	/*a := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, a = a(i)
		fmt.Println(i, sum)

	}*/
	/*f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())*/

	f := Fibonacci()
	printFileContents(f)

}
