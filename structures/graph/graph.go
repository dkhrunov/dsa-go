package graph

import (
	"container/list"
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

	ErrCyclicCheckOnlyForDirected = "cyclic check applied only for directed"
)

type GraphRepr interface {
	setDirected()
	IsDirected() bool
	Vertices() int
	Edges() int
	HasVertex(vertex string) bool
	AddVertex(vertex string) error
	DeleteVertex(vertex string) error
	HasEdge(source, target string) bool
	AddEdge(source, target string) error
	DeleteEdge(source, target string) error
	BFS(start string, callback func(node string)) error
	DFS(start string, callback func(node string)) error
	IsCyclic() bool
	FindComponents() ([]*list.List, error)
	String() string
}

type GraphOption func(gr GraphRepr)

func WithVertices(vertices []string) GraphOption {
	return func(gr GraphRepr) {
		for _, vertex := range vertices {
			gr.AddVertex(vertex)
		}
	}
}

func WithEdges(edges [][2]string) GraphOption {
	return func(gr GraphRepr) {
		for _, v := range edges {
			gr.AddEdge(v[0], v[1])
		}
	}
}

type Graph struct {
	repr GraphRepr
}

func New(opts ...GraphOption) *Graph {
	list := newAdjList(opts...)
	return &Graph{list}
}

func NewDirected(opts ...GraphOption) *Graph {
	list := newAdjList(opts...)
	list.setDirected()
	return &Graph{list}
}

func NewMatrix(opts ...GraphOption) *Graph {
	matrix := newAdjMatrix(opts...)
	return &Graph{matrix}
}

func NewList(opts ...GraphOption) *Graph {
	list := newAdjList(opts...)
	return &Graph{list}
}

func NewDirectedMatrix(opts ...GraphOption) *Graph {
	matrix := newAdjMatrix(opts...)
	matrix.setDirected()
	return &Graph{matrix}
}

func NewDirectedList(opts ...GraphOption) *Graph {
	list := newAdjList(opts...)
	list.setDirected()
	return &Graph{list}
}

func (g *Graph) Vertices() int {
	return g.repr.Vertices()
}

func (g *Graph) Edges() int {
	return g.repr.Edges()
}

// https://www.baeldung.com/cs/graphs-sparse-vs-dense
func (g *Graph) MaxEdges() float64 {
	vertices := g.repr.Vertices()

	if g.repr.IsDirected() {
		return float64(vertices * (vertices - 1))
	}

	return float64(vertices*(vertices-1)) / 2
}

// https://www.baeldung.com/cs/graphs-sparse-vs-dense
func (g *Graph) Density() float64 {
	edges := g.repr.Edges()
	maxEdges := g.MaxEdges()
	return float64(edges) / maxEdges
}

func (g *Graph) HasEdge(source, target string) bool {
	return g.repr.HasEdge(source, target)
}

func (g *Graph) AddVertex(vertex string) error {
	return g.repr.AddVertex(vertex)
}

func (g *Graph) DeleteVertex(vertex string) error {
	return g.repr.DeleteVertex(vertex)
}

func (g *Graph) AddEdge(source, target string) error {
	return g.repr.AddEdge(source, target)
}

func (g *Graph) DeleteEdge(source, target string) error {
	return g.repr.DeleteEdge(source, target)
}

func (g *Graph) BFS(start string, callback func(node string)) error {
	return g.repr.BFS(start, callback)
}

func (g *Graph) DFS(start string, callback func(node string)) error {
	return g.repr.DFS(start, callback)
}

func (g *Graph) IsCyclic() bool {
	return g.repr.IsCyclic()
}

func (g *Graph) FindComponents() ([]*list.List, error) {
	return g.repr.FindComponents()
}

func (g *Graph) String() string {
	return g.repr.String()
}
