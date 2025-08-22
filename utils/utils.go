package utils

import (
	"fmt"
	"errors"
)

type List[V any] struct {
	items []V
}

func NewList[V any]() List[V] {
	items := make([]V)
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
	list.items = append(list.items[0:ielem], ...list.items[ielem + 1:])
}

type VList[V any] struct {
	items	map[string]List[V]
	current string
}

func NewVList[V any]() Vlist[V] {
	default := make(map[string]List[V])
	vlist := VList[V] {
		&default,
		"",
	}
	return vlist
}

func (vlist VList[V]) Add(element V) {
	vlist.items[vlist.current].Add(element)
}

func (vlist VList[V]) Remove(element V) {
	vlist.items[vlist.current].Remove(element)
}

func (vlist VList[V]) AddView(key string) {
	vlist.items[key] = NewList[V]()
}

func (vlist VList[V]) RemoveView(key string) {
	delete(vlist.items, key)
}


type EList[V any] struct {
	items	List[V]
}

func NewEList[V any]() Elist[V] {
	elist := EList[V] {
		NewList[V](),
	}
	return elist
}

func (elist EList[V]) Add(element V, ef func (V, V) bool) {
	for _, elem := range elist.items {
	}
}

func (elist EList[V]) Remove(element V) {}

func Assert(condition bool, elseerror string) {
	if !condition {
		panic(errors.New(elseerror))
	}
}
