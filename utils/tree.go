package utils

import (
	_ "errors"
	"slices"
)

var _index int = 0

func Id() int {
	_index += 1
	return _index
}

type Node[T comparable] struct {
	Id			int
	Parent		*Node[T]
	Value		*T
	Children	[]*Node[T]
}

func (node Node[T]) SetValue(value T) {
	*node.Value = value
}

func (node Node[T]) GetValue() T {
	return *node.Value
}

type Tree[T comparable] struct {
	Root	*Node[T]
}

func NewTree[T comparable](root T) Tree[T] {
	rootNode := Node[T] {
		Id(),
		nil,
		&root,
		make([]*Node[T], 0),
	}
	return Tree[T] {
		&rootNode,
	}
}

func (tree Tree[T]) Subtree(value T) Tree[T] {
	nodeNode := tree.FindNode(value)
	subtreeRoot := Node[T] {
		Id(),
		nil,
		nodeNode.Value,
		nodeNode.Children,
	}
	subtree := Tree[T] {
		&subtreeRoot,
	}
	return subtree
}

func (tree Tree[T]) Flatten() []*Node[T] {
	current := tree.Root
	flatTree := []*Node[T]{current}
	paths := current.Children
	for len(paths) > 0 {
		current = paths[len(paths) - 1]
		paths = paths[:len(paths) - 1]
		if len(current.Children) > 0 {
			paths = append(paths, current.Children...)
		}
		flatTree = append(flatTree, current)
	}
	return flatTree
}

func (tree Tree[T]) FindNode(value T) *Node[T] {
	flatTree := tree.Flatten()
	for i, node := range flatTree {
		if *node.Value == value {
			return flatTree[i]
		}
	} 
	return nil
}

func (tree Tree[T]) AddNode(parent T, child T) {
	parentNode := tree.FindNode(parent)
	childNode :=  Node[T] {
		Id(),
		parentNode,
		&child,
		make([]*Node[T], 0),
	}
	parentNode.Children = append(parentNode.Children, &childNode)
}

func (tree Tree[T]) Remove(value T) {
	nodeNode := tree.FindNode(value)
	parentNode := nodeNode.Parent
	i := slices.IndexFunc(parentNode.Children, func (child *Node[T]) bool {
		return *child.Value == value
	})
	if i >= 0 {
		parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
		parentNode.Children = append(parentNode.Children, nodeNode.Children...)
	}
}

func (tree Tree[T]) RemoveSubtree(value T) {
	valueNode := tree.FindNode(value)
	parentNode := valueNode.Parent
	i := slices.IndexFunc(parentNode.Children, func (child *Node[T]) bool {
		return *child.Value == value
	})
	if i >= 0 {
		parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
	}
}

