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
	
	graph.BFSTraverse(7)
	fmt.Println()
	graph.DFSTraverse(7)
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

func (g *Graph)BFSByLinkedList(s, t int){
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
	visited[s] = true
	prev[s] = -1
	
	for queue.Len() > 0 && notFound{
		top := queue.Front()
		queue.Remove(top)
		if g.adj[top.Value.(int)].Len() <= 0{
			continue
		}
		
		linkedList := g.adj[top.Value.(int)]
	    for e := linkedList.Front(); e != nil; e = e.Next(){
	    	k := e.Value.(int)
	    	if !visited[k]{
	    		prev[k] = top.Value.(int)
	    		if k == t{
	    			notFound = false
	    			break
	    		}
	    		visited[k] = true
	    		queue.PushBack(k)
	    	}
	    }
	}
	
	if !notFound{
		printPrev(7, prev)
	} 
}

func BFSBySlice(s, t int) {
	if s == t{
		return
	}
}

func (g *Graph)BFSTraverse(s int){
	if s >= g.v{
		fmt.Printf("the value of s is large than the number of this graph %d\n", g.v)
		return
	}
	
	var queue []int
	prev := make([]int, g.v)
	for index := range prev{
		prev[index] = -1
	}
	visited := make([]bool, g.v)
	queue = append(queue, s)
	visited[s] = true
	
	for len(queue) > 0{
		top := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", top)
		linkedList := g.adj[top]
		for e := linkedList.Front(); e != nil; e = e.Next(){
			k := e.Value.(int)
			if !visited[k]{
				prev[k] = top
				visited[k] =  true
				queue = append(queue, k)
			}
		}
	}
}

func (g *Graph)DFSTraverse(t int){
	if t >= g.v || t < 0{
		fmt.Println("input error")
		return
	}
	
	visited := make([]bool, g.v)
	visited[t] =  true
	prev := make([]int, g.v)
	for i := range prev{
		prev[i] = -1
	}
	
	g.recurse(t, visited, prev)
}

func (g *Graph)recurse(t int, visited []bool, prev []int){
	fmt.Printf("%d ", t)
	linkedList := g.adj[t]
	for e := linkedList.Front(); e != nil; e = e.Next(){
		k := e.Value.(int)
		if !visited[k]{
			visited[k] = true
			prev[k] = t
			g.recurse(k, visited, prev)
		}
	}
}

func printPrev(t int, prev []int){
	index := t
	for prev[index] != -1 {
		fmt.Printf("%d -> %d\n", index, prev[index])
		index = prev[index]
	}
}