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
		for j := range mazes[i] {
			fmt.Fscan(file, &mazes[i][j])
			//fmt.Print(mazes[i][j], " j=",j," ")
		}
	}
	return mazes
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(t point) point {
	return point{p.i + t.i, p.j + t.j}
}

func (p point) as(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true

}

func walk(mazes [][]int, start point, end point) [][]int {
	steps := make([][]int, len(mazes))
	for i := range steps {
		steps[i] = make([]int, len(mazes[i]))
	}
	Q := []point{start}
	count := 0
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}

		for _, dir := range dirs {
			count++
			next := cur.add(dir)
			//是否在区域内 & 是否是墙
			if v, ok := next.as(mazes); !ok || v != 0 {
				continue
			}
			//已经走过
			if v, ok := next.as(steps); !ok || v != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.as(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	fmt.Println("count: ", count)
	return steps
}

func print(arr [][]int) {
	for _, rows := range arr {
		for _, v := range rows {
			fmt.Printf("\t%d ", v)
		}
		fmt.Println()
	}
}

func main() {
	mazes := readMaze("maze/maze.in")
	fmt.Println("--------------------start main----------------------")
	print(mazes)
	fmt.Println("------------------------------------------")
	mazes2 := walk(mazes, point{0, 0}, point{len(mazes) - 1, len(mazes[0]) - 1})
	print(mazes2)

}
