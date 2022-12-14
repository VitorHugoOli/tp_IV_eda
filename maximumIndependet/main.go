package main

import "fmt"

// Graph represents a graph with vertices and edges.
type Graph struct {
	Vertices []int
	Edges    map[int][]int
}

// NewGraph creates a new graph with the given vertices and edges.
func NewGraph(vertices []int, edges map[int][]int) *Graph {
	return &Graph{vertices, edges}
}

func (g *Graph) graphSets() []int {
	if len(g.Vertices) == 0 {
		return []int{}
	}

	if len(g.Vertices) == 1 {
		return []int{g.Vertices[0]}
	}

	vCurrent := g.Vertices[0]

	graph2 := Graph{
		Vertices: make([]int, 0, len(g.Vertices)),
		Edges:    make(map[int][]int),
	}

	for i, v := range g.Vertices {
		if i == 0 {
			continue
		}
		graph2.Vertices = append(graph2.Vertices, v)
		graph2.Edges[v] = g.Edges[v]
	}

	res1 := graph2.graphSets()

	for _, v := range g.Edges[vCurrent] {
		// if v is in graph2.Vertices
		for i, v2 := range graph2.Vertices {
			if v == v2 {
				graph2.Vertices = append(graph2.Vertices[:i], graph2.Vertices[i+1:]...)
				delete(graph2.Edges, v)
			}
		}
	}

	res2 := []int{vCurrent}
	res2 = append(res2, graph2.graphSets()...)

	if len(res1) > len(res2) {
		return res1
	}
	return res2
}

// maximalIndependentSet returns the maximal independent

func main() {
	// Create a new graph with the given vertices and edges.
	g := NewGraph([]int{1, 2, 3, 4, 5, 6, 7, 8}, map[int][]int{
		1: {2, 3},
		2: {1, 4},
		3: {1},
		4: {2, 8},
		5: {6},
		6: {5, 7},
		7: {6},
		8: {4},
	})

	// Find the maximum independent set of the graph.
	S := g.graphSets()

	// Print the maximum independent set.
	fmt.Println("The maximum independent set is:", S)
}

func (g *Graph) graphSetsNonDeterministic() []int {
	maxSubset := OraculoMaxSubSet(graph)
	//check if the maxSubset is a valid subset
	for i, v := range maxSubset {
		if i == len(maxSubset)-1 {
			break
		}
		for _, v2 := range g.Edges[v] {
			if v2 == maxSubset[i+1] {
				return []int{}
			}
		}
	}
	return maxSubset
}
