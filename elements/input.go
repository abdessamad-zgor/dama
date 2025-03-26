package elements

import "github.com/abdessamad-zgor/dama"

type Input struct {
	*dama.Widget
	*dama.Editable
	*dama.Navigable
	*dama.Scrollable
}

func NewInput(mark rune, title string) *Input {
    return &Input{}
}
