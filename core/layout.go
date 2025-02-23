package core

import (
	"errors"
)

type Position interface{}

type GridLayout struct {
	Columns uint
	Rows    uint
    Widgets map[GridPosition]Widget
}

type GridPosition struct {
	Column     uint
	Row        uint
	RowSpan    uint
	ColumnSpan uint
}

type BaseLayout struct {
    Widgets map[BasePosition]Widget
}

type BasePosition = Direction

const (
    Center BasePosition = "center"
)

type DamaLayout interface {
    GetWidgets() []Widget
    AddWidget(widget Widget, position Position) error
}

func (layout *GridLayout) AddWidget(widget Widget, position Position) error {
    gridPosition, ok := position.(GridPosition)
    if !ok {
        return errors.New("position is not of type GridPosition")
    }
    if gridPosition.Row >= layout.Rows || gridPosition.Column >= layout.Columns || gridPosition.Column + gridPosition.ColumnSpan > layout.Columns || gridPosition.Row + gridPosition.RowSpan > layout.Rows {
        return errors.New("widget position is out of bounds")
    } 
    layout.Widgets[gridPosition] = widget
    return nil 
} 

func (layout *BaseLayout) AddWidget(widget Widget, position Position) error {
    basePosition, ok := position.(BasePosition)
    if !ok {
        return errors.New("position is not of type BasePosition")
    }
    layout.Widgets[basePosition] = widget
    return nil
}

func (layout *GridLayout) GetWidgets() []Widget {
    widgets := []Widget{}
    for _, w := range layout.Widgets {
        widgets = append(widgets, w)
    }
    return widgets
}

func (layout *BaseLayout) GetWidgets() []Widget {
    widgets := []Widget{}
    for _, w := range layout.Widgets {
        widgets = append(widgets, w)
    }
    return widgets
}
