package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

func eval(a, b int, opt string) int {
	var result int
	switch opt {
	case "+":
		result = a + b
		fallthrough
	case "-":
		result = a - b
	default:
		fmt.Println("no find opt")
	}
	return result

}

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		result += "0"
	}
	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}
	return result
}

func readFile(filename string) {
	if file, error := os.Open(filename); error != nil {
		panic(error)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}
func div(opt func(a int, b int) int, a1 int, b1 int) {
	pointer := reflect.ValueOf(opt).Pointer()
	optName := runtime.FuncForPC(pointer).Name()
	fmt.Println(optName)

}

func pow2(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(arr ...int) int {
	total := 0
	/*for i := range arr {
		total += arr[i]
	}*/
	for k, v := range arr {
		fmt.Printf("k=%s,v=%s ", k, v)
	}
	return total
}

func addValue(a, b int) {
	a, b = b, a
}

func addRes(a, b *int) {

	fmt.Println("-------------")
	fmt.Println(a, b)
	*a, *b = *b, *a
}

func funCache() {

}

func main() {
	const filename = "abc.txt"
	/*if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(string(contents))
	}*/
	readFile(filename)

	fmt.Println(eval(1, 2, "+"))
	fmt.Println(convertToBin(0), convertToBin(1), convertToBin(2), convertToBin(10))
	div(pow2, 1, 2)
	fmt.Println(sum(1, 2, 3, 4))
	a, b := 1, 2
	addValue(a, b)
	fmt.Println(a, b)
	addRes(&a, &b)
	fmt.Println(a, b)

	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}
	fmt.Println(arr1, arr2, arr3, len(arr1))

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:6]

	fmt.Println(len(s1), cap(arr), s1, s2)
	var s []int
	fmt.Println(s)

}
