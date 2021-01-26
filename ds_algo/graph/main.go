package main

import "fmt"

// Graph
type Graph struct {
	vertices []*Vertex
}

// Vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Printf("Vertex %v not added because it is an existing key", k)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("Invalid edge (%v-->%v)\n", from, to)
	}
	if contains(g.GetVertex(from).adjacent, to) {
		fmt.Printf("Vertex %v not added because it is an existing edge", to)
		return
	}
	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

// Get Vertex
func (g *Graph) GetVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// Contain
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	// test.AddVertex(0)

	test.AddEdge(1, 2)
	test.AddEdge(1, 3)
	test.AddEdge(1, 3)
	test.AddEdge(2, 4)
	test.AddEdge(4, 3)
	test.Print()
}
