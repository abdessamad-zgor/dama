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

func Assert(condition bool, elseerror string) {
	if !condition {
		panic(errors.New(elseerror))
	}
}
