package tree

import (
	"testing"

	"github.com/dkhrunov/dsa-go/structures/queue"
	"github.com/stretchr/testify/require"
)

func TestNewAVLNode(t *testing.T) {
	t.Parallel()
	node := NewAVLNode[int](10)
	require.NotNil(t, node)
	require.Equal(t, 10, node.value)
	require.Zero(t, node.bf)
	require.Zero(t, node.height)
	require.Zero(t, node.left)
	require.Zero(t, node.right)
	require.Zero(t, node.parent)
}

func TestNewAVLTree(t *testing.T) {
	t.Parallel()
	avl := NewAVLTree[int]()

	require.NotNil(t, avl)
	require.Zero(t, avl.root)
	require.Zero(t, avl.size)
}

func TestAVLTreeRoot(t *testing.T) {
	t.Parallel()

	t.Run("should return root of avl", func(t *testing.T) {
		t.Parallel()
		avl := createAVLTree(t)
		require.NotZero(t, avl.Root())
	})

	t.Run("should return nil root of avl", func(t *testing.T) {
		t.Parallel()
		avl := &AVLTree[int]{}
		require.Zero(t, avl.Root())
	})
}

func TestAVLTreeSize(t *testing.T) {
	t.Parallel()
	type TreeFactory func(t *testing.T) *AVLTree[int]
	tests := []struct {
		name   string
		treeFc TreeFactory
		want   int
	}{
		{
			name:   "should return zero size of avl",
			treeFc: func(t *testing.T) *AVLTree[int] { t.Helper(); return &AVLTree[int]{} },
			want:   0,
		},
		{
			name:   "should return size == 7",
			treeFc: createAVLTree,
			want:   7,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			avl := tt.treeFc(t)
			require.Equal(t, tt.want, avl.Size())
		})
	}
}

func TestAVLTreeHeight(t *testing.T) {
	t.Parallel()
	type TreeFactory func(t *testing.T) *AVLTree[int]
	tests := []struct {
		name   string
		treeFc TreeFactory
		want   int
	}{
		{
			name:   "should return 0 height for empty tree",
			treeFc: func(t *testing.T) *AVLTree[int] { t.Helper(); return &AVLTree[int]{} },
			want:   0,
		},
		{
			name:   "should return height == e",
			treeFc: createAVLTree,
			want:   3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			avl := tt.treeFc(t)
			require.Equal(t, tt.want, avl.Height())
		})
	}
}

func TestAVLTreeInsert(t *testing.T) {
	t.Parallel()
	type iteration struct {
		insert int
		root   int
	}
	tests := []struct {
		name       string
		iterations []iteration
		serialized string
	}{
		{
			name: "left-left case: should right rotate",
			iterations: []iteration{
				{insert: 5, root: 5},
				{insert: 4, root: 5},
				{insert: 3, root: 4},
				{insert: 2, root: 4},
				{insert: 1, root: 4},
				{insert: 0, root: 2},
			},
			serialized: "^2,^1,^0,#,#,#,^4,^3,#,#,^5,#,#,",
		},
		{
			name: "left-right case: should left-right rotate",
			iterations: []iteration{
				{insert: 5, root: 5},
				{insert: 3, root: 5},
				{insert: 4, root: 4},
			},
			serialized: "^4,^3,#,#,^5,#,#,",
		},
		{
			name: "right-right case: should left rotate",
			iterations: []iteration{
				{insert: 5, root: 5},
				{insert: 6, root: 5},
				{insert: 7, root: 6},
				{insert: 8, root: 6},
				{insert: 9, root: 6},
				{insert: 10, root: 8},
			},
			serialized: "^8,^6,^5,#,#,^7,#,#,^9,#,^10,#,#,",
		},
		{
			name: "right-left case: should right-left rotate",
			iterations: []iteration{
				{insert: 3, root: 3},
				{insert: 5, root: 3},
				{insert: 4, root: 4},
			},
			serialized: "^4,^3,#,#,^5,#,#,",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			avl := NewAVLTree[int]()
			for i, v := range tt.iterations {
				require.True(t, avl.Insert(v.insert))
				require.Equal(t, i+1, avl.Size())
				require.Equal(t, v.root, avl.Root().Value())
			}
			require.Equal(t, tt.serialized, avl.Serialize())
			require.True(t, isBalance(t, avl.Root()))
		})
	}

	t.Run("should not add duplicated node", func(t *testing.T) {
		t.Parallel()
		avl := NewAVLTree[int]()
		require.True(t, avl.Insert(5))
		require.True(t, avl.Insert(4))
		require.True(t, avl.Insert(3))
		require.False(t, avl.Insert(3))
		require.False(t, avl.Insert(3))
		require.Equal(t, avl.Size(), 3)
		require.Equal(t, "^4,^3,#,#,^5,#,#,", avl.Serialize())
	})
}

