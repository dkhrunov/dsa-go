package tree

import (
	"fmt"
	"strings"

	"github.com/dkhrunov/dsa-go/gmath"
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

type AVLNode[T constraints.Ordered] struct {
	value               T
	bf                  int
	height              int
	left, right, parent *AVLNode[T]
}

// TODO add comparator to constructor and use it in comparisons

// NewAVLNode creates a new AVL node.
func NewAVLNode[T constraints.Ordered](value T) *AVLNode[T] {
	return &AVLNode[T]{
		value:  value,
		bf:     0,
		height: 0,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

// Value return the value of node.
func (node *AVLNode[T]) Value() T {
	return node.value
}

// Left return the left child of node.
func (node *AVLNode[T]) Left() *AVLNode[T] {
	if node == nil {
		return nil
	}
	return node.left
}

// Right return the right child of node.
func (node *AVLNode[T]) Right() *AVLNode[T] {
	if node == nil {
		return nil
	}
	return node.right
}

// Parent return the parent of node.
func (node *AVLNode[T]) Parent() *AVLNode[T] {
	if node == nil {
		return nil
	}
	return node.parent
}

type AVLTree[T constraints.Ordered] struct {
	// Tracks the number of nodes inside the tree.
	size int
	// The root node of the AVL tree.
	root *AVLNode[T]
}

// NewAVLTree creates a new AVL tree.
func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{size: 0, root: nil}
}

// Root returns the root of the AVLTree.
func (tree *AVLTree[T]) Root() *AVLNode[T] {
	return tree.root
}

// Height the height of a rooted tree is the number of edges between the tree's
// root and its furthest leaf. This means that a tree containing a single
// node has a height of 0.
func (tree *AVLTree[T]) Height() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.height
}

// Size returns the number of nodes in the tree.
func (tree *AVLTree[T]) Size() int {
	return tree.size
}

// Insert inserts a new value to AVLTree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *AVLTree[T]) Insert(value T) bool {
	if tree.contains(tree.Root(), value) {
		return false
	}

	tree.root = tree.insert(tree.root, nil, value)
	tree.size++
	return true
}

// insert inserts a value inside the AVL tree.
func (tree *AVLTree[T]) insert(node, parent *AVLNode[T], value T) *AVLNode[T] {
	// Base case.
	if node == nil {
		newNode := &AVLNode[T]{
			value:  value,
			parent: parent,
		}
		return newNode
	}

	// Insert node in left subtree.
	if value < node.value {
		node.left = tree.insert(node.left, node, value)

		// Insert node in right subtree.
	} else if value > node.value {
		node.right = tree.insert(node.right, node, value)
	}

	// Update balance factor and height values.
	tree.update(node)

	// Re-balance tree.
	return tree.balance(node)
}

// Delete removes value from AVL Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *AVLTree[T]) Delete(value T) bool {
	if !tree.contains(tree.Root(), value) {
		return false
	}

	tree.root = tree.delete(tree.root, value)
	tree.size--
	return true
}

// delete deletes a value from the AVL tree.
func (tree *AVLTree[T]) delete(node *AVLNode[T], value T) *AVLNode[T] {
	if node == nil {
		return nil
	}

	// Dig into left subtree, the value we're looking
	// for is smaller than the current value.
	if value < node.value {
		node.left = tree.delete(node.left, value)

		// Dig into right subtree, the value we're looking
		// for is greater than the current value.
	} else if value > node.value {
		node.right = tree.delete(node.right, value)

		// Found the node we wish to remove.
	} else {
		// This is the case with only a right subtree or no subtree at all.
		// In this situation just swap the node we wish to remove
		// with its right child.
		if node.left == nil {
			return node.right
		}

		// This is the case with only a left subtree or
		// no subtree at all. In this situation just
		// swap the node we wish to remove with its left child.
		if node.right == nil {
			return node.left
		}

		// When removing a node from a binary tree with two links the
		// successor of the node being removed can either be the largest
		// value in the left subtree or the smallest value in the right
		// subtree. As a heuristic, I will remove from the subtree with
		// the greatest height in hopes that this may help with balancing.

		// Choose to remove from left subtree
		if node.left.height > node.right.height {
			// Swap the value of the successor into the node.
			successorValue := tree.findMax(node.left)
			node.value = successorValue

			// Find the largest node in the left subtree.
			node.left = tree.delete(node.left, successorValue)
		} else {
			// Swap the value of the successor into the node.
			successorValue := tree.findMin(node.right)
			node.value = successorValue

			// Go into the right subtree and remove the leftmost node we
			// found and swapped data with. This prevents us from having
			// two nodes in our tree with the same value.
			node.right = tree.delete(node.right, successorValue)
		}

		// or simply do that:
		// node.value = tree.findMin(node.right)
		// node.right = tree.delete(node.right, node.value)
	}

	// Update balance factor and height values.
	tree.update(node)

	// Re-balance tree.
	return tree.balance(node)
}

