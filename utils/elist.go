package utils

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
