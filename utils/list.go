package utils

type List[V comparable] struct {
	items *[]V
}

func NewList[V comparable]() List[V] {
	items := make([]V, 0)
	list := List[V] {
		&items,
	}
	return list
}

func (list List[T]) Add(element T) {
	*list.items = append(*list.items, element)
}

func (list List[T]) Remove(index int) {
	ielem := 0
	for i, _ := range *list.items {
		if index == i {
			ielem = i
			break
		}
	}
	*(list.items) = append((*list.items)[0:ielem], (*list.items)[ielem + 1:]...)
}

func (list List[T]) Insert(item T, index int) {
	if index >= 0 && index <= list.Length() {
		clist := make([]T, list.Length())
		copy(clist, *(list.items))
		*(list.items) = append(clist[:index], item)
		*(list.items) = append(*(list.items), clist[index:]...)
	}
}


func (list List[T]) Empty() {
	*(list.items) = make([]T, 0)
}

func (list List[T]) Items() []T {
	return *(list.items)
}

func (list List[T]) Length() int {
	return len(*list.items)
}
