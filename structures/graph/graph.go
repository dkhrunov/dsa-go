package graph

import (
	"fmt"
)

var (
	ErrVertexNotFound = func(vertex string) error {
		return fmt.Errorf("vertex \"%v\" not found", vertex)
	}

	ErrVertexAlreadyExists = func(vertex string) error {
		return fmt.Errorf("vertex \"%v\" already exists", vertex)
	}

	ErrEdgeNotFound = func(source, target string) error {
		return fmt.Errorf("edge \"%v\" -> \"%v\" not found", source, target)
	}

	ErrEdgeAlreadyExists = func(source, target string) error {
		return fmt.Errorf("edge \"%v\" -> \"%v\" already exists", source, target)
	}

	// ErrEdgeCreatesCycle    = errors.New("edge would create a cycle")
	// ErrVertexHasEdges      = errors.New("vertex has edges")
)

type GraphRepresentation interface {
	AddVertex(vertex string) error
	DeleteVertex(vertex string) error
	AddEdge(source, target string) error
	DeleteEdge(source, target string) error
	BFS(start string, callback func(node string)) error
	DFS(start string, callback func(node string)) error
	Print()
}

type Graph struct {
	representation GraphRepresentation
}

func New(r GraphRepresentation) Graph {
	return Graph{r}
}

func NewWithAdjacencyMatrix(opts ...AdjacencyMatrixOption) Graph {
	return Graph{NewAdjacencyMatrix(opts...)}
}

func NewWithAdjacencyList(opts ...AdjacencyListOption) Graph {
	return Graph{NewAdjacencyList(opts...)}
}

func (g Graph) AddVertex(vertex string) error {
	return g.representation.AddVertex(vertex)
}

func (g Graph) DeleteVertex(vertex string) error {
	return g.representation.DeleteVertex(vertex)
}

func (g Graph) AddEdge(source, target string) error {
	return g.representation.AddEdge(source, target)
}

func (g Graph) DeleteEdge(source, target string) error {
	return g.representation.DeleteEdge(source, target)
}

func (g Graph) BFS(start string, callback func(node string)) error {
	return g.representation.BFS(start, callback)
}

func (g Graph) DFS(start string, callback func(node string)) error {
	return g.representation.DFS(start, callback)
}

func (g Graph) Print() {
	g.representation.Print()
}
