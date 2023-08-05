package binarytree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dkhrunov/dsa-go/structures/queue"
	"github.com/dkhrunov/dsa-go/structures/stack"
)

const (
	SerializationStart     = "^"
	SerializationEnd       = "#"
	SerializationDelimiter = ","
)

// Binary tree node.
type Tree[T any] struct {
	value               T
	left, right, parent *Tree[T]
}

// Creates a new binary tree.
func New[T any](v T) *Tree[T] {
	return &Tree[T]{v, nil, nil, nil}
}

// Gets the value of the tree node.
func (t *Tree[T]) Value() T {
	return t.value
}

// Sets the tree node to a new value.
func (t *Tree[T]) SetValue(v T) {
	t.value = v
}

// Gets the left child of a tree node.
func (t *Tree[T]) Left() *Tree[T] {
	if t == nil {
		return nil
	}
	return t.left
}

// Gets the right child of a tree node.
func (t *Tree[T]) Right() *Tree[T] {
	if t == nil {
		return nil
	}
	return t.right
}

// Insert 'n' node after current node.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (t *Tree[T]) InsertAfter(n *Tree[T]) {
	InsertAfter(t, n)
}

// Insert 'new' node before passed 'node' in the first argument.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (t *Tree[T]) InsertBefore(n *Tree[T]) {
	InsertBefore(t, n)
}

// Delete the current node from the binary tree.
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (t *Tree[T]) Delete() {
	Delete(t)
}

// Preorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (t *Tree[T]) TraversePreorder(callback func(t *Tree[T]), empty func()) {
	TraversePreorder(t, callback, empty)
}

// Inorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (t *Tree[T]) TraverseInorder(callback func(t *Tree[T]), empty func()) {
	TraverseInorder(t, callback, empty)
}

// Postorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (t *Tree[T]) TraversePostorder(callback func(t *Tree[T]), empty func()) {
	TraverseInorder(t, callback, empty)
}

// Levelorder traversal of binary tree.
//
// BFS (Breadth First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(n).
func (t *Tree[T]) TraverseLevelorder(callback func(t *Tree[T]), empty func()) {
	TraverseLevelorder(t, callback, empty)
}

// Checks that s is a subtree of r.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n+m).
//
// Space complexity: O(n+m).
//
// This is because of the recursive calls in the serialize() function.
// The maximum depth of the recursive calls will be O(n) for 'root' and O(m) for 's'.
// So, the total space complexity is O(n+m)
func (t *Tree[T]) SubtreeOf(root *Tree[T]) bool {
	return IsSubtree(root, t)
}

// Insert 'new' node after passed 'node' in the first argument.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func InsertAfter[T any](node, new *Tree[T]) {
	if node.right == nil {
		node.right = new
		new.parent = node
	} else {
		succ := InorderSuccessor(node)
		succ.left = new
		new.parent = succ
	}
}

// Insert 'new' node before passed 'node' in the first argument.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func InsertBefore[T any](node, new *Tree[T]) {
	if node.left == nil {
		node.left = new
		new.parent = node
	} else {
		pred := InorderPredecessor(node)
		pred.right = new
		new.parent = pred
	}
}

// Delete given node from the binary tree.
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func Delete[T any](node *Tree[T]) {
	// [Base case] is leaf
	if node.left == nil && node.right == nil {
		parent := node.parent
		if reflect.DeepEqual(parent.left, node) {
			parent.left = nil
		} else {
			parent.right = nil
		}
		node.parent = nil
		return
	}

	predecessor := InorderPredecessor(node)
	if predecessor == nil {
		panic("Cannot delete left-most node for entire tree")
	}
	// swap values node.value <-> predecessor.value
	node.value, predecessor.value = predecessor.value, node.value
	Delete(predecessor)
}

// Preorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func TraversePreorder[T any](root *Tree[T], callback func(t *Tree[T]), empty func()) {
	if root == nil {
		empty()
		return
	}
	callback(root)
	TraversePreorder(root.left, callback, empty)
	TraversePreorder(root.right, callback, empty)
}

// Inorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func TraverseInorder[T any](root *Tree[T], callback func(t *Tree[T]), empty func()) {
	if root == nil {
		empty()
		return
	}
	TraverseInorder(root.left, callback, empty)
	callback(root)
	TraverseInorder(root.right, callback, empty)
}

// Inorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n*h), where 'n' is the number of nodes in the binary tree and 'h' is the height of tree and.
//
// Space complexity: O(h), where 'h' is the height of tree, since all left nodes are pushed onto the stack and then processed in turn.
func TraverseInorderI[T any](root *Tree[T], callback func(t *Tree[T]), empty func()) {
	stack := stack.New[*Tree[T]]()
	curr := root
	for curr != nil || stack.Len() > 0 {
		for curr != nil {
			stack.Push(curr)
			curr = curr.left
		}
		empty()
		curr, _ = stack.Pop()
		callback(curr)
		curr = curr.right
	}
	empty()
}

// Postorder traversal of binary tree.
//
// DFS (Deep First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func TraversePostorder[T any](root *Tree[T], callback func(t *Tree[T]), empty func()) {
	if root == nil {
		empty()
		return
	}
	TraversePostorder(root.left, callback, empty)
	TraversePostorder(root.right, callback, empty)
	callback(root)
}

// Levelorder traversal of binary tree.
//
// BFS (Breadth First Search) algorithm.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(n).
func TraverseLevelorder[T any](root *Tree[T], callback func(t *Tree[T]), empty func()) {
	queue := queue.New()
	queue.EnQueue(root)

	for queue.Len() > 0 {
		curr := queue.DeQueue().(*Tree[T])
		callback(curr)

		if curr.left != nil {
			queue.EnQueue(curr.left)
		} else {
			empty()
		}
		if curr.right != nil {
			queue.EnQueue(curr.right)
		} else {
			empty()
		}
	}
}

// Gets the node is the next node in Inorder traversal of the Binary Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where h is the height of the tree.
//
// Space complexity: O(1).
func InorderSuccessor[T any](n *Tree[T]) *Tree[T] {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return LeftMostNode(n.right)
	}

	p := n.parent
	for p != nil && reflect.DeepEqual(n, p.right) {
		n = p
		p = p.parent
	}
	return p
}

// Gets the node is the next node in Inorder traversal of the Binary Tree.
//
// Recursive and without parent pointer version of the InorderSuccessor.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n), where 'n' is the number of nodes in the binary tree.
//
// Space complexity: O(n), where 'n' is the number of nodes in the binary tree.
func InorderSuccessorR[T any](root, target *Tree[T]) *Tree[T] {
	var (
		prev           *Tree[T]
		succ           *Tree[T]
		reverseInorder func(root *Tree[T])
	)

	// reverse inorder traversal
	reverseInorder = func(root *Tree[T]) {
		if root == nil {
			return
		}

		reverseInorder(root.right)
		if reflect.DeepEqual(root, target) {
			// this case fot last node in tree
			// in inorder traversal i.e., rightmost node.
			if prev == nil {
				succ = nil
			} else {
				succ = prev
			}
			return
		}
		prev = root
		reverseInorder(root.left)
	}
	reverseInorder(root)

	return succ
}

// Gets the node is the next node in Inorder traversal of the Binary Tree.
//
// Iterative and without parent pointer vesrion of the InorderSuccessor.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n*h), where 'n' is the number of nodes in the binary tree and 'h' is the height of tree and.
//
// Space complexity: O(h), where 'h' is the height of tree, since all left nodes are pushed onto the stack and then processed in turn.
func InorderSuccessorI[T any](root, target *Tree[T]) *Tree[T] {
	if target.right != nil {
		return LeftMostNode(target.right)
	}

	if reflect.DeepEqual(RightMostNode(root), target) {
		return nil
	}

	var (
		stack stack.Stack[*Tree[T]]
		succ  *Tree[T]
		found bool
	)

	curr := root
	for curr != nil || stack.Len() > 0 {
		for curr != nil {
			stack.Push(curr)
			curr = curr.left
		}
		curr, _ = stack.Pop()
		if found {
			succ = curr
			break
		}
		if reflect.DeepEqual(curr, target) {
			found = true
		}
		curr = curr.right
	}

	return succ
}

