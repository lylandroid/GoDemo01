package main

import (
	"../map"
	"testing"
)

/*func main() {
	TestTriangle()
}*/

func TestTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
		{-1, -1, -2},
		{30000, 40000, 70000},
	}

	for _, tt := range tests {
		if c := add(tt.a, tt.b); c != tt.c {
			t.Errorf("Err a(%d) + b(%d) = c(%d)", tt.a, tt.b, c)
			//fmt.Printf("Err a(%d) + b(%d) = c(%d)", tt.a, tt.b, c)
		}
	}
	_map.LenStringSub("jfsadjfldsfjldsjfdjfoewferd你好")
}

func add(a int, b int) int {
	return a + b
}

//---------------------------------------------------

