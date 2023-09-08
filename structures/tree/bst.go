package tree

import (
	"golang.org/x/exp/constraints"
)

type BSTree[T constraints.Ordered] struct {
	size int
	root *BinaryNode[T]
}

// NewBST creates a new binary search tree.
func NewBST[T constraints.Ordered]() BSTree[T] {
	return BSTree[T]{size: 0, root: nil}
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
func (tree *BSTree[T]) Insert(value T) {
	tree.root = tree.insert(tree.root, value)
	tree.size++
}

func (tree *BSTree[T]) insert(root *BinaryNode[T], value T) *BinaryNode[T] {
	if root == nil {
		return &BinaryNode[T]{
			value:  value,
			left:   nil,
			right:  nil,
			parent: nil,
		}
	}

	if value < root.value {
		root.left = tree.insert(root.left, value)
	} else if value > root.value {
		root.right = tree.insert(root.right, value)
	}

	return root
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
func (tree *BSTree[T]) Delete(value T) {
	tree.delete(tree.root, value)
	tree.size--
}

func (tree *BSTree[T]) delete(root *BinaryNode[T], value T) *BinaryNode[T] {
	if root == nil {
		return nil
	}

	if value < root.value {
		root.left = tree.delete(root.left, value)
	} else if value > root.value {
		root.right = tree.delete(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		root.value = LeftMostNode(root.right).value
		root.right = tree.delete(root.right, root.value)
	}

	return root
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

func (tree *BSTree[T]) search(root *BinaryNode[T], value T) *BinaryNode[T] {
	if root == nil {
		return nil
	}

	if value < root.value {
		return tree.search(root.left, value)
	} else if value > root.value {
		return tree.search(root.right, value)
	}

	return root
}
