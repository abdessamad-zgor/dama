package dama

type Navigable interface {
	GetTag() rune
	GetTitle() string
}

type Display struct {
	Title string
	Tag   rune
}

func (display Display) GetTitle() string {
    return display.Title
}

func (display Display) GetTag() rune {
    return display.Tag
}