// findMax is a helper method to find the rightmost node (which has the largest value)
func (tree *AVLTree[T]) findMax(node *AVLNode[T]) T {
	for node != nil && node.right != nil {
		node = node.right
	}
	return node.value
}

// findMin is a helper method to find the leftmost node (which has the smallest value)
func (tree *AVLTree[T]) findMin(node *AVLNode[T]) T {
	for node != nil && node.left != nil {
		node = node.left
	}
	return node.value
}

// update update a node's height and balance factor.
func (tree *AVLTree[T]) update(node *AVLNode[T]) {
	leftNodeHeight := -1
	if node.left != nil {
		leftNodeHeight = node.left.height
	}

	rightNodeHeight := -1
	if node.right != nil {
		rightNodeHeight = node.right.height
	}

	// Update this node's height
	node.height = 1 + gmath.Max(leftNodeHeight, rightNodeHeight)

	// Update balance factor
	node.bf = rightNodeHeight - leftNodeHeight
}

// balance re-balance a node if its balance factor is +2 or -2.
func (tree *AVLTree[T]) balance(node *AVLNode[T]) *AVLNode[T] {
	// Left heavy subtree
	if node.bf == -2 {
		// Left-Left case
		// 			[ A ] bf = -2
		// 			/
		// 		[ B ]  bf = -1
		// 		/
		// 	[ C ]
		if node.left.bf <= 0 {
			return tree.rightRotate(node)
		}

		// Left-Right case
		// 			[ A ] bf = -2
		// 			/
		// 		[ B ]  bf = +1
		// 				\
		// 				[ C ]
		node.left = tree.leftRotate(node.left)
		return tree.rightRotate(node)
	}

	// Right heavy subtree
	if node.bf == 2 {
		// Right-Right case
		// 	[ A ] bf = +2
		// 			\
		// 			[ B ] bf = +1
		// 					\
		// 					[ C ]
		if node.right.bf >= 0 {
			return tree.leftRotate(node)
		}

		// Right-Left case
		// 	[ A ] bf = +2
		// 			\
		// 			[ B ] bf = -1
		// 			/
		// 	[ C ]
		node.right = tree.rightRotate(node.right)
		return tree.leftRotate(node)
	}

	// Node either has a balance factor of 0, +1 or -1 which is fine
	return node
}

// TODO test
// leftRotate rotate the tree to the left and return new root
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(1).
//
// Space complexity: O(1).
func (tree *AVLTree[T]) leftRotate(n *AVLNode[T]) *AVLNode[T] {
	// Legend:

	// n - root to be rotated
	// p - it's parent node of the n
	// b - right child of n
	// c - left subtree of n
	// d - left subtree of lc
	// e - right subtree of lc

	// Before rotate:

	// 				[ P ]
	//   				|
	//				[ N ]
	// 			/      \
	// 	  /	\	 	  [ B ]
	// 	 / C \	 /     \
	// 				 / \     / \
	// 			  / D \   / E \

	// After rotate:

	//		 				[ P ]
	//		   				|
	//						[ B ]
	//		 			/      \
	//		 	 [ N ]		 / \
	//		 	/	    \   / E \
	//		/ \     / \
	//	 / C \   / D \

	if n == nil {
		return nil
	}

	// temp
	p := n.parent
	b := n.right

	// change n.right
	n.right = b.left
	if b.left != nil {
		b.left.parent = n
	}

	// change b.left
	b.left = n
	n.parent = b
	b.parent = p

	// change parent down link
	if p != nil {
		if p.left == n {
			p.left = b
		} else {
			p.right = b
		}
	}

	tree.update(n)
	tree.update(b)

	return b
}

