package utils

import (
	"errors"
)

type List[V comparable] struct {
	items []V
}

func NewList[V comparable]() List[V] {
	items := make([]V, 1)
	list := List[V] {
		items,
	}
	return list
}

func (list List[T]) Add(element T) {
	list.items = append(list.items, element)
}

func (list List[T]) Remove(element T) {
	ielem := 0
	for i, elem := range list.items {
		if elem == element {
			ielem = i
			break
		}
	}
	list.items = append(list.items[0:ielem], list.items[ielem + 1:]...)
}

func (list List[T]) Insert(item T, index int) {
	if index >= 0 && index <= list.Length() {
		clist := make([]T, list.Length())
		copy(clist, list.items)
		list.items = append(clist[:index], item)
		list.items = append(list.items, clist[index:]...)
	}
}

func (list List[T]) Empty() {
	list.items = make([]T, 0)
}

func (list List[T]) Items() []T {
	return list.items
}

func (list List[T]) Length() int {
	return len(list.items)
}

type VList[V comparable] struct {
	items	*map[string]List[V]
	current string
}

func NewVList[V comparable]() VList[V] {
	_default := make(map[string]List[V])
	vlist := VList[V] {
		&_default,
		"",
	}
	return vlist
}

func (vlist VList[V]) Add(element V) {
	(*(vlist.items))[vlist.current].Add(element)
}

func (vlist VList[V]) Remove(element V) {
	(*(vlist.items))[vlist.current].Remove(element)
}

func (vlist VList[V]) AddView(key string) {
	(*(vlist.items))[key] = NewList[V]()
}

func (vlist VList[V]) RemoveView(key string) {
	delete(*(vlist.items), key)
}

func (vlist VList[V]) Length() int {
	return ((*vlist.items)[vlist.current]).Length()
}

func (vlist VList[T]) Items() []T {
	return ((*vlist.items)[vlist.current]).Items()
}

type ExcludeFn[T comparable] func (list List[T], item T) int 

type EList[V comparable] struct {
	items		List[V]
	excludeFn 	ExcludeFn[V]
}

func NewEList[V comparable](excludeFn ExcludeFn[V]) EList[V] {
	elist := EList[V] {
		NewList[V](),
		excludeFn,
	}
	return elist
}

func (elist EList[V]) Add(element V) {
	insertIndex := elist.excludeFn(elist.items, element) 
	if insertIndex == elist.Length() {
		elist.items.Add(element)
	} else if insertIndex >= 0 {
		elist.items.Insert(element, insertIndex)
	}
}

func (elist EList[V]) Remove(element V) {
	elist.items.Remove(element)
}

func (elist EList[V]) Empty() {
	elist.items.Empty()
}

func (elist EList[T]) Items() []T {
	return elist.items.Items()
}

func (elist EList[V]) Length() int {
	return elist.items.Length()
}

type Node[T comparable] struct {
	Id			int
	Parent		*Node
	Value		T
	Children	[]Node[T]
}

type Tree[T comparable] struct {
	Root	*Node
}

func NewTree[T comparable](root T) Tree[T] {
	rootNode := Node[T] {
		iota,
		nil,
		root,
		make([]Node[T], 0),
	}
	return Tree[T] {
		&rootNode,
	}
}

func (tree Tree[T]) Subtree(value T) Tree[T] {
	nodeNode := tree.FindNode(node)
	subtreeRoot := Node[T] {
		iota,
		nil,
		nodeNode.Value,
		nodeNode.Children,
	}
	subtree := Tree[T] {
		&subtreeRoot,
	}
	return subtree
}

func (tree Tree[T]) Flatten() []Node[T] {
	current := *tree.Root
	flatTree := []Node[T]{current}
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
		if node.Value == value {
			return &flatTree[i]
		}
	} 
	return nil
}

func (tree Tree[T]) AddNode(parent T, child T) {
	parentNode := tree.FindNode(parent)
	childNode :=  Node[T] {
		iota,
		parentNode,
		child,
		make([]Node[T], 0),
	}
	parentNode.Children = append(parentNode.Children, childNode)
}

func (tree Tree[T]) Remove(value T) {
	nodeNode := tree.FindNode(value)
	parentNode := nodeNode.Parent
	i := slices.IndexFunc(parentNode.Children, func (child Node[T]) bool {
		return child.Value = value
	})
	if i >= 0 {
		parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
		parentNode.Children = append(parentNode.Children, nodeNode.Children...)
	}
}

func (tree Tree[T]) RemoveSubtree(value T) {
	valueNode := tree.FindNode(value)
	parentNode := nodeNode.Parent
	i := slices.IndexFunc(parentNode.Children, func (child Node[T]) bool {
		return child.Value = value
	})
	if i >= 0 {
		parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
	}
}

func Assert(condition bool, elseerror string) {
	if !condition {
		panic(errors.New(elseerror))
	}
}
