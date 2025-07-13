package utils

import (
	"fmt"
	"errors"
)

type List[V any] struct {
	items []V
}

func NewList[V any] () List[V] {
	items := make([]V)
	list := List[V] {
		items,
	}
	return list
}

func (list List[T]) Add(element T) {

}

func (list List[T]) Remove(element T) {

}

type VListKey interface {
	comparable
}

type VList[K comparable, V any] struct {
	items	map[K]List[V]
	current K
}

func NewVList[K comparable, V any]() Vlist[K, V] {

}

func (vlist VList[K, V]) Add(element V) {}

func (vlist VList[K, V]) Remove(element V) {}

func (vlist VList[K, V]) AddView(key K) {}

func (vlist VList[K, V]) Switch(to K) {}

func (vlist VList[K, V]) Keys() []K {}

func (vlist VList[K, V]) Values() []V {}

func (vlist VList[K, V]) RemoveView(key K) {}

func Assert(condition bool, elseerror string) {
	if !condition {
		panic(errors.New(elseerror))
	}
}
