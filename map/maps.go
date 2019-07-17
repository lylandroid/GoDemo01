package _map

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":  "value",
		"name2": "value2",
		"name3": "value3",
	}

	m2 := make(map[string]string)
	var m3 map[string]int

	//var m2  map[string]string
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println("m[name]=" + m["name"])
	for k, v := range m {
		fmt.Printf("k=%s,v=%s ", k, v)
	}
	fmt.Println("-------------------")
	delete(m, "name")
	if name, ok := m["name"]; ok {
		fmt.Println("name:"+name+" ok:", ok)
	} else {
		fmt.Println("name="+name+"ok: ", ok)
	}
	fmt.Println(len(m))
	s := "abcabcbb"
	fmt.Println(LenStringSub(s))
	fmt.Println(LenStringSub("aaaaaaaaaa"))
	fmt.Println(LenStringSub(""))
	fmt.Println(LenStringSub("abcdefg"))
	/*for k, v := range []byte(s) {
		fmt.Println(k, v)
	}*/
}



func LenStringSub(s string) int {
	for i := 0; i < 20; i++ {
		s = s + s
	}
	var lastCurred = make(map[byte]int)
	fmt.Println("lastCurred", lastCurred)
	startIndex := 0
	maxLen := 0
	for k, v := range []byte(s) {
		if lastI, ok := lastCurred[v]; ok && lastI >= startIndex {
			startIndex = lastI + 1
		}
		if k-startIndex+1 > maxLen {
			maxLen = k - startIndex + 1
		}
		lastCurred[v] = k
	}
	fmt.Println("lastCurred2", lastCurred)
	return maxLen
}
