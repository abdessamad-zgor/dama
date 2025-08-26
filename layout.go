package dama

import (
	"errors"
	dtraits "github.com/abdessamad-zgor/dama/traits"
)

type DamaLayout interface {
	GetElements() []DamaElement
	AddElement(element DamaElement, position Position) error
}

type Position interface{}

type GridLayout struct {
	Container *Container
	Columns   int
	Rows      int
	Elements  map[GridPosition]DamaElement
}

func NewGridLayout(columns int, rows int) *GridLayout {
	grid := new(GridLayout)
	grid.Columns = int(columns)
	grid.Rows = int(rows)
	grid.Elements = make(map[GridPosition]DamaElement)
	return grid
}

type GridPosition struct {
	Column     int
	Row        int
	RowSpan    int
	ColumnSpan int
}

type BaseLayout struct {
	Container *Container
	Elements  map[BasePosition]DamaElement
}

type BasePosition = dtraits.Direction

const (
	Center dtraits.Direction = "center"
)

func (layout *GridLayout) AddElement(element DamaElement, position Position) error {
	gridPosition, ok := position.(GridPosition)
	if !ok {
		return errors.New("position is not of type GridPosition")
	}

	if gridPosition.Row >= layout.Rows || gridPosition.Column >= layout.Columns || gridPosition.Column+gridPosition.ColumnSpan > layout.Columns || gridPosition.Row+gridPosition.RowSpan > layout.Rows {
		return errors.New("widget position is out of bounds")
	}

	x := (layout.Container.X) + int(layout.Container.Width/layout.Columns)*gridPosition.Column
	y := (layout.Container.Y) + int(layout.Container.Height/layout.Rows)*gridPosition.Row
	width := int(layout.Container.Width/layout.Columns) * gridPosition.ColumnSpan
	height := int(layout.Container.Height/layout.Rows) * gridPosition.RowSpan

	element.SetBox(x, y, width, height)

	layout.Elements[gridPosition] = element
	return nil
}


func (layout *BaseLayout) getBoxForPosition(position BasePosition) (int, int, int, int) {
	x, y, width, height := (0), (0), (0), (0)
	switch position {
	case Center:
		x = (layout.Container.X + layout.Container.Width/5)
		y = (layout.Container.Y + layout.Container.Height/5)
		width = (layout.Container.Width / 5 * 3)
		height = (layout.Container.Height / 5 * 3)
	case dtraits.Left:
		x = (layout.Container.X)
		y = (layout.Container.Y) + (layout.Container.Height / 5)
		width = (layout.Container.Width / 5) * 3
		height = (layout.Container.Height / 5) * 3
	case dtraits.Right:
		x = (layout.Container.X) + (layout.Container.Width/5)*4
		y = (layout.Container.Y) + (layout.Container.Height / 5)
		width = (layout.Container.Width / 5) * 3
		height = (layout.Container.Height / 5) * 3
	case dtraits.Top:
		x = (layout.Container.X)
		y = (layout.Container.Y)
		width = layout.Container.Width
		height = (layout.Container.Height / 5)
	case dtraits.Bottom:
		x = (layout.Container.X)
		y = (layout.Container.Y) + (layout.Container.Height/5)*4
		width = layout.Container.Width
		height = (layout.Container.Height / 5)
	}

	return x, y, width, height
}

func (layout *BaseLayout) shrinkAt(element DamaElement, position BasePosition) {
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

func (layout *BaseLayout) expandTo(element DamaElement, position BasePosition) {
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

func (layout *BaseLayout) AddElement(element DamaElement, position Position) error {
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
