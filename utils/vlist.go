package utils

import (
	"github.com/abdessamad-zgor/dama/logger"
)

type VList[V comparable] struct {
	items	*map[string]List[V]
	current string
}

func NewVList[V comparable]() VList[V] {
	_default := make(map[string]List[V])
	_default[""] = NewList[V]()
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
	logger.Log((*(vlist.items))[vlist.current])
	return ((*vlist.items)[vlist.current]).Length()
}

func (vlist VList[T]) Items() []T {
	return ((*vlist.items)[vlist.current]).Items()
}
