package dama

import (
	"errors"
	"fmt"
	"reflect"

	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type DamaContainer interface {
	DamaLayout
	DamaElement
	GetLayout() DamaLayout
	SetLayout(layout DamaLayout) error
	GetNavigables() []DamaElement
}

type Container struct {
	*Element
	Parent DamaContainer
	Layout DamaLayout
}

func NewContainer() *Container {
	container := new(Container)
	container.Element = new(Element)
	layout := new(BaseLayout)
	layout.Elements = make(map[BasePosition]DamaElement)
	layout.Container = container
	container.Layout = layout
	return container
}

func (container *Container) GetLayout() DamaLayout {
	return container.Layout
}

func (container *Container) SetLayout(layout DamaLayout) error {
	glayout, gOk := layout.(*GridLayout)
	blayout, bOk := layout.(*BaseLayout)
	if gOk {
		glayout.Container = container
		container.Layout = glayout
		return nil
	} else if bOk {
		blayout.Container = container
		container.Layout = blayout
		return nil
	} else {
		return errors.New(fmt.Sprintf("invalid layout %v", layout))
	}
}

func (container *Container) GetElements() []DamaElement {
	return container.Layout.GetElements()
}

func (container *Container) AddElement(element DamaElement, position Position) error {
	return container.Layout.AddElement(element, position)
}

func (container *Container) GetBox() Box {
	box := container.Element.GetBox()
	box.Element = container
	return box
}

func (container *Container) Render(screen tcell.Screen, context lcontext.Context) {
	elements := container.GetElements()
	logger.Logger.Println("elements: ", elements)
	for _, element := range elements {
		logger.Logger.Println("element type: ", reflect.TypeOf(element))
		element.Render(screen, context)
	}
}

func (container *Container) GetNavigables() []DamaElement {
	elements := container.GetElements()
	navigables := []DamaElement{}
	for _, element := range elements {
		if element.IsNavigable() {
			navigables = append(navigables, element)
		}
	}
	return navigables
}