func TestAVLTreeDelete(t *testing.T) {
	t.Parallel()
	type TreeFactory func(t *testing.T) *AVLTree[int]
	tests := []struct {
		name       string
		treeFc     TreeFactory
		delete     int
		size       int
		serialized string
		success    bool
	}{
		{
			name:       "should delete: Node to remove is a leaf node",
			treeFc:     createAVLTree,
			delete:     1,
			size:       6,
			serialized: "^8,^5,^3,#,#,^6,#,#,^10,#,^14,#,#,",
			success:    true,
		},
		{
			name:       "should delete: Node to remove has a right subtree but no left subtree",
			treeFc:     createAVLTree,
			delete:     10,
			size:       6,
			serialized: "^6,^3,^1,#,#,^5,#,#,^8,#,^14,#,#,",
			success:    true,
		},
		{
			name:       "should delete: Node to remove has a left subtree but no right subtree",
			treeFc:     createAVLTree,
			delete:     6,
			size:       6,
			serialized: "^8,^3,^1,#,#,^5,#,#,^10,#,^14,#,#,",
			success:    true,
		},
		{
			name:       "should delete: Node to remove has a both a left subtree and a right subtree, where left.height > right.height",
			treeFc:     createAVLTree,
			delete:     8,
			size:       6,
			serialized: "^6,^3,^1,#,#,^5,#,#,^10,#,^14,#,#,",
			success:    true,
		},
		{
			name:       "should delete: Node to remove has a both a left subtree and a right subtree, where left.height < right.height",
			treeFc:     createAVLTree,
			delete:     3,
			size:       6,
			serialized: "^8,^5,^1,#,#,^6,#,#,^10,#,^14,#,#,",
			success:    true,
		},
		{
			name:       "should not delete: Node that doesn't exist",
			treeFc:     createAVLTree,
			delete:     999,
			size:       7,
			serialized: "^8,^3,^1,#,#,^6,^5,#,#,#,^10,#,^14,#,#,",
			success:    false,
		},
		{
			name:       "should not delete: Where Root is nil",
			treeFc:     func(t *testing.T) *AVLTree[int] { t.Helper(); return NewAVLTree[int]() },
			delete:     1,
			size:       0,
			serialized: "",
			success:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			avl := tt.treeFc(t)
			require.Equal(t, tt.success, avl.Delete(tt.delete))
			require.Equal(t, tt.size, avl.Size())
			require.Equal(t, tt.serialized, avl.Serialize())
			require.True(t, isBalance(t, avl.Root()))
			require.False(t, avl.Contains(tt.delete))
		})
	}
}

func TestAVLTreeSearch(t *testing.T) {
	t.Parallel()
	avl := createAVLTree(t)
	tests := []struct {
		name   string
		search int
		want   *AVLNode[int]
	}{
		{
			name:   "should found value",
			search: 6,
			want:   avl.Root().Left().Right(),
		},
		{
			name:   "should not found value",
			search: 999,
			want:   nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Same(t, tt.want, avl.Search(tt.search))
		})
	}
}

func TestAVLTreeContains(t *testing.T) {
	t.Parallel()
	avl := createAVLTree(t)
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
			require.Equal(t, tt.want, avl.Contains(tt.value))
		})
	}
}

func createAVLTree(t *testing.T) *AVLTree[int] {
	t.Helper()
	avl := &AVLTree[int]{}
	avl.Insert(8)
	avl.Insert(3)
	avl.Insert(10)
	avl.Insert(1)
	avl.Insert(6)
	avl.Insert(14)
	avl.Insert(5)
	return avl
}

func isBalance(t *testing.T, root *AVLNode[int]) bool {
	if root == nil {
		return true
	}

	balanced := true

	checkNodeBf := func(n *AVLNode[int]) {
		if n.bf <= -2 || n.bf >= 2 {
			balanced = false
		}
	}

	traverseLevelorder(t, root, checkNodeBf)

	return balanced
}

func traverseLevelorder(t *testing.T, root *AVLNode[int], cb func(node *AVLNode[int])) {
	if root == nil {
		return
	}

	queue := queue.New()
	queue.EnQueue(root)

	for queue.Len() > 0 {
		n := queue.DeQueue().(*AVLNode[int])
		cb(n)

		if n.Left() != nil {
			queue.EnQueue(n.Left())
		}

		if n.Right() != nil {
			queue.EnQueue(n.Right())
		}
	}
}
