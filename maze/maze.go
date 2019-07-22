package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	//bytes, _ := ioutil.ReadAll(file)
	//fmt.Println(string(bytes))

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	mazes := make([][]int, row)
	for i := range mazes {
		mazes[i] = make([]int, col)
		/*for j := range mazes[i] {
			fmt.Fscanf(file, "%d", &mazes[i][j])
			//fmt.Print(mazes[i][j], " j=",j," ")
		}*/
		var l string
		fmt.Fscan(file,&l)
		fmt.Println(l)
		fmt.Fscan(file,&l)
		fmt.Println(l)
	}
	return mazes
}

func main() {
	mazes := readMaze("maze/maze.in")
	fmt.Println("--------------------start main----------------------")
	for _, rows := range mazes {
		for _, v := range rows {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	fmt.Println("--------------------end main----------------------", mazes)
}
