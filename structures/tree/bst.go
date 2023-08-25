package tree

import (
	"golang.org/x/exp/constraints"
)

type BST[T constraints.Ordered] struct {
	size int
	root *BinaryNode[T]
}

// Creates a new binary search tree.
func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{size: 0, root: nil}
}

// Gets the size of the BST.
func (tree *BST[T]) Size() int {
	return tree.size
}

// Gets the root of the BST.
func (tree *BST[T]) Root() *BinaryNode[T] {
	return tree.root
}

// Min gets the minimum value.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (tree *BST[T]) Min() T {
	return LeftMostNode(tree.root).value
}

// Min gets the maximum value.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(h), where the 'h' is height of tree.
//
// Space complexity: O(1).
func (tree *BST[T]) Max() T {
	return RightMostNode(tree.root).value
}

// Insert inserts a new value to BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BST[T]) Insert(value T) {
	tree.root = insertRec(tree.root, value)
	tree.size++
}

// Delete removes value from BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BST[T]) Delete(value T) {
	deleteRec(tree.root, value)
	tree.size--
}

// Search search given value in BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func (tree *BST[T]) Search(value T) *BinaryNode[T] {
	return searchRec(tree.root, value)
}

// Insert inserts a new value to BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func insertRec[T constraints.Ordered](root *BinaryNode[T], value T) *BinaryNode[T] {
	if root == nil {
		return &BinaryNode[T]{value, nil, nil, nil}
	}

	if value < root.value {
		root.left = insertRec(root.left, value)
	} else if value > root.value {
		root.right = insertRec(root.right, value)
	}

	return root
}

// Delete removes value from BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func deleteRec[T constraints.Ordered](root *BinaryNode[T], value T) *BinaryNode[T] {
	if root == nil {
		return nil
	}

	if value < root.value {
		root.left = deleteRec(root.left, value)
	} else if value > root.value {
		root.right = deleteRec(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		root.value = LeftMostNode(root.right).value
		root.right = deleteRec(root.right, root.value)
	}

	return root
}

// Search search given value in BST.
//
// --------------------------------------------------
//
// Complexity:
//
// Time complexity: O(log n).
//
// Space complexity: O(h), where 'h' is the height of tree, if we do consider the stack size for function calls.
// Otherwise, the space complexity of inorder traversal is O(1).
func searchRec[T constraints.Ordered](root *BinaryNode[T], value T) *BinaryNode[T] {
	if root == nil {
		return nil
	}

	if value < root.value {
		return searchRec(root.left, value)
	} else if value > root.value {
		return searchRec(root.right, value)
	}

	return root
}
