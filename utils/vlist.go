package utils

import (
	_ "github.com/abdessamad-zgor/dama/logger"
)

type VList[K, V comparable] struct {
	items	*map[K]List[V]
	current K
}

func NewVList[K, V comparable](views ...K) VList[K, V] {
	Assert(len(views) > 0, "VList can not estantiate with no views")
	_default := make(map[K]List[V])
	for _, view := range views {
		_default[view] = NewList[V]()
	}
	vlist := VList[K, V] {
		&_default,
		views[0],
	}
	return vlist
}

func (vlist VList[K, V]) Add(element V) {
	(*(vlist.items))[vlist.current].Add(element)
}

func (vlist VList[K, V]) AddToView(view K, element V) {
	(*(vlist.items))[view].Add(element)
}

func (vlist VList[K, V]) Remove(element V) {
	(*(vlist.items))[vlist.current].Remove(element)
}

func (vlist VList[K, V]) AddView(key K) {
	(*(vlist.items))[key] = NewList[V]()
}

func (vlist VList[K, V]) SwitchView(key K) {
	vlist.current = key
}

func (vlist VList[K, V]) RemoveView(key K) {
	delete(*(vlist.items), key)
}

func (vlist VList[K, V]) Length() int {
	return ((*vlist.items)[vlist.current]).Length()
}

func (vlist VList[K, T]) Items() []T {
	return ((*vlist.items)[vlist.current]).Items()
}
