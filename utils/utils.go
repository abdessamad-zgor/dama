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

func Assert(condition bool, elseerror string) {
	if !condition {
		panic(errors.New(elseerror))
	}
}
