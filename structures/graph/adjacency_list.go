package graph

import (
	"bytes"
	"container/list"
	"fmt"
	"sort"
	"sync"

	"github.com/dkhrunov/dsa-go/structures/queue"
)

// Space complexity: O(n+m), where n is number of vertices, m is number of edges
//
// But, in the worst case of a complete graph, which contains n^2 edges,
// the time and space complexities reduce to O(n^2)
type adjList struct {
	lock       sync.RWMutex
	v          int
	e          int
	undirected bool
	vertices   map[string]int
	lists      []*list.List
}

// TODO
type listNode struct {
	name   string
	weight int
}

type vertexIdx struct {
	Vertex string
	Index  int
}

func newAdjList(opts ...GraphOption) *adjList {
	list := &adjList{
		undirected: true,
		vertices:   make(map[string]int),
		lists:      make([]*list.List, 0),
	}

	for _, opt := range opts {
		opt(list)
	}

	return list
}

func (l *adjList) setDirected() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.undirected = false
}

func (l *adjList) IsDirected() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return !l.undirected
}

func (l *adjList) Vertices() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.v
}

func (l *adjList) Edges() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.e
}

func (l *adjList) vertexIdx() []vertexIdx {
	var vIdx []vertexIdx
	for vertex, idx := range l.vertices {
		vIdx = append(vIdx, vertexIdx{vertex, idx})
	}

	sort.Slice(vIdx, func(i, j int) bool {
		return vIdx[i].Index < vIdx[j].Index
	})

	return vIdx
}