// Gets the node is the prev node in Inorder traversal of the Binary Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where h is the height of the tree.
//
// Space complexity: O(1).
func InorderPredecessor[T any](n *Tree[T]) *Tree[T] {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return RightMostNode(n.left)
	}

	p := n.parent
	for p != nil && reflect.DeepEqual(n, p.left) {
		n = p
		p = p.parent
	}
	return p
}

// Checks that two root of binary tree are identical.
//
// Two trees ‘x’ and ‘y’ are identical if:
//
// - Val on their roots is the same or both roots are null
//
// - left subtree of ‘x’ is identical to the left sub-tree of ‘y’
//
// - right subtree of ‘x’ is identical to the right subtree of ‘y’
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h) - The recursive solution has O(h) memory complexity
// as it will consume memory on the stack up to the height of binary tree h.
func IsIdentical[T any](x, y *Tree[T]) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return reflect.DeepEqual(x.value, y.value) &&
		IsIdentical(x.left, y.left) &&
		IsIdentical(x.right, y.right) &&
		reflect.DeepEqual(x.parent, y.parent)
}

// Checks that subRoot is a subtree of root.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n+m).
//
// Space complexity: O(n+m) - This is because of the recursive calls in the serialize() function.
// The maximum depth of the recursive calls will be O(n) for 'root' and O(m) for 'subRoot'.
// So, the total space complexity is O(n+m).
func IsSubtree[T any](root, subRoot *Tree[T]) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	return strings.Contains(Serialize(root), Serialize(subRoot))
}

// Get the left most node of the tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func LeftMostNode[T any](node *Tree[T]) *Tree[T] {
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

// Get the right most node of the tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func RightMostNode[T any](node *Tree[T]) *Tree[T] {
	for node != nil && node.right != nil {
		node = node.right
	}
	return node
}

// Get the root of the binary tree for the given node.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func GetRoot[T any](node *Tree[T]) *Tree[T] {
	if node == nil {
		return nil
	}
	for node.parent != nil {
		node = node.parent
	}
	return node
}

// Serialize binary tree.
// The function Serialize() is similar to the preorder traversal of the tree.
//
// Special characters "^" and "#":
// Each node is preceded by a "^" to signify the beginning.
// If the node has no left or right child, a "#" sign is added.
//
// --------------------------------------------------
//
// Complexity:
// Complexity same as DFS algorithm.
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func Serialize[T any](root *Tree[T]) string {
	var sb strings.Builder

	TraversePreorder(
		root,
		func(t *Tree[T]) {
			sb.WriteString(SerializationStart)
			sb.WriteString(fmt.Sprintf("%v", t.value))
			sb.WriteString(SerializationDelimiter)
		},
		func() {
			sb.WriteString(SerializationEnd)
			sb.WriteString(SerializationDelimiter)
		},
	)

	return sb.String()
}

// Deserialize the serialized binary tree string and create a binary tree from the passed string.
//
// --------------------------------------------------
//
// Complexity:
// Complexity same as DFS algorithm.
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func Deserialize(str string) *Tree[string] {
	var dfs func(parent *Tree[string]) *Tree[string]

	tokens := strings.Split(str, SerializationDelimiter)
	dfs = func(parent *Tree[string]) *Tree[string] {
		token := strings.TrimPrefix(tokens[0], SerializationStart)
		tokens = tokens[1:]
		if token == SerializationEnd {
			return nil
		}
		node := &Tree[string]{
			value:  token,
			parent: parent,
		}
		node.left = dfs(node)
		node.right = dfs(node)
		return node
	}

	return dfs(nil)
}
