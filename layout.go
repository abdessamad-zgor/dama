package dama

import (
	"errors"
)

type Position interface{}

type GridLayout struct {
	Columns  uint
	Rows     uint
	Elements map[GridPosition]DamaElement
}

type GridPosition struct {
	Column     uint
	Row        uint
	RowSpan    uint
	ColumnSpan uint
}

type BaseLayout struct {
	Elements map[BasePosition]DamaElement
}

type BasePosition = Direction

const (
	Center BasePosition = "center"
)

type DamaLayout interface {
	GetElements() []DamaElement
	AddElement(element DamaElement, position Position) error
}

func (layout *GridLayout) AddWidget(element DamaElement, position Position) error {
	gridPosition, ok := position.(GridPosition)
	if !ok {
		return errors.New("position is not of type GridPosition")
	}
	if gridPosition.Row >= layout.Rows || gridPosition.Column >= layout.Columns || gridPosition.Column+gridPosition.ColumnSpan > layout.Columns || gridPosition.Row+gridPosition.RowSpan > layout.Rows {
		return errors.New("widget position is out of bounds")
	}
	layout.Elements[gridPosition] = element
	return nil
}

func (layout *BaseLayout) AddWidget(element DamaElement, position Position) error {
	basePosition, ok := position.(BasePosition)
	if !ok {
		return errors.New("position is not of type BasePosition")
	}
	layout.Elements[basePosition] = element
	return nil
}

func (layout *GridLayout) GetElements() []DamaElement {
	elements := []DamaElement{}
	for _, w := range layout.Elements {
		elements = append(elements, w)
	}
	return elements
}

func (layout *BaseLayout) Getelements() []DamaElement {
	elements := []DamaElement{}
	for _, w := range layout.Elements {
		elements = append(elements, w)
	}
	return elements
}
