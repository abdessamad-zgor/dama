package dama

import (
	"errors"
)

type DamaLayout interface {
	GetElements() []DamaElement
	AddElement(element DamaElement, position Position) error
}

type Position interface{}

type GridLayout struct {
	Container *Container
	Columns   uint
	Rows      uint
	Elements  map[GridPosition]DamaElement
}

type GridPosition struct {
	Column     uint
	Row        uint
	RowSpan    uint
	ColumnSpan uint
}

type BaseLayout struct {
	Container *Container
	Elements  map[BasePosition]DamaElement
}

type BasePosition = Direction

const (
	Center BasePosition = "center"
)

func (layout *GridLayout) AddElement(element DamaElement, position Position) error {
	gridPosition, ok := position.(GridPosition)
	if !ok {
		return errors.New("position is not of type GridPosition")
	}

	if gridPosition.Row >= layout.Rows || gridPosition.Column >= layout.Columns || gridPosition.Column+gridPosition.ColumnSpan > layout.Columns || gridPosition.Row+gridPosition.RowSpan > layout.Rows {
		return errors.New("widget position is out of bounds")
	}

	x := uint(layout.Container.X) + uint(layout.Container.Width/layout.Columns)*gridPosition.Column
	y := uint(layout.Container.Y) + uint(layout.Container.Height/layout.Rows)*gridPosition.Row
	width := uint(layout.Container.Width/layout.Columns) * gridPosition.ColumnSpan
	height := uint(layout.Container.Height/layout.Rows) * gridPosition.RowSpan

	element.SetBox(x, y, width, height)

	layout.Elements[gridPosition] = element
	return nil
}

func (layout *BaseLayout) isPositionFree(position BasePosition) bool {
	_, ok := layout.Elements[position]
	return !ok
}

func (layout *BaseLayout) AddElement(element DamaElement, position Position) error {
	basePosition, ok := position.(BasePosition)
	if !ok {
		return errors.New("position is not of type BasePosition")
	}

	x, y, width, height := uint(0), uint(0), uint(0), uint(0)
	switch basePosition {
	case Center:
		x = (layout.Container.X + layout.Container.Width/5)
		y = (layout.Container.Y + layout.Container.Height/5)
		width = (layout.Container.Width/5 * 3)
		height = (layout.Container.Height/5 * 3)
	case Left:
		x = (layout.Container.X)
		y = (layout.Container.Y) + (layout.Container.Height/5)
		width = (layout.Container.Width/5) * 3
		height = (layout.Container.Height/5) * 3
	case Right:
		x = (layout.Container.X) + (layout.Container.Width/5) * 4
		y = (layout.Container.Y) + (layout.Container.Height/5)
		width = (layout.Container.Width/5) * 3
		height = (layout.Container.Height/5) * 3
	case Top:
		x = (layout.Container.X) 
		y = (layout.Container.Y)
		width = layout.Container.Width
		height = (layout.Container.Height/5) 
	case Bottom:
		x = (layout.Container.X)
		y = (layout.Container.Y) + (layout.Container.Height/5) * 4
		width = layout.Container.Width
		height = (layout.Container.Height/5)
	}

    element.SetBox(uint(x), uint(y), uint(width), uint(height))
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

func (layout *BaseLayout) GetElements() []DamaElement {
	elements := []DamaElement{}
	for _, w := range layout.Elements {
		elements = append(elements, w)
	}
	return elements
}
