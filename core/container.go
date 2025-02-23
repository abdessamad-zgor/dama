package core

type DamaContainer interface {
	DamaLayout
	GetLayout() DamaLayout
    SetLayout(layout DamaLayout)
	GetBox() Box
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

func (container *Container) GetLayout() DamaLayout{
	return container.Layout
}

func (container *Container) SetLayout(layout DamaLayout) {
    container.Layout = layout
}

func (container *Container) GetWidgets() []Widget {
	return container.Layout.GetWidgets()
}

func (container *Container) GetBox() Box {
    nilWidget := any(nil).(Widget)
    return Box{container.X, container.Y, container.Width, container.Height, &nilWidget} 
}
