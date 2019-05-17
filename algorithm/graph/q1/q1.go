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
	graph := newGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)
	
	graph.BFS(0, 7)
	
}

func newGraph(v int) *Graph{
	graph := &Graph{}
	adj := make([]*list.List, v)
	for i := range adj{
		adj[i] = list.New()
	}
	graph.adj = adj
	graph.v = v
	
	return graph
}

func (g *Graph)addEdge(s, t int){
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

func (g *Graph)BFS(s, t int){
	if s == t{
		return
	}
	
	visited := make([]bool, g.v)
	queue := list.New()
	prev := make([]int, g.v)
	for i := range prev{
		prev[i] = -1
	}
	notFound := true
	
	queue.PushBack(s)
	prev[s] = -1
	for queue.Len() > 0 && notFound{
		top := queue.Front()
		if visited[top.Value.(int)] == false{
			visited[top.Value.(int)] = true
	    }
		if g.adj[top.Value.(int)].Len() <= 0{
			continue
		}
		pre := top
		e := g.adj[top.Value.(int)].Front()
	    for{
	    	if e.Value.(int) == t{
	    		notFound = false
	    		prev[e.Value.(int)] = pre.Value.(int)
	    		break
	    	}
	    	if visited[e.Value.(int)] == false{
	    		visited[e.Value.(int)] = true
	    		queue.PushBack(e.Value.(int))
	    		prev[e.Value.(int)] = pre.Value.(int)
	    	}
	    	if e == g.adj[top.Value.(int)].Back(){
	    		break
	    	}
	    	pre = e
	    	e =  e.Next()
	    }
	    
	    queue.Remove(top)
	}
	if !notFound{
		printPrev(7, prev)
	}
	
	 
}

func printPrev(t int, prev []int){
	index := t
	for prev[index] != -1 {
		fmt.Printf("%d -> %d\n", index, prev[index])
		index = prev[index]
	}
}