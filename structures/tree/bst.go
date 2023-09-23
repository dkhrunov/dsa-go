package tree

import (
	"golang.org/x/exp/constraints"
)

type BSTree[T constraints.Ordered] struct {
	size int
	root *BinaryNode[T]
}

// NewBST creates a new binary search tree.
func NewBST[T constraints.Ordered]() *BSTree[T] {
	return &BSTree[T]{size: 0, root: nil}
}

// Size returns the size of the BST.
func (tree *BSTree[T]) Size() int {
	return tree.size
}

// Root returns the root of the BST.
func (tree *BSTree[T]) Root() *BinaryNode[T] {
	return tree.root
}

// Min gets the minimum value.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(1).
func (tree *BSTree[T]) Min() T {
	return LeftMostNode(tree.root).value
}

// Min gets the maximum value.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(1).
func (tree *BSTree[T]) Max() T {
	return RightMostNode(tree.root).value
}

// Insert inserts a new value to BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BSTree[T]) Insert(value T) bool {
	if tree.contains(tree.Root(), value) {
		return false
	}

	tree.root = tree.insert(tree.root, nil, value)
	tree.size++
	return true
}

func (tree *BSTree[T]) insert(node, parent *BinaryNode[T], value T) *BinaryNode[T] {
	if node == nil {
		return &BinaryNode[T]{
			value:  value,
			left:   nil,
			right:  nil,
			parent: parent,
		}
	}

	if value < node.value {
		node.left = tree.insert(node.left, node, value)
	} else if value > node.value {
		node.right = tree.insert(node.right, node, value)
	}

	return node
}

// Delete removes value from BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BSTree[T]) Delete(value T) bool {
	if !tree.contains(tree.Root(), value) {
		return false
	}

	tree.delete(tree.root, value)
	tree.size--
	return true
}

// delete deletes a value from the AVL tree.
func (tree *BSTree[T]) delete(node *BinaryNode[T], value T) *BinaryNode[T] {
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

		node.value = LeftMostNode(node.right).value
		node.right = tree.delete(node.right, node.value)
	}

	return node
}

// Search search given value in BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BSTree[T]) Search(value T) *BinaryNode[T] {
	return tree.search(tree.root, value)
}

func (tree *BSTree[T]) search(node *BinaryNode[T], value T) *BinaryNode[T] {
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

// Contains checks for the presence of a value in the BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BSTree[T]) Contains(value T) bool {
	return tree.contains(tree.root, value)
}

func (tree *BSTree[T]) contains(node *BinaryNode[T], value T) bool {
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
