package binarytree

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	if got := New(0); got == nil {
		t.Errorf("New() can't create new binary tree node struct")
	}
}

func TestValue(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree()
	node := &Tree[int]{}
	type args struct {
		n *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Should get the node value", args{bt}, 1},
		{"Should get the nested node value", args{bt.Left().Left()}, 4},
		{"Should get default zero value for empty node", args{node}, 0},
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

func TestSetValue(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree()
	type args struct {
		n      *Tree[int]
		newVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Should set new value for node", args{bt, 99}, 99},
		{"Should set new value for nested node", args{bt.Left().Left(), 88}, 88},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			prevVal := tt.args.n.Value()
			tt.args.n.SetValue(tt.args.newVal)
			got := tt.args.n.Value()
			if prevVal == got {
				t.Error("SetValue(...args) doesnt change the value", got, tt.want)
			}
			if got != tt.want {
				t.Errorf("SetValue(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree()
	node := &Tree[int]{}
	type args struct {
		n *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should get the left child of the node", args{bt}, bt.left},
		{"Should get the nil if left child does not exist ", args{node}, nil},
		{"Should get node pointer nil of nil", args{nil}, nil},
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
	bt := newPerfectBinaryTree()
	node := &Tree[int]{}
	type args struct {
		n *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should get the right child of the node", args{bt}, bt.right},
		{"Should get the nil if right child does not exist ", args{node}, nil},
		{"Should get node pointer nil of nil", args{nil}, nil},
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
	t.Run("Should be inserted after leaf node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree()
		after := LeftMostNode(bt)
		new := New(99)
		InsertAfter(after, new)
		if !reflect.DeepEqual(after.right, new) {
			t.Errorf("InsertAfter(...args) should be inserted as right child of %v node", after)
		}
		if !reflect.DeepEqual(new.parent, after) {
			t.Error("InsertAfter(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("Should be inserted after root node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree()
		after := LeftMostNode(bt.right)
		new := New(99)
		InsertAfter(bt, new)
		if !reflect.DeepEqual(LeftMostNode(bt.right), new) {
			t.Errorf("InsertAfter(...args) should be inserted as left child of %v node", after)
		}
		if !reflect.DeepEqual(new.parent, after) {
			t.Error("InsertAfter(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("Should be inserted after internal mode", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree()
		after := LeftMostNode(bt).parent
		parent := LeftMostNode(after.right)
		new := New(99)
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
	t.Run("Should be inserted before leaf node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree()
		before := LeftMostNode(bt)
		new := New(99)
		InsertBefore(before, new)
		if !reflect.DeepEqual(before.left, new) {
			t.Errorf("InsertBefore(...args) should be inserted as left child of %v node", before)
		}
		if !reflect.DeepEqual(new.parent, before) {
			t.Error("InsertBefore(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("Should be inserted before root node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree()
		before := RightMostNode(bt.left)
		new := New(99)
		InsertBefore(bt, new)
		if !reflect.DeepEqual(RightMostNode(bt.left), new) {
			t.Errorf("InsertBefore(...args) should be inserted as right child of %v node", before)
		}
		if !reflect.DeepEqual(new.parent, before) {
			t.Error("InsertBefore(...args) should set correct parent pointer for inserted node")
		}
	})

	t.Run("Should be inserted before internal node", func(t *testing.T) {
		t.Parallel()
		bt := newFullBinaryTree()
		before := RightMostNode(bt).parent
		parent := RightMostNode(before.left)
		new := New(99)
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
		node *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Delete root node", args{newFullBinaryTree()}, "^5,^2,^4,^8,#,#,^9,#,#,#,^3,^6,#,#,^7,^10,#,#,^11,#,#,"},
		{"Delete leaf node", args{newFullBinaryTree().Left().Right()}, "^1,^2,^4,^8,#,#,^9,#,#,#,^3,^6,#,#,^7,^10,#,#,^11,#,#,"},
		{"Delete node without left child", args{
			func() *Tree[int] {
				r := newFullBinaryTree()
				// Delete from left subtree from right subtree
				r.right.left.parent = nil
				r.right.left = nil
				return r.right
			}(),
		}, "^5,^2,^4,^8,#,#,^9,#,#,#,^1,#,^7,^10,#,#,^11,#,#,"},
		{"Delete node without right child", args{
			func() *Tree[int] {
				r := newFullBinaryTree()
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
		r := newFullBinaryTree()
		// Delete from left subtree from right subtree
		r.left.left.left.parent = nil
		r.left.left.left = nil
		Delete(r.left.left)
		t.Errorf("Should have panicked")
	})
}

func TestTraversePreorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Preorder traverse of full binary tree", args{newFullBinaryTree()}, "1,2,4,8,#,#,9,#,#,5,#,#,3,6,#,#,7,10,#,#,11,#,#,"},
		{"Preorder traverse of complete binary tree", args{newCompleteBinaryTree()}, "1,2,4,8,#,#,9,#,#,5,10,#,#,11,#,#,3,6,#,#,7,#,#,"},
		{"Preorder traverse of prefect binary tree", args{newPerfectBinaryTree()}, "1,2,4,#,#,5,#,#,3,6,#,#,7,#,#,"},
		{"Preorder traverse of degenerate binary tree", args{newDegenerateBinaryTree()}, "1,2,#,3,#,4,#,#,#,"},
		{"Preorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, "1,2,3,4,#,#,#,#,#,"},
		{"Preorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree()}, "1,#,2,#,3,#,4,#,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraversePreorder(tt.args.root, traverseCallbackFactory[int](sb), traverseEmptyFactory(sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraversePreorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraverseInorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Inorder traverse of full binary tree", args{newFullBinaryTree()}, "#,8,#,4,#,9,#,2,#,5,#,1,#,6,#,3,#,10,#,7,#,11,#,"},
		{"Inorder traverse of complete binary tree", args{newCompleteBinaryTree()}, "#,8,#,4,#,9,#,2,#,10,#,5,#,11,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of prefect binary tree", args{newPerfectBinaryTree()}, "#,4,#,2,#,5,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of degenerate binary tree", args{newDegenerateBinaryTree()}, "#,2,#,3,#,4,#,1,#,"},
		{"Inorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, "#,4,#,3,#,2,#,1,#,"},
		{"Inorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree()}, "#,1,#,2,#,3,#,4,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraverseInorder(tt.args.root, traverseCallbackFactory[int](sb), traverseEmptyFactory(sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseInorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraverseInorderI(t *testing.T) {
	t.Parallel()
	type args struct {
		root *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Inorder traverse of full binary tree", args{newFullBinaryTree()}, "#,8,#,4,#,9,#,2,#,5,#,1,#,6,#,3,#,10,#,7,#,11,#,"},
		{"Inorder traverse of complete binary tree", args{newCompleteBinaryTree()}, "#,8,#,4,#,9,#,2,#,10,#,5,#,11,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of prefect binary tree", args{newPerfectBinaryTree()}, "#,4,#,2,#,5,#,1,#,6,#,3,#,7,#,"},
		{"Inorder traverse of degenerate binary tree", args{newDegenerateBinaryTree()}, "#,2,#,3,#,4,#,1,#,"},
		{"Inorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, "#,4,#,3,#,2,#,1,#,"},
		{"Inorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree()}, "#,1,#,2,#,3,#,4,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraverseInorderI(tt.args.root, traverseCallbackFactory[int](sb), traverseEmptyFactory(sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseInorderI(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraversePostorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Postorder traverse of full binary tree", args{newFullBinaryTree()}, "#,#,8,#,#,9,4,#,#,5,2,#,#,6,#,#,10,#,#,11,7,3,1,"},
		{"Postorder traverse of complete binary tree", args{newCompleteBinaryTree()}, "#,#,8,#,#,9,4,#,#,10,#,#,11,5,2,#,#,6,#,#,7,3,1,"},
		{"Postorder traverse of prefect binary tree", args{newPerfectBinaryTree()}, "#,#,4,#,#,5,2,#,#,6,#,#,7,3,1,"},
		{"Postorder traverse of degenerate binary tree", args{newDegenerateBinaryTree()}, "#,#,#,#,4,3,2,#,1,"},
		{"Postorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, "#,#,4,#,3,#,2,#,1,"},
		{"Postorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree()}, "#,#,#,#,#,4,3,2,1,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraversePostorder(tt.args.root, traverseCallbackFactory[int](sb), traverseEmptyFactory(sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraversePostorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestTraverseLevelorder(t *testing.T) {
	t.Parallel()
	type args struct {
		root *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TraverseLevelorder traverse of full binary tree", args{newFullBinaryTree()}, "1,2,3,4,5,#,#,6,#,#,7,8,#,#,9,#,#,10,#,#,11,#,#,"},
		{"TraverseLevelorder traverse of complete binary tree", args{newCompleteBinaryTree()}, "1,2,3,4,5,6,#,#,7,#,#,8,#,#,9,#,#,10,#,#,11,#,#,"},
		{"TraverseLevelorder traverse of prefect binary tree", args{newPerfectBinaryTree()}, "1,2,3,4,#,#,5,#,#,6,#,#,7,#,#,"},
		{"TraverseLevelorder traverse of degenerate binary tree", args{newDegenerateBinaryTree()}, "1,#,2,#,3,#,4,#,#,"},
		{"TraverseLevelorder traverse of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, "1,#,2,#,3,#,4,#,#,"},
		{"TraverseLevelorder traverse of right-skewed binary tree", args{newRightSkewedBinaryTree()}, "1,#,2,#,3,#,4,#,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sb := &strings.Builder{}
			TraverseLevelorder(tt.args.root, traverseCallbackFactory[int](sb), traverseEmptyFactory(sb))
			if got := sb.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseLevelorder(...args) = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

func TestInorderSuccessor(t *testing.T) {
	t.Parallel()
	bt := newPerfectBinaryTree()
	type args struct {
		node *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should be the left most node of the target's right subtree", args{bt}, bt.Right().Left()},
		{"Should be the parent node of the target", args{bt.Left().Left()}, bt.Left()},
		{"Should be the grandparent node of the target", args{bt.Left().Right()}, bt},
		{"Should be the nil, if target last node in tree", args{bt.Right().Right()}, nil},
		{"Should be the right node of the target's right subtree", args{bt.Right()}, bt.Right().Right()},
		{"Should be a nil of nil root", args{nil}, nil},
		{"Should be the nil of right-most node in tree", args{RightMostNode(bt)}, nil},
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
	bt := newPerfectBinaryTree()
	type args struct {
		root   *Tree[int]
		target *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should be the left most node of the target's right subtree", args{bt, bt}, bt.Right().Left()},
		{"Should be the parent node of the target", args{bt, bt.Left().Left()}, bt.Left()},
		{"Should be the grandparent node of the target", args{bt, bt.Left().Right()}, bt},
		{"Should be the nil, if target last node in tree", args{bt, bt.Right().Right()}, nil},
		{"Should be the right node of the target's right subtree", args{bt, bt.Right()}, bt.Right().Right()},
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
	bt := newPerfectBinaryTree()
	type args struct {
		root   *Tree[int]
		target *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should be the left-most node of the target's right subtree", args{bt, bt}, bt.Right().Left()},
		{"Should be the parent node of the target", args{bt, bt.Left().Left()}, bt.Left()},
		{"Should be the grandparent of the target", args{bt, bt.Left().Right()}, bt},
		{"Should be null if the target node is the last node in the inorder traversal", args{bt, bt.Right().Right()}, nil},
		{"Should be the right node of the target's right subtree", args{bt, bt.Right()}, bt.Right().Right()},
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
	bt := newFullBinaryTree()
	type args struct {
		node *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should be the right-most node of the target's left subtree", args{bt}, bt.Left().Right()},
		{"Should be the parent node of the target #1", args{bt.Left().Left().Right()}, bt.Left().Left()},
		{"Should be root node for the left-most node in the right subtree", args{bt.Right().Left()}, bt},
		{"Should be null if the target node is the first node in the inorder traversal", args{LeftMostNode(bt)}, nil},
		{"Should be a nil of nil root", args{nil}, nil},
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
	tree1X := newFullBinaryTree()
	tree1Y := newFullBinaryTree()

	tree2X := newCompleteBinaryTree()
	tree2Y := newCompleteBinaryTree()
	tree2Y.Right().Right().SetValue(99)

	type args struct {
		x *Tree[int]
		y *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Should be identical", args{tree1X, tree1Y}, true},
		{"Should not be identical", args{tree2X, tree2Y}, false},
		{"Two nil roots should be identical", args{nil, nil}, true},
		{"If one is nil and the other is not, then it should not be identical", args{newFullBinaryTree(), nil}, false},
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
	tree1 := newFullBinaryTree()

	subtree1 := tree1.left

	tree2 := newFullBinaryTree()

	subtree2 := New(2)
	subtree2.left = New(4)
	subtree2.left.left = New(8)
	subtree2.right = New(5)

	tree3 := New(12)

	subtree3 := New(2)

	tree4 := New(1)
	tree4.left = New(2)
	tree4.left.left = New(4)
	tree4.right = New(3)

	subtree4 := New(1)
	subtree4.left = New(2)
	subtree4.right = New(3)

	tree5 := newFullBinaryTree()

	var subtree5 *Tree[int]

	tree6 := New(1)

	subtree6 := New(1)

	type args struct {
		root    *Tree[int]
		subRoot *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Should be a subtree", args{tree1, subtree1}, true},
		{"Should not be a subtree", args{tree2, subtree2}, false},
		{"Should not be a subtree", args{tree3, subtree3}, false},
		{"Should not be a subtree", args{tree4, subtree4}, false},
		{"Nil should not be a subtree", args{tree5, subtree5}, false},
		{"Should be a subtree", args{tree6, subtree6}, true},
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
		node *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should find left most node of tree", args{newFullBinaryTree()}, newFullBinaryTree().Left().Left().Left()},
		{"Should be a root node of right-skewed binary tree", args{newRightSkewedBinaryTree()}, newRightSkewedBinaryTree()},
		{"Should be a last node of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, newLeftSkewedBinaryTree().Left().Left().Left()},
		{"Should be a nil of nil", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := LeftMostNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftMostNode(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightMostNode(t *testing.T) {
	t.Parallel()
	type args struct {
		node *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Should find right most node of tree", args{newFullBinaryTree()}, newFullBinaryTree().Right().Right().Right()},
		{"Should be a last node of right-skewed binary tree", args{newRightSkewedBinaryTree()}, newRightSkewedBinaryTree().Right().Right().Right()},
		{"Should be a root node of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, newLeftSkewedBinaryTree()},
		{"Should be a nil of nil", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := RightMostNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightMostNode(...args) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRoot(t *testing.T) {
	t.Parallel()
	type args struct {
		node *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want *Tree[int]
	}{
		{"Get root for left-most node in the tree", args{LeftMostNode(newFullBinaryTree())}, newFullBinaryTree()},
		{"Get root for right-most node in the tree", args{RightMostNode(newFullBinaryTree())}, newFullBinaryTree()},
		{"Get root for nested node in the tree", args{newFullBinaryTree().Left().Right()}, newFullBinaryTree()},
		{"Get the same node if the node and is the root of the tree", args{newFullBinaryTree()}, newFullBinaryTree()},
		{"Get the nil for nil node", args{nil}, nil},
	}
	for _, tt := range tests {
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
		root *Tree[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Serialize of full binary tree", args{newFullBinaryTree()}, serializedFullBinaryTree},
		{"Serialize of complete binary tree", args{newCompleteBinaryTree()}, serializedCompleteBinaryTree},
		{"Serialize of prefect binary tree", args{newPerfectBinaryTree()}, serializedPerfectBinaryTree},
		{"Serialize of degenerate binary tree", args{newDegenerateBinaryTree()}, serializedDegenerateBinaryTree},
		{"Serialize of left-skewed binary tree", args{newLeftSkewedBinaryTree()}, serializedLeftSkewedBinaryTree},
		{"Serialize of right-skewed binary tree", args{newRightSkewedBinaryTree()}, serializedRightSkewedBinaryTree},
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
		want *Tree[string]
	}{
		{"Deserialize serialized tree", args{Serialize(newFullBinaryTreeS())}, newFullBinaryTreeS()},
		{"Deserialize of full binary tree", args{serializedFullBinaryTree}, newFullBinaryTreeS()},
		{"Deserialize of complete binary tree", args{serializedCompleteBinaryTree}, newCompleteBinaryTreeS()},
		{"Deserialize of prefect binary tree", args{serializedPerfectBinaryTree}, newPerfectBinaryTreeS()},
		{"Deserialize of degenerate binary tree", args{serializedDegenerateBinaryTree}, newDegenerateBinaryTreeS()},
		{"Deserialize of left-skewed binary tree", args{serializedLeftSkewedBinaryTree}, newLeftSkewedBinaryTreeS()},
		{"Deserialize of right-skewed binary tree", args{serializedRightSkewedBinaryTree}, newRightSkewedBinaryTreeS()},
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

func traverseCallbackFactory[T any](sb *strings.Builder) func(n *Tree[T]) {
	return func(n *Tree[T]) {
		sb.WriteString(fmt.Sprintf("%v", n.Value()))
		sb.WriteString(SerializationDelimiter)
	}
}

func traverseEmptyFactory(sb *strings.Builder) func() {
	return func() {
		sb.WriteString(SerializationEnd)
		sb.WriteString(SerializationDelimiter)
	}
}

func newFullBinaryTree() *Tree[int] {
	bt := New(1)
	bt.left = New(2)
	bt.left.parent = bt
	bt.left.left = New(4)
	bt.left.left.parent = bt.left
	bt.left.left.left = New(8)
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = New(9)
	bt.left.left.right.parent = bt.left.left
	bt.left.right = New(5)
	bt.left.right.parent = bt.left
	bt.right = New(3)
	bt.right.parent = bt
	bt.right.left = New(6)
	bt.right.left.parent = bt.right
	bt.right.right = New(7)
	bt.right.right.parent = bt.right
	bt.right.right.left = New(10)
	bt.right.right.left.parent = bt.right.right
	bt.right.right.right = New(11)
	bt.right.right.right.parent = bt.right.right
	return bt
}

func newFullBinaryTreeS() *Tree[string] {
	bt := New("1")
	bt.left = New("2")
	bt.left.parent = bt
	bt.left.left = New("4")
	bt.left.left.parent = bt.left
	bt.left.left.left = New("8")
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = New("9")
	bt.left.left.right.parent = bt.left.left
	bt.left.right = New("5")
	bt.left.right.parent = bt.left
	bt.right = New("3")
	bt.right.parent = bt
	bt.right.left = New("6")
	bt.right.left.parent = bt.right
	bt.right.right = New("7")
	bt.right.right.parent = bt.right
	bt.right.right.left = New("10")
	bt.right.right.left.parent = bt.right.right
	bt.right.right.right = New("11")
	bt.right.right.right.parent = bt.right.right
	return bt
}

func newCompleteBinaryTree() *Tree[int] {
	bt := New(1)
	bt.left = New(2)
	bt.left.parent = bt
	bt.left.left = New(4)
	bt.left.left.parent = bt.left
	bt.left.left.left = New(8)
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = New(9)
	bt.left.left.right.parent = bt.left.left
	bt.left.right = New(5)
	bt.left.right.parent = bt.left
	bt.left.right.left = New(10)
	bt.left.right.left.parent = bt.left.right
	bt.left.right.right = New(11)
	bt.left.right.right.parent = bt.left.right
	bt.right = New(3)
	bt.right.parent = bt
	bt.right.left = New(6)
	bt.right.left.parent = bt.right
	bt.right.right = New(7)
	bt.right.right.parent = bt.right
	return bt
}

func newCompleteBinaryTreeS() *Tree[string] {
	bt := New("1")
	bt.left = New("2")
	bt.left.parent = bt
	bt.left.left = New("4")
	bt.left.left.parent = bt.left
	bt.left.left.left = New("8")
	bt.left.left.left.parent = bt.left.left
	bt.left.left.right = New("9")
	bt.left.left.right.parent = bt.left.left
	bt.left.right = New("5")
	bt.left.right.parent = bt.left
	bt.left.right.left = New("10")
	bt.left.right.left.parent = bt.left.right
	bt.left.right.right = New("11")
	bt.left.right.right.parent = bt.left.right
	bt.right = New("3")
	bt.right.parent = bt
	bt.right.left = New("6")
	bt.right.left.parent = bt.right
	bt.right.right = New("7")
	bt.right.right.parent = bt.right
	return bt
}

func newPerfectBinaryTree() *Tree[int] {
	bt := New(1)
	bt.left = New(2)
	bt.left.parent = bt
	bt.left.left = New(4)
	bt.left.left.parent = bt.left
	bt.left.right = New(5)
	bt.left.right.parent = bt.left
	bt.right = New(3)
	bt.right.parent = bt
	bt.right.left = New(6)
	bt.right.left.parent = bt.right
	bt.right.right = New(7)
	bt.right.right.parent = bt.right
	return bt
}

func newPerfectBinaryTreeS() *Tree[string] {
	bt := New("1")
	bt.left = New("2")
	bt.left.parent = bt
	bt.left.left = New("4")
	bt.left.left.parent = bt.left
	bt.left.right = New("5")
	bt.left.right.parent = bt.left
	bt.right = New("3")
	bt.right.parent = bt
	bt.right.left = New("6")
	bt.right.left.parent = bt.right
	bt.right.right = New("7")
	bt.right.right.parent = bt.right
	return bt
}

func newDegenerateBinaryTree() *Tree[int] {
	bt := New(1)
	bt.left = New(2)
	bt.left.parent = bt
	bt.left.right = New(3)
	bt.left.right.parent = bt.left
	bt.left.right.right = New(4)
	bt.left.right.right.parent = bt.left.right
	return bt
}

func newDegenerateBinaryTreeS() *Tree[string] {
	bt := New("1")
	bt.left = New("2")
	bt.left.parent = bt
	bt.left.right = New("3")
	bt.left.right.parent = bt.left
	bt.left.right.right = New("4")
	bt.left.right.right.parent = bt.left.right
	return bt
}

func newLeftSkewedBinaryTree() *Tree[int] {
	bt := New(1)
	bt.left = New(2)
	bt.left.parent = bt
	bt.left.left = New(3)
	bt.left.left.parent = bt.left
	bt.left.left.left = New(4)
	bt.left.left.left.parent = bt.left.left
	return bt
}

func newLeftSkewedBinaryTreeS() *Tree[string] {
	bt := New("1")
	bt.left = New("2")
	bt.left.parent = bt
	bt.left.left = New("3")
	bt.left.left.parent = bt.left
	bt.left.left.left = New("4")
	bt.left.left.left.parent = bt.left.left
	return bt
}

func newRightSkewedBinaryTree() *Tree[int] {
	bt := New(1)
	bt.right = New(2)
	bt.right.parent = bt
	bt.right.right = New(3)
	bt.right.right.parent = bt.right
	bt.right.right.right = New(4)
	bt.right.right.right.parent = bt.right.right
	return bt
}

func newRightSkewedBinaryTreeS() *Tree[string] {
	bt := New("1")
	bt.right = New("2")
	bt.right.parent = bt
	bt.right.right = New("3")
	bt.right.right.parent = bt.right
	bt.right.right.right = New("4")
	bt.right.right.right.parent = bt.right.right
	return bt
}
