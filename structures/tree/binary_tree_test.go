package tree

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	t.Parallel()
	if got := NewBinaryTree(0); got == nil {
		t.Errorf("NewBinaryTree() can't create new binary tree node struct")
	}
}

func TestValue(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree(t)
	node := &BinaryNode[int]{}
	type args struct {
		n *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should get the node value", args{bt}, 1},
		{"should get the nested node value", args{bt.Left().Left()}, 4},
		{"should get default zero value for empty node", args{node}, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.args.n.Value(); got != tt.want {
				t.Errorf("Value(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree(t)
	node := &BinaryNode[int]{}
	type args struct {
		n *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should get the left child of the node", args{bt}, bt.left},
		{"should get the nil if left child does not exist ", args{node}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.args.n.Left(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree(t)
	node := &BinaryNode[int]{}
	type args struct {
		n *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should get the right child of the node", args{bt}, bt.right},
		{"should get the nil if right child does not exist ", args{node}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.args.n.Right(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {
	t.Parallel()
	t.Run("should be inserted after leaf node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree(t)
		after := LeftMostNode(bt)
		new := NewBinaryTree(99)
		InsertAfter(after, new)
		if !reflect.DeepEqual(after.right, new) {
			t.Errorf("InsertAfter(...args) should be inserted as right child of %v node", after)
		}
		if !reflect.DeepEqual(new.parent, after) {
			t.Error("InsertAfter(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("should be inserted after root node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree(t)
		after := LeftMostNode(bt.right)
		new := NewBinaryTree(99)
		InsertAfter(bt, new)
		if !reflect.DeepEqual(LeftMostNode(bt.right), new) {
			t.Errorf("InsertAfter(...args) should be inserted as left child of %v node", after)
		}
		if !reflect.DeepEqual(new.parent, after) {
			t.Error("InsertAfter(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("should be inserted after internal mode", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree(t)
		after := LeftMostNode(bt).parent
		parent := LeftMostNode(after.right)
		new := NewBinaryTree(99)
		InsertAfter(after, new)
		if !reflect.DeepEqual(LeftMostNode(after.right), new) {
			t.Errorf("InsertAfter(...args) should be inserted as left child of %v node", parent)
		}
		if !reflect.DeepEqual(new.parent, parent) {
			t.Error("InsertAfter(...args) should set correct parent pointer for inserted node")
		}
	})
}

func TestInsertBefore(t *testing.T) {
	t.Parallel()
	t.Run("should be inserted before leaf node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree(t)
		before := LeftMostNode(bt)
		new := NewBinaryTree(99)
		InsertBefore(before, new)
		if !reflect.DeepEqual(before.left, new) {
			t.Errorf("InsertBefore(...args) should be inserted as left child of %v node", before)
		}
		if !reflect.DeepEqual(new.parent, before) {
			t.Error("InsertBefore(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("should be inserted before root node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree(t)
		before := RightMostNode(bt.left)
		new := NewBinaryTree(99)
		InsertBefore(bt, new)
		if !reflect.DeepEqual(RightMostNode(bt.left), new) {
			t.Errorf("InsertBefore(...args) should be inserted as right child of %v node", before)
		}
		if !reflect.DeepEqual(new.parent, before) {
			t.Error("InsertBefore(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("should be inserted before internal node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree(t)
		before := RightMostNode(bt).parent
		parent := RightMostNode(before.left)
		new := NewBinaryTree(99)
		InsertBefore(before, new)
		if !reflect.DeepEqual(RightMostNode(before.left), new) {
			t.Errorf("InsertBefore(...args) should be inserted as right child of %v node", parent)
		}
		if !reflect.DeepEqual(new.parent, parent) {
			t.Error("InsertBefore(...args) should set correct parent pointer for inserted node")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Delete root node", args{newFullBinaryTree(t)}, "^5,^2,^4,^8,#,#,^9,#,#,#,^3,^6,#,#,^7,^10,#,#,^11,#,#,"},
		{"Delete leaf node", args{newFullBinaryTree(t).Left().Right()}, "^1,^2,^4,^8,#,#,^9,#,#,#,^3,^6,#,#,^7,^10,#,#,^11,#,#,"},
		{"Delete node without left child", args{
			func() *BinaryNode[int] {
				r := newFullBinaryTree(t)
				// Delete from left subtree from right subtree
				r.right.left.parent = nil
				r.right.left = nil
				return r.right
			}(),
		}, "^5,^2,^4,^8,#,#,^9,#,#,#,^1,#,^7,^10,#,#,^11,#,#,"},
		{"Delete node without right child", args{
			func() *BinaryNode[int] {
				r := newFullBinaryTree(t)
				// Delete from left subtree from right subtree
				r.left.left.right.parent = nil
				r.left.left.right = nil
				return r.left.left
			}(),
		}, "^1,^2,^8,#,#,^5,#,#,^3,^6,#,#,^7,^10,#,#,^11,#,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			root := GetRoot(tt.args.node)
			Delete(tt.args.node)
			if got := Serialize(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete(node) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}

	t.Run("Cannot delete left-most node because this node has no predecessor", func(t *testing.T) {
		t.Parallel()
		defer func() { recover() }()
		r := newFullBinaryTree(t)
		// Delete from left subtree from right subtree
		r.left.left.left.parent = nil
		r.left.left.left = nil
		Delete(r.left.left)
		t.Errorf("should have panicked")
	})
}

func TestTraversePreorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Preorder traverse of full binary tree", args{newFullBinaryTree(t)}, "1,2,4,8,#,#,9,#,#,5,#,#,3,6,#,#,7,10,#,#,11,#,#,"},
		{"Preorder traverse of complete binary tree", args{newCompleteBinaryTree(t)}, "1,2,4,8,#,#,9,#,#,5,10,#,#,11,#,#,3,6,#,#,7,#,#,"},
		{"Preorder traverse of prefect binary tree", args{newPerfectBinaryTree(t)}, "1,2,4,#,#,5,#,#,3,6,#,#,7,#,#,"},
		{"Preorder traverse of degenerate binary tree", args{newDegenerateBinaryTree(t)}, "1,2,#,3,#,4,#,#,#,"},
		{"Preorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, "1,2,3,4,#,#,#,#,#,"},
		{"Preorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, "1,#,2,#,3,#,4,#,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraversePreorder(tt.args.root, traverseCallbackFactory[int](t, sb), traverseEmptyFactory(t, sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraversePreorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraverseInorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Inorder traverse of full binary tree", args{newFullBinaryTree(t)}, "#,8,#,4,#,9,#,2,#,5,#,1,#,6,#,3,#,10,#,7,#,11,#,"},
		{"Inorder traverse of complete binary tree", args{newCompleteBinaryTree(t)}, "#,8,#,4,#,9,#,2,#,10,#,5,#,11,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of prefect binary tree", args{newPerfectBinaryTree(t)}, "#,4,#,2,#,5,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of degenerate binary tree", args{newDegenerateBinaryTree(t)}, "#,2,#,3,#,4,#,1,#,"},
		{"Inorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, "#,4,#,3,#,2,#,1,#,"},
		{"Inorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, "#,1,#,2,#,3,#,4,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraverseInorder(tt.args.root, traverseCallbackFactory[int](t, sb), traverseEmptyFactory(t, sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseInorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraverseInorderI(t *testing.T) {
	t.Parallel()
	type args struct {
		root *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Inorder traverse of full binary tree", args{newFullBinaryTree(t)}, "#,8,#,4,#,9,#,2,#,5,#,1,#,6,#,3,#,10,#,7,#,11,#,"},
		{"Inorder traverse of complete binary tree", args{newCompleteBinaryTree(t)}, "#,8,#,4,#,9,#,2,#,10,#,5,#,11,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of prefect binary tree", args{newPerfectBinaryTree(t)}, "#,4,#,2,#,5,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of degenerate binary tree", args{newDegenerateBinaryTree(t)}, "#,2,#,3,#,4,#,1,#,"},
		{"Inorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, "#,4,#,3,#,2,#,1,#,"},
		{"Inorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, "#,1,#,2,#,3,#,4,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraverseInorderI(tt.args.root, traverseCallbackFactory[int](t, sb), traverseEmptyFactory(t, sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseInorderI(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraversePostorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Postorder traverse of full binary tree", args{newFullBinaryTree(t)}, "#,#,8,#,#,9,4,#,#,5,2,#,#,6,#,#,10,#,#,11,7,3,1,"},
		{"Postorder traverse of complete binary tree", args{newCompleteBinaryTree(t)}, "#,#,8,#,#,9,4,#,#,10,#,#,11,5,2,#,#,6,#,#,7,3,1,"},
		{"Postorder traverse of prefect binary tree", args{newPerfectBinaryTree(t)}, "#,#,4,#,#,5,2,#,#,6,#,#,7,3,1,"},
		{"Postorder traverse of degenerate binary tree", args{newDegenerateBinaryTree(t)}, "#,#,#,#,4,3,2,#,1,"},
		{"Postorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, "#,#,4,#,3,#,2,#,1,"},
		{"Postorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, "#,#,#,#,#,4,3,2,1,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraversePostorder(tt.args.root, traverseCallbackFactory[int](t, sb), traverseEmptyFactory(t, sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraversePostorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraverseLevelorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TraverseLevelorder traverse of full binary tree", args{newFullBinaryTree(t)}, "1,2,3,4,5,#,#,6,#,#,7,8,#,#,9,#,#,10,#,#,11,#,#,"},
		{"TraverseLevelorder traverse of complete binary tree", args{newCompleteBinaryTree(t)}, "1,2,3,4,5,6,#,#,7,#,#,8,#,#,9,#,#,10,#,#,11,#,#,"},
		{"TraverseLevelorder traverse of prefect binary tree", args{newPerfectBinaryTree(t)}, "1,2,3,4,#,#,5,#,#,6,#,#,7,#,#,"},
		{"TraverseLevelorder traverse of degenerate binary tree", args{newDegenerateBinaryTree(t)}, "1,#,2,#,3,#,4,#,#,"},
		{"TraverseLevelorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, "1,#,2,#,3,#,4,#,#,"},
		{"TraverseLevelorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, "1,#,2,#,3,#,4,#,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraverseLevelorder(tt.args.root, traverseCallbackFactory[int](t, sb), traverseEmptyFactory(t, sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseLevelorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestInorderSuccessor(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree(t)
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should be the left most node of the target's right subtree", args{bt}, bt.Right().Left()},
		{"should be the parent node of the target", args{bt.Left().Left()}, bt.Left()},
		{"should be the grandparent node of the target", args{bt.Left().Right()}, bt},
		{"should be the nil, if target last node in tree", args{bt.Right().Right()}, nil},
		{"should be the right node of the target's right subtree", args{bt.Right()}, bt.Right().Right()},
		{"should be a nil of nil root", args{nil}, nil},
		{"should be the nil of right-most node in tree", args{RightMostNode(bt)}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := InorderSuccessor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderSuccessor(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderSuccessorR(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree(t)
	type args struct {
		root   *BinaryNode[int]
		target *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should be the left most node of the target's right subtree", args{bt, bt}, bt.Right().Left()},
		{"should be the parent node of the target", args{bt, bt.Left().Left()}, bt.Left()},
		{"should be the grandparent node of the target", args{bt, bt.Left().Right()}, bt},
		{"should be the nil, if target last node in tree", args{bt, bt.Right().Right()}, nil},
		{"should be the right node of the target's right subtree", args{bt, bt.Right()}, bt.Right().Right()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := InorderSuccessorR(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderSuccessorR(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderSuccessorI(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree(t)
	type args struct {
		root   *BinaryNode[int]
		target *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should be the left-most node of the target's right subtree", args{bt, bt}, bt.Right().Left()},
		{"should be the parent node of the target", args{bt, bt.Left().Left()}, bt.Left()},
		{"should be the grandparent of the target", args{bt, bt.Left().Right()}, bt},
		{"should be null if the target node is the last node in the inorder traversal", args{bt, bt.Right().Right()}, nil},
		{"should be the right node of the target's right subtree", args{bt, bt.Right()}, bt.Right().Right()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := InorderSuccessorI(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderSuccessorI(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderPredecessor(t *testing.T) {
	t.Parallel()
	bt := newFullBinaryTree(t)
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should be the right-most node of the target's left subtree", args{bt}, bt.Left().Right()},
		{"should be the parent node of the target #1", args{bt.Left().Left().Right()}, bt.Left().Left()},
		{"should be root node for the left-most node in the right subtree", args{bt.Right().Left()}, bt},
		{"should be null if the target node is the first node in the inorder traversal", args{LeftMostNode(bt)}, nil},
		{"should be a nil of nil root", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := InorderPredecessor(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderPredecessor(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIdentical(t *testing.T) {
	t.Parallel()
	tree1X := newFullBinaryTree(t)
	tree1Y := newFullBinaryTree(t)

	tree2X := newCompleteBinaryTree(t)
	tree2Y := newCompleteBinaryTree(t)
	tree2Y.Right().Right().value = 99

	type args struct {
		x *BinaryNode[int]
		y *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should be identical", args{tree1X, tree1Y}, true},
		{"should not be identical", args{tree2X, tree2Y}, false},
		{"Two nil roots should be identical", args{nil, nil}, true},
		{"If one is nil and the other is not, then it should not be identical", args{newFullBinaryTree(t), nil}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsIdentical(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsIdentical(x , y) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubtree(t *testing.T) {
	t.Parallel()
	tree1 := newFullBinaryTree(t)

	subtree1 := tree1.left

	tree2 := newFullBinaryTree(t)

	subtree2 := NewBinaryTree(2)
	subtree2.left = NewBinaryTree(4)
	subtree2.left.left = NewBinaryTree(8)
	subtree2.right = NewBinaryTree(5)

	tree3 := NewBinaryTree(12)

	subtree3 := NewBinaryTree(2)

	tree4 := NewBinaryTree(1)
	tree4.left = NewBinaryTree(2)
	tree4.left.left = NewBinaryTree(4)
	tree4.right = NewBinaryTree(3)

	subtree4 := NewBinaryTree(1)
	subtree4.left = NewBinaryTree(2)
	subtree4.right = NewBinaryTree(3)

	tree5 := newFullBinaryTree(t)

	var subtree5 *BinaryNode[int]

	tree6 := NewBinaryTree(1)

	subtree6 := NewBinaryTree(1)

	type args struct {
		root    *BinaryNode[int]
		subRoot *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should be a subtree", args{tree1, subtree1}, true},
		{"should not be a subtree", args{tree2, subtree2}, false},
		{"should not be a subtree", args{tree3, subtree3}, false},
		{"should not be a subtree", args{tree4, subtree4}, false},
		{"Nil should not be a subtree", args{tree5, subtree5}, false},
		{"should be a subtree", args{tree6, subtree6}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("IsSubtree(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftMostNode(t *testing.T) {
	t.Parallel()
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should find left most node of tree", args{newFullBinaryTree(t)}, newFullBinaryTree(t).Left().Left().Left()},
		{"should be a root node of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, newRightSkewedBinaryTree(t)},
		{"should be a last node of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, newLeftSkewedBinaryTree(t).Left().Left().Left()},
		{"should be a nil of nil", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := LeftMostNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lLeftMostNode(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightMostNode(t *testing.T) {
	t.Parallel()
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should find right most node of tree", args{newFullBinaryTree(t)}, newFullBinaryTree(t).Right().Right().Right()},
		{"should be a last node of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, newRightSkewedBinaryTree(t).Right().Right().Right()},
		{"should be a root node of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, newLeftSkewedBinaryTree(t)},
		{"should be a nil of nil", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := RightMostNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightMostNode(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	t.Parallel()
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should get max depth = 4", args{newFullBinaryTree(t)}, 4},
		{"should get zero for nil root", args{nil}, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := MaxDepth(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxDepth(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRoot(t *testing.T) {
	t.Parallel()
	type args struct {
		node *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"Get root for left-most node in the tree", args{LeftMostNode(newFullBinaryTree(t))}, newFullBinaryTree(t)},
		{"Get root for right-most node in the tree", args{RightMostNode(newFullBinaryTree(t))}, newFullBinaryTree(t)},
		{"Get root for nested node in the tree", args{newFullBinaryTree(t).Left().Right()}, newFullBinaryTree(t)},
		{"Get the same node if the node and is the root of the tree", args{newFullBinaryTree(t)}, newFullBinaryTree(t)},
		{"Get the nil for nil node", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetRoot(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerialize(t *testing.T) {
	t.Parallel()
	type args struct {
		root *BinaryNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Serialize of full binary tree", args{newFullBinaryTree(t)}, serializedFullBinaryTree},
		{"Serialize of complete binary tree", args{newCompleteBinaryTree(t)}, serializedCompleteBinaryTree},
		{"Serialize of prefect binary tree", args{newPerfectBinaryTree(t)}, serializedPerfectBinaryTree},
		{"Serialize of degenerate binary tree", args{newDegenerateBinaryTree(t)}, serializedDegenerateBinaryTree},
		{"Serialize of left-skewed binary tree", args{newLeftSkewedBinaryTree(t)}, serializedLeftSkewedBinaryTree},
		{"Serialize of right-skewed binary tree", args{newRightSkewedBinaryTree(t)}, serializedRightSkewedBinaryTree},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Serialize(tt.args.root); got != tt.want {
				t.Errorf("Serialize(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeserialize(t *testing.T) {
	t.Parallel()
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[string]
	}{
		{"Deserialize serialized tree", args{Serialize(newFullBinaryTreeS(t))}, newFullBinaryTreeS(t)},
		{"Deserialize of full binary tree", args{serializedFullBinaryTree}, newFullBinaryTreeS(t)},
		{"Deserialize of complete binary tree", args{serializedCompleteBinaryTree}, newCompleteBinaryTreeS(t)},
		{"Deserialize of prefect binary tree", args{serializedPerfectBinaryTree}, newPerfectBinaryTreeS(t)},
		{"Deserialize of degenerate binary tree", args{serializedDegenerateBinaryTree}, newDegenerateBinaryTreeS(t)},
		{"Deserialize of left-skewed binary tree", args{serializedLeftSkewedBinaryTree}, newLeftSkewedBinaryTreeS(t)},
		{"Deserialize of right-skewed binary tree", args{serializedRightSkewedBinaryTree}, newRightSkewedBinaryTreeS(t)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Deserialize(tt.args.str); !IsIdentical(got, tt.want) {
				t.Errorf("Deserialize(str) = %v, want %v", *got, *tt.want)
			}
		})
	}
}

// DONT CHANGE
// HELPERS FOR TESTS

const (
	serializedFullBinaryTree        = "^1,^2,^4,^8,#,#,^9,#,#,^5,#,#,^3,^6,#,#,^7,^10,#,#,^11,#,#,"
	serializedCompleteBinaryTree    = "^1,^2,^4,^8,#,#,^9,#,#,^5,^10,#,#,^11,#,#,^3,^6,#,#,^7,#,#,"
	serializedPerfectBinaryTree     = "^1,^2,^4,#,#,^5,#,#,^3,^6,#,#,^7,#,#,"
	serializedDegenerateBinaryTree  = "^1,^2,#,^3,#,^4,#,#,#,"
	serializedLeftSkewedBinaryTree  = "^1,^2,^3,^4,#,#,#,#,#,"
	serializedRightSkewedBinaryTree = "^1,#,^2,#,^3,#,^4,#,#,"
)

func traverseCallbackFactory[T any](t *testing.T, sb *strings.Builder) func(n *BinaryNode[T]) {
	t.Helper()
	return func(n *BinaryNode[T]) {
		sb.WriteString(fmt.Sprintf("%v", n.Value()))
		sb.WriteString(SerializationDelimiter)
	}
}

func traverseEmptyFactory(t *testing.T, sb *strings.Builder) func() {
	t.Helper()
	return func() {
		sb.WriteString(SerializationEnd)
		sb.WriteString(SerializationDelimiter)
	}
}

func newFullBinaryTree(t *testing.T) *BinaryNode[int] {
	t.Helper()
	bt := NewBinaryTree(1)
	bt.left = NewBinaryTree(2)
	bt.left.parent = bt
	bt.left.left = NewBinaryTree(4)
	bt.left.left.parent = bt.left
	bt.left.left.left = NewBinaryTree(8)
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = NewBinaryTree(9)
	bt.left.left.right.parent = bt.left.left
	bt.left.right = NewBinaryTree(5)
	bt.left.right.parent = bt.left
	bt.right = NewBinaryTree(3)
	bt.right.parent = bt
	bt.right.left = NewBinaryTree(6)
	bt.right.left.parent = bt.right
	bt.right.right = NewBinaryTree(7)
	bt.right.right.parent = bt.right
	bt.right.right.left = NewBinaryTree(10)
	bt.right.right.left.parent = bt.right.right
	bt.right.right.right = NewBinaryTree(11)
	bt.right.right.right.parent = bt.right.right
	return bt
}

func newFullBinaryTreeS(t *testing.T) *BinaryNode[string] {
	t.Helper()
	bt := NewBinaryTree("1")
	bt.left = NewBinaryTree("2")
	bt.left.parent = bt
	bt.left.left = NewBinaryTree("4")
	bt.left.left.parent = bt.left
	bt.left.left.left = NewBinaryTree("8")
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = NewBinaryTree("9")
	bt.left.left.right.parent = bt.left.left
	bt.left.right = NewBinaryTree("5")
	bt.left.right.parent = bt.left
	bt.right = NewBinaryTree("3")
	bt.right.parent = bt
	bt.right.left = NewBinaryTree("6")
	bt.right.left.parent = bt.right
	bt.right.right = NewBinaryTree("7")
	bt.right.right.parent = bt.right
	bt.right.right.left = NewBinaryTree("10")
	bt.right.right.left.parent = bt.right.right
	bt.right.right.right = NewBinaryTree("11")
	bt.right.right.right.parent = bt.right.right
	return bt
}

func newCompleteBinaryTree(t *testing.T) *BinaryNode[int] {
	t.Helper()
	bt := NewBinaryTree(1)
	bt.left = NewBinaryTree(2)
	bt.left.parent = bt
	bt.left.left = NewBinaryTree(4)
	bt.left.left.parent = bt.left
	bt.left.left.left = NewBinaryTree(8)
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = NewBinaryTree(9)
	bt.left.left.right.parent = bt.left.left
	bt.left.right = NewBinaryTree(5)
	bt.left.right.parent = bt.left
	bt.left.right.left = NewBinaryTree(10)
	bt.left.right.left.parent = bt.left.right
	bt.left.right.right = NewBinaryTree(11)
	bt.left.right.right.parent = bt.left.right
	bt.right = NewBinaryTree(3)
	bt.right.parent = bt
	bt.right.left = NewBinaryTree(6)
	bt.right.left.parent = bt.right
	bt.right.right = NewBinaryTree(7)
	bt.right.right.parent = bt.right
	return bt
}

func newCompleteBinaryTreeS(t *testing.T) *BinaryNode[string] {
	t.Helper()
	bt := NewBinaryTree("1")
	bt.left = NewBinaryTree("2")
	bt.left.parent = bt
	bt.left.left = NewBinaryTree("4")
	bt.left.left.parent = bt.left
	bt.left.left.left = NewBinaryTree("8")
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = NewBinaryTree("9")
	bt.left.left.right.parent = bt.left.left
	bt.left.right = NewBinaryTree("5")
	bt.left.right.parent = bt.left
	bt.left.right.left = NewBinaryTree("10")
	bt.left.right.left.parent = bt.left.right
	bt.left.right.right = NewBinaryTree("11")
	bt.left.right.right.parent = bt.left.right
	bt.right = NewBinaryTree("3")
	bt.right.parent = bt
	bt.right.left = NewBinaryTree("6")
	bt.right.left.parent = bt.right
	bt.right.right = NewBinaryTree("7")
	bt.right.right.parent = bt.right
	return bt
}

func newPerfectBinaryTree(t *testing.T) *BinaryNode[int] {
	t.Helper()
	bt := NewBinaryTree(1)
	bt.left = NewBinaryTree(2)
	bt.left.parent = bt
	bt.left.left = NewBinaryTree(4)
	bt.left.left.parent = bt.left
	bt.left.right = NewBinaryTree(5)
	bt.left.right.parent = bt.left
	bt.right = NewBinaryTree(3)
	bt.right.parent = bt
	bt.right.left = NewBinaryTree(6)
	bt.right.left.parent = bt.right
	bt.right.right = NewBinaryTree(7)
	bt.right.right.parent = bt.right
	return bt
}

func newPerfectBinaryTreeS(t *testing.T) *BinaryNode[string] {
	t.Helper()
	bt := NewBinaryTree("1")
	bt.left = NewBinaryTree("2")
	bt.left.parent = bt
	bt.left.left = NewBinaryTree("4")
	bt.left.left.parent = bt.left
	bt.left.right = NewBinaryTree("5")
	bt.left.right.parent = bt.left
	bt.right = NewBinaryTree("3")
	bt.right.parent = bt
	bt.right.left = NewBinaryTree("6")
	bt.right.left.parent = bt.right
	bt.right.right = NewBinaryTree("7")
	bt.right.right.parent = bt.right
	return bt
}

func newDegenerateBinaryTree(t *testing.T) *BinaryNode[int] {
	t.Helper()
	bt := NewBinaryTree(1)
	bt.left = NewBinaryTree(2)
	bt.left.parent = bt
	bt.left.right = NewBinaryTree(3)
	bt.left.right.parent = bt.left
	bt.left.right.right = NewBinaryTree(4)
	bt.left.right.right.parent = bt.left.right
	return bt
}

func newDegenerateBinaryTreeS(t *testing.T) *BinaryNode[string] {
	t.Helper()
	bt := NewBinaryTree("1")
	bt.left = NewBinaryTree("2")
	bt.left.parent = bt
	bt.left.right = NewBinaryTree("3")
	bt.left.right.parent = bt.left
	bt.left.right.right = NewBinaryTree("4")
	bt.left.right.right.parent = bt.left.right
	return bt
}

func newLeftSkewedBinaryTree(t *testing.T) *BinaryNode[int] {
	t.Helper()
	bt := NewBinaryTree(1)
	bt.left = NewBinaryTree(2)
	bt.left.parent = bt
	bt.left.left = NewBinaryTree(3)
	bt.left.left.parent = bt.left
	bt.left.left.left = NewBinaryTree(4)
	bt.left.left.left.parent = bt.left.left
	return bt
}

func newLeftSkewedBinaryTreeS(t *testing.T) *BinaryNode[string] {
	t.Helper()
	bt := NewBinaryTree("1")
	bt.left = NewBinaryTree("2")
	bt.left.parent = bt
	bt.left.left = NewBinaryTree("3")
	bt.left.left.parent = bt.left
	bt.left.left.left = NewBinaryTree("4")
	bt.left.left.left.parent = bt.left.left
	return bt
}

func newRightSkewedBinaryTree(t *testing.T) *BinaryNode[int] {
	t.Helper()
	bt := NewBinaryTree(1)
	bt.right = NewBinaryTree(2)
	bt.right.parent = bt
	bt.right.right = NewBinaryTree(3)
	bt.right.right.parent = bt.right
	bt.right.right.right = NewBinaryTree(4)
	bt.right.right.right.parent = bt.right.right
	return bt
}

func newRightSkewedBinaryTreeS(t *testing.T) *BinaryNode[string] {
	t.Helper()
	bt := NewBinaryTree("1")
	bt.right = NewBinaryTree("2")
	bt.right.parent = bt
	bt.right.right = NewBinaryTree("3")
	bt.right.right.parent = bt.right
	bt.right.right.right = NewBinaryTree("4")
	bt.right.right.right.parent = bt.right.right
	return bt
}
