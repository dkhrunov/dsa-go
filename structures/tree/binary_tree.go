package tree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dkhrunov/dsa-go/gmath"
	"github.com/dkhrunov/dsa-go/structures/queue"
	"github.com/dkhrunov/dsa-go/structures/stack"
)

const (
	SerializationStart     = "^"
	SerializationEnd       = "#"
	SerializationDelimiter = ","
)

// Binary tree node.
type BinaryNode[T any] struct {
	value               T
	left, right, parent *BinaryNode[T]
}

// TODO использовать данную структуру для BT вместе BTNode
// type BinaryTree[T any] struct {
// 	size int
// 	root *BinaryNode[T]
// }

// TODO add comparator to constructor and use it in comparisons

// NewBinaryTree creates a new binary tree.
func NewBinaryTree[T any](value T) *BinaryNode[T] {
	return &BinaryNode[T]{
		value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

// Value return the value of node.
func (n *BinaryNode[T]) Value() T {
	return n.value
}

// Left return the left child of node.
func (n *BinaryNode[T]) Left() *BinaryNode[T] {
	if n == nil {
		return nil
	}
	return n.left
}

// Right return the right child of node.
func (n *BinaryNode[T]) Right() *BinaryNode[T] {
	if n == nil {
		return nil
	}
	return n.right
}

// Parent return the parent of node.
func (n *BinaryNode[T]) Parent() *BinaryNode[T] {
	if n == nil {
		return nil
	}
	return n.parent
}

// InsertAfter inserts 'n' node after current node.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (n *BinaryNode[T]) InsertAfter(node *BinaryNode[T]) {
	InsertAfter(n, node)
}

// InsertBefore inserts 'new' node before passed 'node' in the first argument.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (n *BinaryNode[T]) InsertBefore(node *BinaryNode[T]) {
	InsertBefore(n, node)
}

// Delete removes the current node from the binary tree.
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (n *BinaryNode[T]) Delete() {
	Delete(n)
}

// TraversePreorder traverses in preorder of traversal of a binary tree.
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
func (n *BinaryNode[T]) TraversePreorder(callback func(node *BinaryNode[T]), empty func()) {
	TraversePreorder(n, callback, empty)
}

// TraverseInorder traverses in inorder traversal of traversal of a binary tree.
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
func (n *BinaryNode[T]) TraverseInorder(callback func(node *BinaryNode[T]), empty func()) {
	TraverseInorder(n, callback, empty)
}

// TraversePostorder traverses in postorder traversal of traversal of a binary tree.
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
func (n *BinaryNode[T]) TraversePostorder(callback func(node *BinaryNode[T]), empty func()) {
	TraverseInorder(n, callback, empty)
}

// TraverseLevelorder traverses in levelorder traversal of traversal of a binary tree.
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
func (n *BinaryNode[T]) TraverseLevelorder(callback func(node *BinaryNode[T]), empty func()) {
	TraverseLevelorder(n, callback, empty)
}

// SubtreeOf checks that s is a subtree of r.
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
func (n *BinaryNode[T]) SubtreeOf(root *BinaryNode[T]) bool {
	return IsSubtree(root, n)
}

// MaxDepth gets the maximum depth of the node.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(n).
func (n *BinaryNode[T]) MaxDepth() int {
	return MaxDepth(n)
}

// InsertAfter inserts 'new' node after passed 'node' in the first argument.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func InsertAfter[T any](node, new *BinaryNode[T]) {
	if node.right == nil {
		node.right = new
		new.parent = node
	} else {
		succ := InorderSuccessor(node)
		succ.left = new
		new.parent = succ
	}
}

// InsertBefore inserts 'new' node before passed 'node' in the first argument.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func InsertBefore[T any](node, new *BinaryNode[T]) {
	if node.left == nil {
		node.left = new
		new.parent = node
	} else {
		pred := InorderPredecessor(node)
		pred.right = new
		new.parent = pred
	}
}

