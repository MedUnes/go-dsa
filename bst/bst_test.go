package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewlyCreatedBSTShouldBeEmpty(t *testing.T) {
	tree := new(BST)
	assert.True(t, tree.IsEmpty(), "Newly created BST should be empty")
}

func TestBSTInsertion(t *testing.T) {
	tree := new(BST)
	tree.Insert(1)
	assert.False(t, tree.IsEmpty())
	assert.Equal(t, tree.Root.Val, 1)
	tree.Insert(10)
	tree.Insert(100)
	tree.Insert(0)
	tree.Insert(5)
	tree.Insert(15)
	assert.Equal(t, tree.Root.Right.Val, 10)
	assert.Equal(t, tree.Root.Right.Right.Val, 100)
	assert.Equal(t, tree.Root.Left.Val, 0)
	assert.Equal(t, tree.Root.Right.Right.Left.Val, 15)
}

func TestBSTHas(t *testing.T) {
	tree := new(BST)
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(-10)
	tree.Insert(13)
	tree.Insert(1120)

	assert.True(t, tree.Has(15))
	assert.True(t, tree.Has(10))
	assert.True(t, tree.Has(-10))
	assert.True(t, tree.Has(13))
	assert.True(t, tree.Has(1120))

	assert.False(t, tree.Has(0))

}

func TestBSTRemove(t *testing.T) {
	tree := new(BST)

	t.Run("empty returns empty", func(t *testing.T) {
		tree.Remove(0)
		assert.True(t, tree.IsEmpty())

	})

	t.Run("single noded tree returns empty", func(t *testing.T) {

		tree.Insert(10)
		tree.Remove(10)
		assert.True(t, tree.IsEmpty())
	})

	t.Run("intermediate from right-only tree", func(t *testing.T) {

		tree.Insert(10)
		tree.Insert(15)
		tree.Insert(25)

		tree.Remove(15)
		assert.Equal(t, tree.Root.Right.Val, 25)
		assert.Nil(t, tree.Root.Right.Right)
	})

	t.Run("leaf from right-only tree", func(t *testing.T) {
		tree.Insert(10)
		tree.Insert(15)
		tree.Insert(25)

		tree.Remove(25)
		assert.Equal(t, tree.Root.Right.Val, 15)
		assert.Nil(t, tree.Root.Right.Right)
	})

	t.Run("root from single-node tree", func(t *testing.T) {
		tree = new(BST)
		tree.Insert(42)
		tree.Remove(42)
		assert.True(t, tree.IsEmpty())
	})

	t.Run("root (has only left subtree)", func(t *testing.T) {
		tree.Insert(20)
		tree.Insert(10)
		tree.Insert(5)
		tree.Remove(20)
		assert.Equal(t, 10, tree.Root.Val)
		assert.Equal(t, 5, tree.Root.Left.Val)
	})

	t.Run("root (has only right subtree)", func(t *testing.T) {
		tree = new(BST)
		tree.Insert(20)
		tree.Insert(30)
		tree.Insert(40)
		tree.Remove(20)
		assert.Equal(t, 30, tree.Root.Val)
		assert.Equal(t, 40, tree.Root.Right.Val)

	})

	t.Run("root (has both children) – inorder successor scenario", func(t *testing.T) {
		tree = new(BST)
		tree.Insert(20)
		tree.Insert(10)
		tree.Insert(30)
		tree.Insert(25)
		tree.Insert(35)
		tree.Remove(20)
		assert.Equal(t, 25, tree.Root.Val)
		assert.Equal(t, 30, tree.Root.Right.Val)
		assert.Nil(t, tree.Root.Right.Left)
	})

	t.Run("leaf on left side", func(t *testing.T) {
		tree = new(BST)
		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(3)
		tree.Remove(3)
		assert.Nil(t, tree.Root.Left.Left)
	})
	t.Run("Remove non-existent key (tree unchanged)", func(t *testing.T) {
		tree = new(BST)
		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Remove(42)
		assert.Equal(t, 10, tree.Root.Val)
		assert.Equal(t, 5, tree.Root.Left.Val)
		assert.Equal(t, 15, tree.Root.Right.Val)
	})

}
