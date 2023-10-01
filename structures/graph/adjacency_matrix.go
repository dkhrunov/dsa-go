package graph

import (
	"fmt"
	"sync"

	"github.com/dkhrunov/dsa-go/structures/queue"
)

// Space complexity: O(n^2 + 2n), where n is number of vertices
type AdjacencyMatrix struct {
	lock         sync.RWMutex
	vertices     map[string]int
	verticeNames map[int]string
	matrix       [][]int8
}

type AdjacencyMatrixOption func(m *AdjacencyMatrix)

func NewAdjacencyMatrix(opts ...AdjacencyMatrixOption) *AdjacencyMatrix {
	matrix := &AdjacencyMatrix{
		vertices:     make(map[string]int),
		verticeNames: make(map[int]string),
		matrix:       make([][]int8, 0),
	}

	for _, opt := range opts {
		opt(matrix)
	}

	return matrix
}

func AdjacencyMatrixWithVerticales(vertices []string) AdjacencyMatrixOption {
	return func(m *AdjacencyMatrix) {
		for _, vertex := range vertices {
			m.AddVertex(vertex)
		}
	}
}

func AdjacencyMatrixWithEdges(edges [][2]string) AdjacencyMatrixOption {
	return func(m *AdjacencyMatrix) {
		for _, v := range edges {
			m.AddEdge(v[0], v[1])
		}
	}
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(1)
func (m *AdjacencyMatrix) AddVertex(vertex string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.vertices[vertex]; ok {
		return ErrVertexAlreadyExists(vertex)
	}

	idx := len(m.vertices)
	// Add new vertex to maps
	m.vertices[vertex] = idx
	m.verticeNames[idx] = vertex
	// Add to each row new column
	for i := range m.matrix {
		m.matrix[i] = append(m.matrix[i], 0)
	}
	// Add new row
	m.matrix = append(m.matrix, make([]int8, idx+1))

	return nil
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(n), where n is number of vertices
func (m *AdjacencyMatrix) DeleteVertex(vertex string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	vertexId, ok := m.vertices[vertex]
	if !ok {
		return ErrVertexNotFound(vertex)
	}

	// Delete row
	m.matrix = append(m.matrix[:vertexId], m.matrix[vertexId+1:]...)
	// Delete col
	for i := range m.matrix {
		m.matrix[i] = append(m.matrix[i][:vertexId], m.matrix[i][vertexId+1:]...)
	}

	// Update vertices map
	delete(m.vertices, vertex)
	for k, v := range m.vertices {
		if v > vertexId {
			m.vertices[k] = v - 1
		}
	}

	// Update verticeNames map
	delete(m.verticeNames, vertexId)
	for k, v := range m.vertices {
		m.verticeNames[v] = k
	}

	return nil
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (m *AdjacencyMatrix) AddEdge(source, target string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	i, ok := m.vertices[source]
	if !ok {
		return ErrVertexNotFound(source)
	}

	j, ok := m.vertices[target]
	if !ok {
		return ErrVertexNotFound(target)
	}

	if m.matrix[i][j] == 1 && m.matrix[j][i] == 1 {
		return ErrEdgeAlreadyExists(source, target)
	}

	m.matrix[i][j] = 1
	m.matrix[j][i] = 1

	return nil
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (m *AdjacencyMatrix) DeleteEdge(source, target string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	i, ok := m.vertices[source]
	if !ok {
		return ErrVertexNotFound(source)
	}

	j, ok := m.vertices[target]
	if !ok {
		return ErrVertexNotFound(target)
	}

	if m.matrix[i][j] == 0 && m.matrix[j][i] == 0 {
		return ErrEdgeNotFound(source, target)
	}

	m.matrix[i][j] = 0
	m.matrix[j][i] = 0

	return nil
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(n), where n is number of vertices
func (m *AdjacencyMatrix) BFS(start string, callback func(vertex string)) error {
	visited := make([]bool, len(m.vertices))
	queue := queue.New()

	i, ok := m.vertices[start]
	if !ok {
		return ErrVertexNotFound(start)
	}

	visited[i] = true
	queue.EnQueue(start)

	for queue.Len() > 0 {
		curr := queue.DeQueue().(string)
		callback(curr)

		currIdx, ok := m.vertices[curr]
		if !ok {
			return ErrVertexNotFound(curr)
		}

		for i := 0; i < len(m.vertices); i++ {
			if !visited[i] && m.matrix[currIdx][i] == 1 {
				vertex, ok := m.verticeNames[i]
				if !ok {
					return ErrVertexNotFound("nil")
				}
				queue.EnQueue(vertex)
				visited[i] = true
			}
		}
	}

	return nil
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(n), where n is number of vertices
func (m *AdjacencyMatrix) DFS(start string, callback func(vertex string)) error {
	visited := make([]bool, len(m.vertices))

	var dfs func(vertex string, callback func(vertex string), visited []bool) error
	dfs = func(vertex string, callback func(vertex string), visited []bool) error {
		i, ok := m.vertices[vertex]
		if !ok {
			return ErrVertexNotFound(vertex)
		}

		visited[i] = true
		callback(vertex)

		for j := 0; j < len(m.matrix[i]); j++ {
			if !visited[j] && m.matrix[i][j] == 1 {
				v, ok := m.verticeNames[j]
				if !ok {
					return ErrVertexNotFound("nil")
				}
				dfs(v, callback, visited)
			}
		}

		return nil
	}

	return dfs(start, callback, visited)
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(1)
func (m *AdjacencyMatrix) Print() {
	if len(m.matrix) == 0 {
		fmt.Print("[]")
		return
	}

	fmt.Print("   ")
	for i := range m.matrix {
		fmt.Printf("%v ", m.verticeNames[i])
	}
	fmt.Println()
	for i, row := range m.matrix {
		fmt.Printf("%v ", m.verticeNames[i])
		fmt.Println(row)
	}
}
