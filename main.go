package main

import (
	"fmt"

	"github.com/dkhrunov/dsa-go/structures/graph"
)

func main() {

	createComponents := func(g *graph.Graph) {
		g.AddVertex("A")
		g.AddVertex("B")
		g.AddVertex("C")
		g.AddVertex("D")
		g.AddVertex("E")
		g.AddVertex("F")
		g.AddVertex("G")
		g.AddVertex("H")
		g.AddVertex("I")
		g.AddVertex("J")
		g.AddVertex("K")

		g.AddEdge("A", "B")
		g.AddEdge("B", "C")
		g.AddEdge("C", "A")

		g.AddEdge("E", "F")

		g.AddEdge("G", "H")
		g.AddEdge("H", "I")
		g.AddEdge("I", "J")
		g.AddEdge("J", "K")
		g.AddEdge("J", "H")
		g.AddEdge("K", "G")
		g.AddEdge("K", "H")
	}

	printComponents := func(g *graph.Graph) {
		components, _ := g.FindComponents()
		for i, group := range components {
			fmt.Printf("[%v] ", i)
			for e := group.Front(); e != nil; e = e.Next() {
				fmt.Printf("%v, ", e.Value)
			}
			fmt.Println()
		}
	}

	gL := graph.New()

	createComponents(gL)
	fmt.Println("Graph (List):")
	fmt.Println("--------------------")
	fmt.Println(gL)
	fmt.Println("Connected Components:")
	printComponents(gL)
	fmt.Println()

	gM := graph.NewMatrix()

	createComponents(gM)
	fmt.Println("Graph (Matrix):")
	fmt.Println("--------------------")
	fmt.Println(gM)
	fmt.Println("Connected Components:")
	printComponents(gM)
	fmt.Println()

	digL := graph.NewDirected()

	createComponents(digL)
	fmt.Println("Digraph (List):")
	fmt.Println("--------------------")
	fmt.Println(digL)
	fmt.Println("Connected Components:")
	printComponents(digL)
	fmt.Println()

	digM := graph.NewDirectedMatrix()

	createComponents(digM)
	fmt.Println("Digraph (Matrix):")
	fmt.Println("--------------------")
	fmt.Println(digM)
	fmt.Println("Connected Components:")
	printComponents(digM)
	fmt.Println()

	// ----------------
	// Digraph
	// ----------------

	//            |---> [C] ----|
	//            |             |
	//            |             ˅
	//  [A] ---> [B]           [E] ---> [F]
	//            ^             |
	//            |             |
	//            |---- [D] <---|

	digraph := graph.NewDirected()

	digraph.AddVertex("A")
	digraph.AddVertex("B")
	digraph.AddVertex("C")
	digraph.AddVertex("D")
	digraph.AddVertex("E")
	digraph.AddVertex("F")

	digraph.AddEdge("A", "B")
	digraph.AddEdge("B", "C")
	digraph.AddEdge("C", "E")
	digraph.AddEdge("E", "F")
	digraph.AddEdge("E", "D")
	digraph.AddEdge("D", "B")

	fmt.Println("Directed Graph (List):")
	fmt.Println("--------------------")

	fmt.Println(digraph)
	fmt.Printf("Vertices: %v\n", digraph.Vertices())
	fmt.Printf("Edges:    %v\n", digraph.Edges())
	fmt.Println()

	fmt.Printf("MaxEdges: %v\n", digraph.MaxEdges())
	fmt.Printf("Density:  %v\n", digraph.Density())
	fmt.Println()

	fmt.Printf("IsCyclic: %v\n", digraph.IsCyclic())
	fmt.Println()

	fmt.Print("BFS: ")
	if err := digraph.BFS("A", printVertex); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Print("DFS: ")
	if err := digraph.DFS("A", printVertex); err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println()

	// ------------------
	//  Graph
	// ------------------

	// 	[A] ----- [C] ----- [E]
	// 		 \     /   \     /   \
	//      \   /     \   /     \
	// 			 [B]       [D]       \
	//          \     /           \
	//           \   /             \
	//            [F] ------------ [G]

	graph := graph.NewMatrix()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")
	graph.AddVertex("G")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "F")
	graph.AddEdge("C", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "E")
	graph.AddEdge("D", "F")
	graph.AddEdge("E", "G")
	graph.AddEdge("F", "G")

	// udGraph.DeleteVertex("D")
	// udGraph.DeleteVertex("A")

	// Delete edge
	// udGraph.DeleteEdge("D", "C")
	// udGraph.DeleteEdge("D", "E")
	// udGraph.DeleteEdge("D", "F")

	fmt.Println("Undirected Graph (Matrix):")
	fmt.Println("--------------------")

	fmt.Println(graph)
	fmt.Printf("Vertices: %v\n", graph.Vertices())
	fmt.Printf("Edges:    %v\n", graph.Edges())
	fmt.Println()

	fmt.Printf("MaxEdges: %v\n", graph.MaxEdges())
	fmt.Printf("Density:  %v\n", graph.Density())
	fmt.Println()

	fmt.Print("BFS: ")
	if err := graph.BFS("A", printVertex); err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Print("DFS: ")
	if err := graph.DFS("A", printVertex); err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

func printVertex(vertex string) {
	fmt.Printf("{%v}, ", vertex)
}
