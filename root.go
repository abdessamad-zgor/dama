package utils

import (
	"errors"
)

type ViewList[K comparable, V] struct {
	items 	*map[K][]V
	current K
}

func NewViewList[K comparable, V]() ViewList[K, V] {
	itemsMap = make(map[K][]V)
	viewList := ViewList[K, V]{
		&itemsMap,
	}

	return viewList
}


func (vlist ViewList[K, V]) Add(key K, items V) {
	vlist.items[vlist.current] = append(vlist.items[vlist.current], V)
}

func (vlist ViewList[K, V] Remove(itemIndex int) {
	vlist.items[vlist.current] = append(vlist.items[vlist.current][:itemIndex], vlist.items[vlist.current][itemIndex+1:])
}

func (vlist ViewList[K, V]) Find(predicate func(value V) bool) (V, error) {
	var value V
	for _, v := range vlist.items[vlist.current] {
		if predicate(v) {
			return v, nil
		}
	}
	return value, errors.New("No element found that satifies the predicate")
}

func (vlist ViewList[K, V] FindAll(predicate func(V) bool) []V {
	found := []V{}
	for _, v := range vlist.items[vlist.current] {
		if predicate(v) {
			found = append(found, v)
		}
	}
	return found
}

func (vlist ViewList[K, V]) FindIndex(predicate func(V) bool) int {
	for i, v := range vlist.items[vlist.current] {
		if predicate(v) {
			return i
		}
	}
	return -1
}