// Delete removes given node from the binary tree.
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func Delete[T any](node *BinaryNode[T]) {
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

// TraversePreorder traverses in preorder of traversal of a binary tree.
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
func TraversePreorder[T any](root *BinaryNode[T], callback func(t *BinaryNode[T]), empty func()) {
	if root == nil {
		empty()
		return
	}
	callback(root)
	TraversePreorder(root.left, callback, empty)
	TraversePreorder(root.right, callback, empty)
}

// TraverseInorder traverses in inorder traversal of traversal of a binary tree.
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
func TraverseInorder[T any](root *BinaryNode[T], callback func(t *BinaryNode[T]), empty func()) {
	if root == nil {
		empty()
		return
	}
	TraverseInorder(root.left, callback, empty)
	callback(root)
	TraverseInorder(root.right, callback, empty)
}

// TraverseInorder traverses in inorder traversal of traversal of a binary tree.
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
func TraverseInorderI[T any](root *BinaryNode[T], callback func(t *BinaryNode[T]), empty func()) {
	stack := stack.New[*BinaryNode[T]]()
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

// TraversePostorder traverses in postorder traversal of traversal of a binary tree.
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
func TraversePostorder[T any](root *BinaryNode[T], callback func(t *BinaryNode[T]), empty func()) {
	if root == nil {
		empty()
		return
	}
	TraversePostorder(root.left, callback, empty)
	TraversePostorder(root.right, callback, empty)
	callback(root)
}

// TraverseLevelorder traverses in levelorder traversal of traversal of a binary tree.
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
func TraverseLevelorder[T any](root *BinaryNode[T], callback func(t *BinaryNode[T]), empty func()) {
	queue := queue.New()
	queue.EnQueue(root)

	for queue.Len() > 0 {
		curr := queue.DeQueue().(*BinaryNode[T])
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

// InorderSuccessor gets the node is the next node in Inorder traversal of the Binary Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where h is the height of the tree.
//
// Space complexity: O(1).
func InorderSuccessor[T any](n *BinaryNode[T]) *BinaryNode[T] {
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

// InorderSuccessorR gets the node is the next node in Inorder traversal of the Binary Tree.
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
func InorderSuccessorR[T any](root, target *BinaryNode[T]) *BinaryNode[T] {
	var (
		prev           *BinaryNode[T]
		succ           *BinaryNode[T]
		reverseInorder func(root *BinaryNode[T])
	)

	// reverse inorder traversal
	reverseInorder = func(root *BinaryNode[T]) {
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

// InorderSuccessorI gets the node is the next node in Inorder traversal of the Binary Tree.
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
func InorderSuccessorI[T any](root, target *BinaryNode[T]) *BinaryNode[T] {
	if target.right != nil {
		return LeftMostNode(target.right)
	}

	if reflect.DeepEqual(RightMostNode(root), target) {
		return nil
	}

	var (
		succ  *BinaryNode[T]
		found bool
	)

	stack := stack.New[*BinaryNode[T]]()
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

// InorderPredecessor gets the node is the prev node in Inorder traversal of the Binary Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where h is the height of the tree.
//
// Space complexity: O(1).
func InorderPredecessor[T any](node *BinaryNode[T]) *BinaryNode[T] {
	if node == nil {
		return nil
	}

	if node.left != nil {
		return RightMostNode(node.left)
	}

	p := node.parent
	for p != nil && reflect.DeepEqual(node, p.left) {
		node = p
		p = p.parent
	}
	return p
}

// IsIdentical checks that two root of binary tree are identical.
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
func IsIdentical[T any](x, y *BinaryNode[T]) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return reflect.DeepEqual(x.value, y.value) &&
		IsIdentical(x.left, y.left) &&
		IsIdentical(x.right, y.right) &&
		reflect.DeepEqual(x.parent, y.parent)
}

// IsSubtree checks that subRoot is a subtree of root.
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
func IsSubtree[T any](root, subRoot *BinaryNode[T]) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	return strings.Contains(Serialize(root), Serialize(subRoot))
}

// LeftMostNode gets the left most node of the tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func LeftMostNode[T any](node *BinaryNode[T]) *BinaryNode[T] {
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

// RightMostNode gets the right most node of the tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func RightMostNode[T any](node *BinaryNode[T]) *BinaryNode[T] {
	for node != nil && node.right != nil {
		node = node.right
	}
	return node
}

// MaxDepth gets the maximum depth of the node.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(n).
func MaxDepth[T any](node *BinaryNode[T]) int {
	if node == nil {
		return 0
	}

	lDepth := MaxDepth(node.left)
	rDepth := MaxDepth(node.right)
	return gmath.Max(lDepth, rDepth) + 1
}

// GetRoot gets the root of the binary tree for the given node.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func GetRoot[T any](node *BinaryNode[T]) *BinaryNode[T] {
	if node == nil {
		return nil
	}
	for node.parent != nil {
		node = node.parent
	}
	return node
}

// Serialize serializes binary tree.
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
func Serialize[T any](root *BinaryNode[T]) string {
	if root == nil {
		return ""
	}

	var sb strings.Builder

	TraversePreorder(
		root,
		func(t *BinaryNode[T]) {
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

// Deserialize deserializes the serialized before binary tree string and create a binary tree from the passed string.
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
func Deserialize(str string) *BinaryNode[string] {
	if len(str) == 0 {
		return &BinaryNode[string]{}
	}

	var dfs func(parent *BinaryNode[string]) *BinaryNode[string]

	tokens := strings.Split(str, SerializationDelimiter)
	dfs = func(parent *BinaryNode[string]) *BinaryNode[string] {
		token := strings.TrimPrefix(tokens[0], SerializationStart)
		tokens = tokens[1:]
		if token == SerializationEnd {
			return nil
		}
		node := &BinaryNode[string]{
			value:  token,
			parent: parent,
		}
		node.left = dfs(node)
		node.right = dfs(node)
		return node
	}

	return dfs(nil)
}