// TODO test
// rightRotate rotate the tree to the right and return new root
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(1).
//
// Space complexity: O(1).
func (tree *AVLTree[T]) rightRotate(n *AVLNode[T]) *AVLNode[T] {
	// Legend:

	// n - root to be rotated
	// p - it's parent node of the n
	// b - left child of n
	// c - right subtree of n
	// d - left subtree of lc
	// e - right subtree of lc

	// Before rotate:

	//		 				[ P ]
	//		   				|
	//						[ N ]
	//		 			/      \
	//		 	 [ B ]		 / \
	//		 	/	    \   / C \
	//		/ \     / \
	//	 / D \   / E \

	// After rotate:

	// 				[ P ]
	//   				|
	//				[ B ]
	// 			/      \
	// 	  /	\	 	  [ N ]
	// 	 / D \	 /     \
	// 				 / \     / \
	// 			  / E \   / C \

	if n == nil {
		return nil
	}

	// temp
	p := n.parent
	b := n.left

	// change n.left
	n.left = b.right
	if b.right != nil {
		b.right.parent = n
	}

	// change b.right
	b.right = n
	n.parent = b
	b.parent = p

	// change parent down link
	if p != nil {
		if p.left == n {
			p.left = b
		} else {
			p.right = b
		}
	}

	tree.update(n)
	tree.update(b)

	return b
}

// Search search given value in AVL Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *AVLTree[T]) Search(value T) *AVLNode[T] {
	return tree.search(tree.root, value)
}

func (tree *AVLTree[T]) search(node *AVLNode[T], value T) *AVLNode[T] {
	if node == nil {
		return nil
	}

	if value < node.value {
		return tree.search(node.left, value)
	} else if value > node.value {
		return tree.search(node.right, value)
	}

	return node
}

// Contains checks for the presence of a value in the AVL Tree.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *AVLTree[T]) Contains(value T) bool {
	return tree.contains(tree.root, value)
}

func (tree *AVLTree[T]) contains(node *AVLNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value < node.value {
		return tree.contains(node.left, value)
	} else if value > node.value {
		return tree.contains(node.right, value)
	}

	return true
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
func (tree *AVLTree[T]) TraversePreorder(callback func(t *AVLNode[T]), empty func()) {
	tree.traversePreorder(tree.root, callback, empty)
}

func (tree *AVLTree[T]) traversePreorder(node *AVLNode[T], callback func(node *AVLNode[T]), empty func()) {
	if node == nil {
		empty()
		return
	}
	callback(node)
	tree.traversePreorder(node.left, callback, empty)
	tree.traversePreorder(node.right, callback, empty)
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
func (tree *AVLTree[T]) Serialize() string {
	if tree.root == nil {
		return ""
	}

	var sb strings.Builder

	tree.TraversePreorder(
		func(t *AVLNode[T]) {
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
func (tree *AVLTree[T]) Deserialize(str string) *AVLNode[string] {
	if len(str) == 0 {
		return NewAVLNode(utils.Zero[string]())
	}

	var dfs func(parent *AVLNode[string]) *AVLNode[string]

	tokens := strings.Split(str, SerializationDelimiter)
	dfs = func(parent *AVLNode[string]) *AVLNode[string] {
		token := strings.TrimPrefix(tokens[0], SerializationStart)
		tokens = tokens[1:]
		if token == SerializationEnd {
			return nil
		}
		node := &AVLNode[string]{
			value:  token,
			parent: parent,
		}
		node.left = dfs(node)
		node.right = dfs(node)
		return node
	}

	return dfs(nil)
}
