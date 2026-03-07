package dama

import (
	"errors"
	"fmt"
	_ "reflect"

	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Container interface {
	Layout
	Element
	GetLayout() Layout
	SetLayout(layout Layout) error
	GetElements() []Element
}

type container_s struct {
	Element
	Parent Container
	Layout Layout
}

func NewContainer() Container {
	container := new(container_s)
	container.Element = NewElement()
	layout := new(BaseLayout)
	layout.Elements = make(map[BasePosition]Element)
	layout.Container = container
	container.Layout = layout
	return container
}

func (container *container_s) GetLayout() Layout {
	return container.Layout
}

func (container *container_s) SetLayout(layout Layout) error {
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

func (container *container_s) GetElements() []Element {
	return container.Layout.GetElements()
}

func (container *container_s) AddElement(element Element, position Position) error {
	return container.Layout.AddElement(element, position)
}

func (container *container_s) GetBox() Box {
	box := container.Element.GetBox()
	box.Element = container
	return box
}

func (container *container_s) Render(screen tcell.Screen) {
	elements := container.GetElements()
	for _, element := range elements {
		logger.Log(element)
		element.Render(screen)
	}
}
