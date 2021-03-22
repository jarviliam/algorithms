package main

import (
	"fmt"
	"time"
)

type Graph struct {
	visited  []bool
	adjacent [][]int
}
type Vertex struct {
	value   int
	visited bool
	adj     []*Vertex
}

func NewVert(v int) *Vertex {
	return &Vertex{value: v}
}

func (g *Graph) DFS(v *Vertex) {
	if v.visited {
		return
	}
	v.visited = true
	for _, x := range v.adj {
		g.DFS(x)
	}
}

func (v *Vertex) add(vert ...*Vertex) {
	v.adj = append(v.adj, vert...)
}

func main() {

	graph := &Graph{}
	vert1 := NewVert(0)
	vert2 := NewVert(1)
	vert3 := NewVert(2)
	vert4 := NewVert(9)
	vert1.add(vert1, vert4, vert2)
	vert2.add(vert1)
	vert4.add(vert3)
	start := time.Now()
	fmt.Println("Starting")
	graph.DFS(vert1)
	end := time.Now()
	fmt.Printf("Finished at %d\n", end.Sub(start))
}
