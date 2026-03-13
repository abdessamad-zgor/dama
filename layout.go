package dama

import (
	"errors"
)

type Layout interface {
	GetElements() []Element
	AddElement(element Element, position Position) error
}

type Position interface{}

type GridLayout struct {
	Container *container_s
	Columns   int
	Rows      int
	Elements  map[GridPosition]Element
}

func NewGridLayout(columns int, rows int) *GridLayout {
	grid := new(GridLayout)
	grid.Columns = int(columns)
	grid.Rows = int(rows)
	grid.Elements = make(map[GridPosition]Element)
	return grid
}

type GridPosition struct {
	Column     int
	Row        int
	RowSpan    int
	ColumnSpan int
}

type BaseLayout struct {
	Container *container_s
	Elements  map[BasePosition]Element
}

type BasePosition = Direction

func (layout *GridLayout) AddElement(element Element, position Position) error {
	gridPosition, ok := position.(GridPosition)
	if !ok {
		return errors.New("position is not of type GridPosition")
	}

	if gridPosition.Row >= layout.Rows || gridPosition.Column >= layout.Columns || gridPosition.Column+gridPosition.ColumnSpan > layout.Columns || gridPosition.Row+gridPosition.RowSpan > layout.Rows {
		return errors.New("widget position is out of bounds")
	}

	cont := layout.Container.GetBox()
	x := (cont.X) + int(cont.Width/layout.Columns)*gridPosition.Column
	y := (cont.Y) + int(cont.Height/layout.Rows)*gridPosition.Row
	width := int(cont.Width/layout.Columns) * gridPosition.ColumnSpan
	height := int(cont.Height/layout.Rows) * gridPosition.RowSpan

	element.SetBox(x, y, width, height)

	layout.Elements[gridPosition] = element
	return nil
}


func (layout *BaseLayout) getBoxForPosition(position BasePosition) (int, int, int, int) {
	cont := layout.Container.GetBox()
	x, y, width, height := (0), (0), (0), (0)
	switch position {
	case Center:
		x = (cont.X + cont.Width/5)
		y = (cont.Y + cont.Height/5)
		width = (cont.Width / 5 * 3)
		height = (cont.Height / 5 * 3)
	case Left:
		x = (cont.X)
		y = (cont.Y) + (cont.Height / 5)
		width = (cont.Width / 5) * 3
		height = (cont.Height / 5) * 3
	case Right:
		x = (cont.X) + (cont.Width/5)*4
		y = (cont.Y) + (cont.Height / 5)
		width = (cont.Width / 5) * 3
		height = (cont.Height / 5) * 3
	case Top:
		x = (cont.X)
		y = (cont.Y)
		width = cont.Width
		height = (cont.Height / 5)
	case Bottom:
		x = (cont.X)
		y = (cont.Y) + (cont.Height/5)*4
		width = cont.Width
		height = (cont.Height / 5)
	}

	return x, y, width, height
}

func (layout *BaseLayout) shrinkAt(element Element, position BasePosition) {
	for keyPosition, elementAtPosition := range layout.Elements {
		if elementAtPosition == element {
			box := element.GetBox()
			x, y, width, height := layout.getBoxForPosition(position)
			px, py, pwidth, pheight := layout.getBoxForPosition(keyPosition)

			if x + width > box.X {
				box.X = px
			}

			if y + height > box.Y {
				box.Y = py
			}

			if x < box.X + box.Width {
				box.Width = pwidth
			}

			if y < box.Y + box.Height {
				box.Height = pheight
			}
			element.SetBox((x), (y), (width), (height))
		}
	}
}

func (layout *BaseLayout) isPositionFree(position BasePosition) bool {
	_, ok := layout.Elements[position]
	return !ok
}

func (layout *BaseLayout) expandTo(element Element, position BasePosition) {
	for _, elementAtPosition := range layout.Elements {
		if elementAtPosition == element {
			box := element.GetBox()
			x, y, width, height := layout.getBoxForPosition(position)
			if x < box.X {
				box.X = x
			}

			if y < box.Y {
				box.Y = y
			}

			if x + width > box.X + box.Width {
				box.Width += x + width - box.X - box.Width
			}

			if y + height > box.Y + box.Height {
				box.Height +=  y + height - box.Y - box.Height
			}
			element.SetBox((x), (y), (width), (height))
		}
	}
}

func (layout *BaseLayout) AddElement(element Element, position Position) error {
	basePosition, ok := position.(BasePosition)
	if !ok {
		return errors.New("position is not of type BasePosition")
	}

	x, y, width, height := layout.getBoxForPosition(basePosition)
	element.SetBox((x), (y), (width), (height))
	layout.Elements[basePosition] = element

	//positions := []BasePosition{Center, Top, Bottom, Left, Right}

	//for _, lElement := range layout.Elements {
	//	for _, pos := range positions {
	//		if pos != basePosition{
	//			if layout.isPositionFree(pos) {
	//				layout.expandTo(lElement, pos)
	//			} else {
	//				layout.shrinkAt(lElement, pos)
	//			}
	//		}
	//	}
	//}

	return nil
}

func (layout *GridLayout) GetElements() []Element {
	elements := []Element{}
	for _, w := range layout.Elements {
		elements = append(elements, w)
	}
	return elements
}

func (layout *BaseLayout) GetElements() []Element {
	elements := []Element{}
	for _, w := range layout.Elements {
		elements = append(elements, w)
	}
	return elements
}
