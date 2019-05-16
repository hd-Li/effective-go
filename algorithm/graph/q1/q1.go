package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj []*list.List
	v int
}

func main(){
	graph := newGraph(5)
	fmt.Printf("the type of graph is %T", graph)
}

func newGraph(v int) *Graph{
	graph := &Graph{}
	adj := make([]*list.List, v)
	for i := range adj{
		adj[i] = list.New()
		fmt.Println(i)
	}
	graph.adj = adj
	graph.v = v
	
	return graph
}