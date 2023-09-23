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

func TestBSTSize(t *testing.T) {
	t.Parallel()

	t.Run("should return size of BST", func(t *testing.T) {
		t.Parallel()
		bst := createBST(t)
		require.NotZero(t, bst.Size())
	})

	t.Run("should return zero size of BST", func(t *testing.T) {
		t.Parallel()
		bst := &BSTree[int]{}
		require.Zero(t, bst.Size())
	})
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
		bst := &BSTree[int]{}
		require.Zero(t, bst.Root())
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
	tests := []struct {
		name   string
		insert int
		want   func(root *BinaryNode[int]) int
		size   int
	}{
		{
			name:   "should insert node with value 4",
			insert: 4,
			want:   func(root *BinaryNode[int]) int { return root.Left().Right().Left().value },
			size:   8,
		},
		{
			name:   "should insert node with value 2",
			insert: 2,
			want:   func(root *BinaryNode[int]) int { return root.Left().Left().Right().value },
			size:   8,
		},
		{
			name:   "should insert node with value 9",
			insert: 9,
			want:   func(root *BinaryNode[int]) int { return root.Right().Left().value },
			size:   8,
		},
		{
			name:   "should insert node with value 5",
			insert: 5,
			want:   func(root *BinaryNode[int]) int { return root.Left().Right().Left().value },
			size:   8,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bst := createBST(t)
			require.True(t, bst.Insert(tt.insert))
			require.Equal(t, tt.want(bst.Root()), tt.insert)
			require.Equal(t, bst.Size(), tt.size)
		})
	}

	t.Run("should not add duplicated node", func(t *testing.T) {
		t.Parallel()
		bst := createBST(t)
		require.True(t, bst.Insert(9))
		require.False(t, bst.Insert(9))
		require.False(t, bst.Insert(9))
		require.Equal(t, bst.Size(), 8)
	})
}

func TestBSTDelete(t *testing.T) {
	t.Parallel()
	type TreeFactory func(t *testing.T) *BSTree[int]
	tests := []struct {
		name       string
		treeFc     TreeFactory
		delete     int
		serialized string
		size       int
		success    bool
	}{
		{
			name:       "should delete leaf node",
			treeFc:     createBST,
			delete:     1,
			serialized: "^8,^3,#,^6,#,^7,#,#,^10,#,^14,#,#,",
			size:       6,
			success:    true,
		},
		{
			name:       "should delete node with single child node",
			treeFc:     createBST,
			delete:     6,
			serialized: "^8,^3,^1,#,#,^7,#,#,^10,#,^14,#,#,",
			size:       6,
			success:    true,
		},
		{
			name:       "should delete node with two children",
			treeFc:     createBST,
			delete:     3,
			serialized: "^8,^6,^1,#,#,^7,#,#,^10,#,^14,#,#,",
			size:       6,
			success:    true,
		},
		{
			name:       "should do nothing with not existing value",
			treeFc:     createBST,
			delete:     999,
			serialized: "^8,^3,^1,#,#,^6,#,^7,#,#,^10,#,^14,#,#,",
			size:       7,
			success:    false,
		},
		{
			name:       "should not delete: Where Root is nil",
			treeFc:     func(t *testing.T) *BSTree[int] { t.Helper(); return NewBST[int]() },
			delete:     1,
			serialized: "",
			size:       0,
			success:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			bst := tt.treeFc(t)
			require.Equal(t, tt.success, bst.Delete(tt.delete))
			require.Equal(t, tt.size, bst.Size())
			require.Equal(t, tt.serialized, Serialize(bst.Root()))
			require.False(t, bst.Contains(tt.delete))
		})
	}
}

func TestBSTSearch(t *testing.T) {
	t.Parallel()
	bst := createBST(t)
	tests := []struct {
		name   string
		search int
		want   *BinaryNode[int]
	}{
		{
			name:   "should found value",
			search: 6,
			want:   bst.Root().Left().Right(),
		},
		{
			name:   "should not found value",
			search: -999,
			want:   nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bst.Search(tt.search)
			require.Same(t, tt.want, got)
		})
	}
}

func TestBSTContains(t *testing.T) {
	t.Parallel()
	bst := createBST(t)
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{
			name:  "Should contain",
			value: 6,
			want:  true,
		},
		{
			name:  "Should not contain",
			value: 999,
			want:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, bst.Contains(tt.value))
		})
	}
}

func createBST(t *testing.T) *BSTree[int] {
	t.Helper()
	bst := &BSTree[int]{}
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
