package main

import (
	"os"
	"runtime"
	"time"
	"fmt"
)

const UINT_MAX = 1024

func main() {
	//var mpas = make(map[int]string)
	fmt.Println(UINT_MAX)
	var arr [UINT_MAX] int
	for i := 0; i < UINT_MAX; i++ {
		go func(i int) {
			for {
				arr[i] = arr[i] + 1
				pid := os.Getegid()
				fmt.Println("pid: ",pid, "i=",arr[i])
				runtime.Gosched() //交出控制权
				//mpas[i] = "Hello from goroutine  i=" + string(i)
				//fmt.Printf("Hello from goroutine  i=%d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
	/*for k, v := range arr {
		fmt.Println(k, v)
	}*/
	fmt.Println(arr)

}
