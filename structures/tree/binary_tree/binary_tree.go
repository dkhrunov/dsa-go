package binarytree

import "github.com/dkhrunov/dsa-go/structures/queue"

type Node[T any] struct {
	Data        T
	Left, Right *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data}
}

// DFS
func TraversePreOrder[T any](root *Node[T], callback func(data T)) {
	if root == nil {
		return
	}
	callback(root.Data)
	TraversePreOrder(root.Left, callback)
	TraversePreOrder(root.Right, callback)
}

// DFS
func TraverseInOrder[T any](root *Node[T], callback func(data T)) {
	if root == nil {
		return
	}
	TraverseInOrder(root.Left, callback)
	callback(root.Data)
	TraverseInOrder(root.Right, callback)
}

// DFS
func TraversePostOrder[T any](root *Node[T], callback func(data T)) {
	if root == nil {
		return
	}
	TraversePostOrder(root.Left, callback)
	TraversePostOrder(root.Right, callback)
	callback(root.Data)
}

// BFS
func TraveseLevelOrder[T any](root *Node[T], callback func(data T)) {
	queue := queue.New()
	queue.EnQueue(root)

	for queue.Len() > 0 {
		temp := queue.DeQueue().(*Node[T])
		callback(temp.Data)

		if temp.Left != nil {
			queue.EnQueue(temp.Left)
		}
		if temp.Right != nil {
			queue.EnQueue(temp.Right)
		}
	}
}
