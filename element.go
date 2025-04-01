package dama

import (
    lcontext "github.com/abdessamad-zgor/dama/context"
    "github.com/gdamore/tcell/v2"
)

type DamaElement interface {
    Render(screen tcell.Screen, context lcontext.Context)
    GetBox() Box
    SetBox(x uint, y uint, width uint, height uint)
}

type Element struct {
	X           uint
	Y           uint
	Width       uint
	Height      uint
}

func (element *Element) Render(screen tcell.Screen, context lcontext.Context) {

}

func (element *Element) GetBox() Box {
    return Box{
        element.X,
        element.Y,
        element.Width,
        element.Height,
        nil,
    }
}

func (element *Element) SetBox(x uint, y uint, width uint, height uint) {
    element.X = x
    element.Y = y
    element.Width = width
    element.Height = height
}