// Time complexity: O(1)
//
// Space complexity: O(1)
func (l *adjList) HasVertex(vertex string) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	_, has := l.vertices[vertex]
	return has
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(1)
func (l *adjList) AddVertex(vertex string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if _, ok := l.vertices[vertex]; ok {
		return ErrVertexAlreadyExists(vertex)
	}

	nextIdx := l.v
	// Add new vertex to maps
	l.vertices[vertex] = nextIdx
	// Add new list
	l.lists = append(l.lists, list.New())

	l.v++

	return nil
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(1)
func (l *adjList) DeleteVertex(vertex string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	// Delete all edges associated with that vertex
	for _, list := range l.lists {
		for e := list.Front(); e != nil; e = e.Next() {
			if e.Value == vertex {
				list.Remove(e)
			}
		}
	}

	vertexIdx, ok := l.vertices[vertex]
	if !ok {
		return ErrVertexNotFound(vertex)
	}

	// Delete from slice
	l.lists = append(l.lists[:vertexIdx], l.lists[vertexIdx+1:]...)

	// Delete from map
	delete(l.vertices, vertex)
	for k, idx := range l.vertices {
		if idx > vertexIdx {
			l.vertices[k] = idx - 1
		}
	}

	l.v--

	return nil
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(1)
func (l *adjList) HasEdge(source, target string) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	i, ok := l.vertices[source]
	if !ok {
		return false
	}

	j, ok := l.vertices[target]
	if !ok {
		return false
	}

	if l.undirected {
		sourceToTarget := false
		for e := l.lists[i].Front(); e != nil; e = e.Next() {
			if e.Value.(string) == target {
				sourceToTarget = true
			}
		}

		targetToSource := false
		for e := l.lists[j].Front(); e != nil; e = e.Next() {
			if e.Value.(string) == source {
				targetToSource = true
			}
		}

		return sourceToTarget && targetToSource
	} else {
		sourceToTarget := false
		for e := l.lists[i].Front(); e != nil; e = e.Next() {
			if e.Value.(string) == target {
				sourceToTarget = true
			}
		}

		return sourceToTarget
	}
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(1)
func (l *adjList) AddEdge(source, target string) error {
	if ok := l.HasEdge(source, target); ok {
		return ErrEdgeAlreadyExists(source, target)
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	i, ok := l.vertices[source]
	if !ok {
		return ErrVertexNotFound(source)
	}

	l.lists[i].PushBack(target)

	if l.undirected {
		j, ok := l.vertices[target]
		if !ok {
			return ErrVertexNotFound(target)
		}

		l.lists[j].PushBack(source)
	}

	l.e++

	return nil
}

// Time complexity: O(n), where n is number of vertices
//
// Space complexity: O(1)
func (l *adjList) DeleteEdge(source, target string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	i, ok := l.vertices[source]
	if !ok {
		return ErrVertexNotFound(source)
	}

	// Remove source -> target edge from source's list
	for e := l.lists[i].Front(); e != nil; e = e.Next() {
		if e.Value == target {
			l.lists[i].Remove(e)
		}
	}

	if l.undirected {
		j, ok := l.vertices[target]
		if !ok {
			return ErrVertexNotFound(target)
		}

		// Remove target -> sourcr edge from target's list
		for e := l.lists[j].Front(); e != nil; e = e.Next() {
			if e.Value == target {
				l.lists[j].Remove(e)
			}
		}
	}

	l.e--

	return nil
}

// Time complexity: O(v+e), where v is number of vertices, and e is number of edges
//
// Space complexity: O(v), where v is number of vertices
func (l *adjList) BFS(start string, callback func(vertex string)) error {
	l.lock.RLock()
	defer l.lock.RUnlock()

	visited := make(map[string]bool, l.Vertices())
	queue := queue.New()

	visited[start] = true
	queue.EnQueue(start)

	for queue.Len() > 0 {
		curr := queue.DeQueue().(string)
		callback(curr)

		currIdx, ok := l.vertices[curr]
		if !ok {
			return ErrVertexNotFound(curr)
		}

		for e := l.lists[currIdx].Front(); e != nil; e = e.Next() {
			vertex := e.Value.(string)

			if !visited[vertex] {
				visited[vertex] = true
				queue.EnQueue(vertex)
			}
		}
	}

	return nil
}

// Time complexity: O(v+e), where v is number of vertices, and e is number of edges
//
// Space complexity: O(v+e), where v is number of vertices, and e is number of edges
func (l *adjList) DFS(start string, callback func(vertex string)) error {
	visited := make(map[string]bool, l.Vertices())
	return l.dfs(start, callback, visited)
}

func (l *adjList) dfs(vertex string, callback func(vertex string), visited map[string]bool) error {
	l.lock.RLock()
	defer l.lock.RUnlock()

	visited[vertex] = true
	callback(vertex)

	i, ok := l.vertices[vertex]
	if !ok {
		return ErrVertexNotFound(vertex)
	}

	for e := l.lists[i].Front(); e != nil; e = e.Next() {
		vertex := e.Value.(string)

		if !visited[vertex] {
			if err := l.dfs(vertex, callback, visited); err != nil {
				return err
			}
		}
	}

	return nil
}

// Time complexity: O(v+e), where v is number of vertices, and e is number of edges
//
// Space complexity: O(v), where v is number of vertices
func (l *adjList) IsCyclic() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if l.undirected {
		panic(ErrCyclicCheckOnlyForDirected)
	}

	visited := make(map[string]bool, l.Vertices())
	recMap := make(map[string]bool, l.Vertices())

	for vertex := range l.vertices {
		if !visited[vertex] && l.isCyclicRec(vertex, visited, recMap) {
			return true
		}
	}

	return false
}

func (l *adjList) isCyclicRec(vertex string, visited, recMap map[string]bool) bool {
	if !visited[vertex] {
		// Mark the current node as visited
		// and part of recursion map
		visited[vertex] = true
		recMap[vertex] = true

		i := l.vertices[vertex]
		// TODO
		// if !ok {
		// return ErrVertexNotFound(vertex)
		// }

		for e := l.lists[i].Front(); e != nil; e = e.Next() {
			v := e.Value.(string)
			if !visited[v] && l.isCyclicRec(v, visited, recMap) {
				return true
			} else if vis, ok := recMap[v]; ok && vis {
				return true
			}
		}
	}

	// Remove the vertex from recursion stack
	recMap[vertex] = false
	return false
}

func (l *adjList) FindComponents() ([]*list.List, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	groupId := -1
	components := make([]*list.List, 0, l.v)
	visited := make(map[string]bool, l.v)

	grouping := func(vertex string) {
		if len(components)-1 < groupId {
			components = append(components, list.New())
		}
		components[groupId].PushBack(vertex)
	}

	for vertex := range l.vertices {
		if !visited[vertex] {
			groupId++
			if err := l.dfs(vertex, grouping, visited); err != nil {
				return nil, err
			}
		}
	}

	return components, nil
}

// Time complexity: O(n^2), where n is number of vertices
//
// Space complexity: O(1)
func (l *adjList) String() string {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if len(l.lists) == 0 {
		return "[]"
	}

	vertexIdx := l.vertexIdx()

	var buffer bytes.Buffer
	for i, list := range l.lists {
		buffer.WriteString(fmt.Sprintf("%v ", vertexIdx[i].Vertex))
		buffer.WriteString("[")
		for e := list.Front(); e != nil; e = e.Next() {
			if e == list.Front() {
				buffer.WriteString(fmt.Sprintf("%v", e.Value))
			} else {
				buffer.WriteString(fmt.Sprintf(", %v", e.Value))
			}
		}
		if list.Len() == 0 {
			buffer.WriteString("]\n")
		} else {
			buffer.WriteString("]\n")
		}
	}

	return buffer.String()
}
