package main

// ColorGraph colors the vertices of the graph using the simple coloring algorithm.
// The resulting map contains the color assigned to each vertex.
//func (g *Graph) ColorGraph() map[int]int {
//	colors := make(map[int]int)
//
//	for _, v := range g.vertices {
//		// Consider each of the current vertex's neighbors that have already been colored.
//		usedColors := make(map[int]bool)
//		for _, n := range g.edges[v] {
//			if c, ok := colors[n]; ok {
//				usedColors[c] = true
//			}
//		}
//
//		// Assign the current vertex the first unused color.
//		for c := 1; c <= len(g.vertices); c++ {
//			if !usedColors[c] {
//				colors[v] = c
//				break
//			}
//		}
//	}
//
//	return colors
//}
//
//// ColorGraphMinColors colors the vertices of the graph using the minimum number of colors possible.
//// The resulting map contains the color assigned to each vertex.
//func (g *Graph) ColorGraphMinColors() map[int]int {
//	colors := make(map[int]int)
//
//	// Initialize the available colors to the maximum possible.
//	availableColors := make([]int, len(g.vertices))
//	for i := range availableColors {
//		availableColors[i] = i + 1
//	}
//
//	for _, v := range g.vertices {
//		// Consider each of the current vertex's neighbors that have already been colored.
//		usedColors := make(map[int]bool)
//		for _, n := range g.edges[v] {
//			if c, ok := colors[n]; ok {
//				usedColors[c] = true
//			}
//		}
//
//		// Remove the used colors from the list of available colors.
//		for c := range usedColors {
//			availableColors = removeFromSlice(availableColors, c)
//		}
//
//		// Assign the current vertex the first available color.
//		colors[v] = availableColors[0]
//	}
//
//	return colors
//}
//
//// removeFromSlice removes the given element from the slice.
//func removeFromSlice(s []int, elem int) []int {
//	for i, e := range s {
//		if e == elem {
//			return append(s[:i], s[i+1:]...)
//		}
//	}
//	return s
//}
//
//// ColorGraphBacktracking colors the vertices of the graph using a backtracking algorithm.
//// The resulting map contains the color assigned to each vertex.
//func (g *Graph) ColorGraphBacktracking() map[int]int {
//	colors := make(map[int]int)
//
//	// Initialize the available colors to the maximum possible.
//	availableColors := make([]int, len(g.vertices))
//	for i := range availableColors {
//		availableColors[i] = i + 1
//	}
//
//	// Recursively color the vertices of the graph.
//	if !g.colorGraphBacktrackingRecursive(colors, availableColors, 0) {
//		return nil
//	}
//
//	return colors
//}
//
//// colorGraphBacktrackingRecursive is a helper function for the backtracking algorithm.
//func (g *Graph) colorGraphBacktrackingRecursive(colors map[int]int, availableColors []int, cur int) bool {
//	// If all vertices have been colored, return true.
//	if cur == len(g.vertices) {
//		return true
//	}
//
//	// Consider each of the current vertex's neighbors that have already been colored.
//	usedColors := make(map[int]bool)
//	for _, n := range g.edges[g.vertices[cur]] {
//		if c, ok := colors[n]; ok {
//			usedColors[c] = true
//		}
//	}
//
//	// Remove the used colors from the list of available colors.
//	for c := range usedColors {
//		availableColors = removeFromSlice(availableColors, c)
//	}
//
//	// Try each available color for the current vertex.
//	for _, c := range availableColors {
//		colors[g.vertices[cur]] = c
//		if g.colorGraphBacktrackingRecursive(colors, availableColors, cur+1) {
//			return true
//		}
//	}
//
//	// If no valid coloring was found, revert the changes and return false.
//	delete(colors, g.vertices[cur])
//	return false
//}
//
////func nonDeterministicColoring(g *Graph) map[int]int {
////	colors := make(map[int]int)
////
////	for i, e := range g.vertices {
////		colors[e] = Oracle(i)
////	}
////
////	// Check if the coloring is valid.
////	for _, v := range g.vertices {
////		for _, n := range g.edges[v] {
////			if colors[n] == colors[v] {
////				return nil
////			}
////		}
////	}
////	return colors
////}
//
//func (g *Graph) nonPolynomialGraphColoring(colors []int) map[int]int {
//	coloredVertices := make(map[int]int)
//	for _, v1 := range g.vertices {
//		for _, v2 := range g.edges[v1] {
//			if _, ok := coloredVertices[v1]; !ok {
//				for _, c := range colors {
//					if c != coloredVertices[v2] {
//						coloredVertices[v1] = c
//						break
//					}
//				}
//			}
//		}
//	}
//	return coloredVertices
//}
