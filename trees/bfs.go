package main

import (
	"fmt"
)

type Edge struct {
	source      int
	destination int
}

type Graph struct {
	adjacentList [][]int
}

func New(edges []Edge, n int) *Graph {

	g := Graph{[][]int{}}
	for _, v := range edges {
		adS := append(g.adjacentList[v.source], v.destination)
		adD := append(g.adjacentList[v.destination], v.source)
	}
	return &g
}

func recursiveBFS(graph Graph, queue []int, disc []bool) {
	if len(queue) == 0 {
		return
	}

}

func main() {
	fmt.Println("BFS")
}
