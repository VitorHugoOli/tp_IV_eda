package main

import "fmt"

// Graph represents a graph with vertices and edges.
type Graph struct {
	vertices []int
	edges    map[int][]int
}

// NewGraph creates a new graph with the given vertices and edges.
func NewGraph(vertices []int, edges map[int][]int) *Graph {
	return &Graph{vertices, edges}
}

func main() {
	// Create a graph with four vertices and five edges.
	g := NewGraph([]int{1, 2, 3, 4}, map[int][]int{
		1: []int{2, 3},
		2: []int{1, 3},
		3: []int{1, 2, 4},
		4: []int{3},
	})
	g2 := NewGraph([]int{1, 2, 3, 4, 5, 6, 7}, map[int][]int{
		1: []int{2, 3, 4},
		2: []int{1, 4, 5},
		3: []int{1, 6, 7},
		4: []int{1, 2, 7},
		5: []int{2, 6},
		6: []int{3, 5, 7},
		7: []int{3, 4, 6},
	})

	// Color the vertices of the graph.
	_, k := g.MinColorGraph()
	fmt.Println(k)

	_, k = g2.MinColorGraph()
	fmt.Println(k)
}

// colorGraph colors a given graph using at most k colors. It returns a
// map where the keys are vertices and the values are the assigned colors.
func (g *Graph) ColorGraph(k int) map[int]int {
	// Create a map to store the assigned colors for each vertex.
	colors := make(map[int]int)

	// Assign a color to each vertex.
	for _, vertex := range g.vertices {
		// Get the colors of the adjacent vertices.
		usedColors := make(map[int]bool)
		for _, adjacentVertex := range g.edges[vertex] {
			if c, ok := colors[adjacentVertex]; ok {
				usedColors[c] = true
			}
		}

		// Choose the first available color.
		for color := 1; color <= k; color++ {
			if !usedColors[color] {
				colors[vertex] = color
				break
			}
		}
	}

	return colors
}

func (g *Graph) MinColorGraph() (map[int]int, int) {

	colors := make(map[int]int)
	usedColors := make(map[int]bool)
	k := 0

	// Brute force over the number of vertices, until we find the chromatic number.
	for k, _ = range g.vertices {
		// Create a map to store the assigned colors for each vertex.

		// Assign a color to each vertex.
		for _, vertex := range g.vertices {
			// Get the colors of the adjacent vertices.
			for _, adjacentVertex := range g.edges[vertex] {
				if c, ok := colors[adjacentVertex]; ok {
					usedColors[c] = true
				}
			}

			// Choose the first available color.
			for color := 1; color <= k; color++ {
				if !usedColors[color] {
					colors[vertex] = color
					break
				}
			}
		}

		if len(colors) == len(g.vertices) {
			break
		}
	}

	return colors, k
}
