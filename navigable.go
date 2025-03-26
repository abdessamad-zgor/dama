package dama

type DamaNavigable interface {
	GetTag() rune
	GetTitle() string
}

type Navigable struct {
	Title string
	Tag   rune
}

func (navigable Navigable) GetTitle() string {
	return navigable.Title
}

func (navigable Navigable) GetTag() rune {
	return navigable.Tag
}

func (navigable Navigable) SetTitle(title string) {
	navigable.Title = title
}

func (navigable Navigable) SetTag(tag rune) {
	navigable.Tag = tag
}
