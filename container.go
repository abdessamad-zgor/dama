package dama

import (
    lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/gdamore/tcell/v2"
)

type DamaContainer interface {
	DamaLayout
	DamaElement
	GetLayout() DamaLayout
	SetLayout(layout DamaLayout)
}


type Container struct {
	X      uint
	Y      uint
	Width  uint
	Height uint
	Parent DamaContainer
	Layout DamaLayout
}

func CreateContainer() *Container {
	container := new(Container)
	return container
}

func (container *Container) GetLayout() DamaLayout {
	return container.Layout
}

func (container *Container) SetLayout(layout DamaLayout) {
	container.Layout = layout
}

func (container *Container) GetElements() []DamaElement {
	return container.Layout.GetElements()
}

func (container *Container) AddElement(element DamaElement, position Position) error {
	return container.Layout.AddElement(element, position)
}

func (container *Container) GetBox() Box {
	return Box{container.X, container.Y, container.Width, container.Height, container}
}

func (container *Container) Render(screen tcell.Screen, context lcontext.Context) {
	elements := container.GetElements()
	for _, element := range elements {
		element.Render(screen, context)
	}
}
