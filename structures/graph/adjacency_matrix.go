package graph

import (
	"bytes"
	"container/list"
	"fmt"
	"sync"

	"github.com/dkhrunov/dsa-go/structures/queue"
)

// Space complexity: O(n^2), where n is number of vertices
type adjMatrix struct {
	lock         sync.RWMutex
	v            int
	e            int
	undirected   bool
	vertices     map[string]int
	verticeNames map[int]string
	matrix       [][]int8
}

func newAdjMatrix(opts ...GraphOption) *adjMatrix {
	matrix := &adjMatrix{
		undirected:   true,
		vertices:     make(map[string]int),
		verticeNames: make(map[int]string),
		matrix:       make([][]int8, 0),
	}

	for _, opt := range opts {
		opt(matrix)
	}

	return matrix
}

func (m *adjMatrix) setDirected() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.undirected = false
}

func (m *adjMatrix) IsDirected() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return !m.undirected
}

func (m *adjMatrix) Vertices() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.v
}

func (m *adjMatrix) Edges() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.e
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (m *adjMatrix) HasVertex(vertex string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	_, has := m.vertices[vertex]
	return has
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(1)
func (m *adjMatrix) AddVertex(vertex string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.vertices[vertex]; ok {
		return ErrVertexAlreadyExists(vertex)
	}

	nextIdx := m.v
	// Add new vertex to maps
	m.vertices[vertex] = nextIdx
	m.verticeNames[nextIdx] = vertex
	// Add to each row new column
	for i := range m.matrix {
		m.matrix[i] = append(m.matrix[i], 0)
	}
	// Add new row
	m.matrix = append(m.matrix, make([]int8, nextIdx+1))

	m.v++

	return nil
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(n), where n is number of vertices
func (m *adjMatrix) DeleteVertex(vertex string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	vertexIdx, ok := m.vertices[vertex]
	if !ok {
		return ErrVertexNotFound(vertex)
	}

	// Delete row
	m.matrix = append(m.matrix[:vertexIdx], m.matrix[vertexIdx+1:]...)
	// Delete column
	for i := range m.matrix {
		m.matrix[i] = append(m.matrix[i][:vertexIdx], m.matrix[i][vertexIdx+1:]...)
	}

	// Delete from vertices map
	delete(m.vertices, vertex)
	// Delete from verticeNames map
	delete(m.verticeNames, vertexIdx)
	for k, idx := range m.vertices {
		if idx > vertexIdx {
			idx -= 1
			// Update vertices map
			m.vertices[k] = idx
		}
		// Update verticeNames map
		m.verticeNames[idx] = k
	}

	m.v--

	return nil
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (m *adjMatrix) HasEdge(source, target string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	i, ok := m.vertices[source]
	if !ok {
		return false
	}

	j, ok := m.vertices[target]
	if !ok {
		return false
	}

	if m.undirected {
		return m.matrix[i][j] == 1 && m.matrix[j][i] == 1
	} else {
		return m.matrix[i][j] == 1
	}
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (m *adjMatrix) AddEdge(source, target string) error {
	if ok := m.HasEdge(source, target); ok {
		return ErrEdgeAlreadyExists(source, target)
	}

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

	m.matrix[i][j] = 1

	if m.undirected {
		m.matrix[j][i] = 1
	}

	m.e++

	return nil
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (m *adjMatrix) DeleteEdge(source, target string) error {
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

	if m.undirected {
		m.matrix[j][i] = 0
	}

	m.e--

	return nil
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(n), where n is number of vertices
func (m *adjMatrix) BFS(start string, callback func(vertex string)) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	visited := make([]bool, m.v)
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

		for i := 0; i < m.v; i++ {
			if !visited[i] && m.matrix[currIdx][i] == 1 {
				vertex, ok := m.verticeNames[i]
				if !ok {
					return ErrVertexNotFound("nil")
				}
				visited[i] = true
				queue.EnQueue(vertex)
			}
		}
	}

	return nil
}

// Time complexity: O(v+e), where v is number of vertices, and e is number of edges
//
// Space complexity: O(v+e), where v is number of vertices, and e is number of edges
func (m *adjMatrix) DFS(start string, callback func(vertex string)) error {
	visited := make([]bool, m.Vertices())
	return m.dfs(start, callback, visited)
}

func (m *adjMatrix) dfs(vertex string, callback func(vertex string), visited []bool) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

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
			m.dfs(v, callback, visited)
		}
	}

	return nil
}

// Time complexity: O(v+e), where v is number of vertices, and e is number of edges
//
// Space complexity: O(v), where v is number of vertices
func (m *adjMatrix) IsCyclic() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.undirected {
		panic(ErrCyclicCheckOnlyForDirected)
	}

	visited := make([]bool, m.Vertices())
	recStack := make([]bool, m.Vertices())

	for i := range m.matrix {
		if !visited[i] && m.isCyclicRec(i, visited, recStack) {
			return true
		}
	}

	return false
}

func (m *adjMatrix) isCyclicRec(i int, visited, recStack []bool) bool {
	if !visited[i] {
		// Mark the current node as visited
		// and part of recursion stack
		visited[i] = true
		recStack[i] = true

		for j := range m.matrix[i] {
			// Check only nodes that has edges
			if m.matrix[i][j] != 1 {
				continue
			}

			if !visited[j] && m.isCyclicRec(j, visited, recStack) {
				return true
			} else if recStack[j] {
				return true
			}
		}
	}

	// Remove the vertex from recursion stack
	recStack[i] = false
	return false
}

func (m *adjMatrix) FindComponents() ([]*list.List, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	groupId := -1
	components := make([]*list.List, 0, m.Vertices())
	visited := make([]bool, m.Vertices())

	grouping := func(vertex string) {
		if len(components)-1 < groupId {
			components = append(components, list.New())
		}
		components[groupId].PushBack(vertex)
	}

	for vertex := range m.vertices {
		i, ok := m.vertices[vertex]
		if !ok {
			return nil, ErrVertexNotFound(vertex)
		}

		if !visited[i] {
			groupId++
			if err := m.dfs(vertex, grouping, visited); err != nil {
				return nil, err
			}
		}
	}

	return components, nil
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(1)
func (m *adjMatrix) String() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if len(m.matrix) == 0 {
		return "[]"
	}

	var buffer bytes.Buffer
	buffer.WriteString("   ")
	for i := range m.matrix {
		buffer.WriteString(fmt.Sprintf("%v ", m.verticeNames[i]))
	}
	buffer.WriteString("\n")
	for i, row := range m.matrix {
		buffer.WriteString(fmt.Sprintf("%v ", m.verticeNames[i]))
		buffer.WriteString(fmt.Sprintf("%v\n", row))
	}

	return buffer.String()
}
