package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBST(t *testing.T) {
	t.Parallel()
	bst := NewBST[int]()

	require.NotNil(t, bst)
	require.Zero(t, bst.root)
	require.Zero(t, bst.size)

}

func TestBSTRoot(t *testing.T) {
	t.Parallel()

	t.Run("should return root of BST", func(t *testing.T) {
		t.Parallel()
		bst := createBST(t)
		require.NotZero(t, bst.Root())
	})

	t.Run("should return nil root of BST", func(t *testing.T) {
		t.Parallel()
		bst := &BST[int]{}
		require.Zero(t, bst.Root())
	})
}

func TestBSTSize(t *testing.T) {
	t.Parallel()

	t.Run("should return size of BST", func(t *testing.T) {
		t.Parallel()
		bst := createBST(t)
		require.NotZero(t, bst.Root())
	})

	t.Run("should return zero size of BST", func(t *testing.T) {
		t.Parallel()
		bst := &BST[int]{}
		require.Zero(t, bst.Size())
	})
}

func TestBSTMin(t *testing.T) {
	t.Parallel()
	bst := createBST(t)
	tests := []struct {
		name string
		want int
	}{
		{"should return min value of BST", 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bst.Min()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBSTMax(t *testing.T) {
	t.Parallel()
	bst := createBST(t)
	tests := []struct {
		name string
		want int
	}{
		{"should return max value of BST", 14},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bst.Max()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBSTInsert(t *testing.T) {
	t.Parallel()
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want func(root *BinaryNode[int]) int
	}{
		{"should insert node with value 4", args{4}, func(root *BinaryNode[int]) int { return root.Left().Right().Left().value }},
		{"should insert node with value 2", args{2}, func(root *BinaryNode[int]) int { return root.Left().Left().Right().value }},
		{"should insert node with value 9", args{9}, func(root *BinaryNode[int]) int { return root.Right().Left().value }},
		{"should insert node with value 5", args{5}, func(root *BinaryNode[int]) int { return root.Left().Right().Left().value }},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bst := createBST(t)
			sizeBefore := bst.Size()
			bst.Insert(tt.args.value)
			require.Equal(t, tt.want(bst.Root()), tt.args.value)
			require.Equal(t, bst.Size(), sizeBefore+1)
		})
	}

	t.Run("should not add duplicated node", func(t *testing.T) {
		t.Parallel()
		bst := createBST(t)
		sizeBefore := bst.Size()
		bst.Insert(8)
		require.Empty(t, searchRec(bst.root.left, 8))
		require.Empty(t, searchRec(bst.root.right, 8))
		require.NotEmpty(t, searchRec(bst.root, 8))
		require.Equal(t, bst.Size(), sizeBefore+1)
	})
}

func TestBSTDelete(t *testing.T) {
	t.Parallel()
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"should delete leaf node", args{1}, "^8,^3,#,^6,#,^7,#,#,^10,#,^14,#,#,"},
		{"should delete node with single child node", args{6}, "^8,^3,^1,#,#,^7,#,#,^10,#,^14,#,#,"},
		{"should delete node with two children", args{3}, "^8,^6,^1,#,#,^7,#,#,^10,#,^14,#,#,"},
		{"should do nothing with not existing value", args{-999}, "^8,^3,^1,#,#,^6,#,^7,#,#,^10,#,^14,#,#,"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bst := createBST(t)
			sizeBefore := bst.Size()
			bst.Delete(tt.args.value)
			require.Equal(t, tt.want, Serialize(bst.Root()))
			require.Equal(t, bst.Size(), sizeBefore-1)
		})
	}
}

func TestBSTSearch(t *testing.T) {
	t.Parallel()
	bst := createBST(t)
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode[int]
	}{
		{"should found value", args{6}, bst.Root().Left().Right()},
		{"should found smallest value than root", args{1}, bst.Root().Left().Left()},
		{"should found bigger value than root", args{14}, bst.Root().Right().Right()},
		{"should not found value", args{-999}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bst.Search(tt.args.value)
			require.Same(t, tt.want, got)
		})
	}
}

func createBST(t *testing.T) *BST[int] {
	t.Helper()
	bst := &BST[int]{}
	bst.root = &BinaryNode[int]{value: 8}
	bst.root.left = &BinaryNode[int]{value: 3}
	bst.root.right = &BinaryNode[int]{value: 10}
	bst.root.left.left = &BinaryNode[int]{value: 1}
	bst.root.left.right = &BinaryNode[int]{value: 6}
	bst.root.right.right = &BinaryNode[int]{value: 14}
	bst.root.left.right.right = &BinaryNode[int]{value: 7}
	bst.size = 7
	return bst
}
