package elements


import (
	"github.com/abdessamad-zgor/dama"
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/gdamore/tcell/v2"
)


type TextArea struct {
	*dama.Widget
	*dama.Editable
	*dama.Scrollable
}

func NewTextArea() *TextArea {
	area := TextArea {
		dama.NewWidget(),
		dama.NewEditable(),
		new(dama.Scrollable),
	}
	
	area.BorderColor(tcell.ColorDefault)
	return &area
}

func (textarea *TextArea) Render(screen tcell.Screen, context lcontext.Context) {
	 
}
