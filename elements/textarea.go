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
	box := textarea.GetBox()
	box.Render(screen, context)
	textarea.RenderTag(screen)
	textarea.RenderTitle(screen)

	text := dama.Text{textarea.Contents, &box}
	text.Render(screen)
	screen.ShowCursor(textarea.Cursor.Column+1+int(box.X), textarea.Cursor.Line+1+int(box.Y))

	screen.SetCursorStyle(tcell.CursorStyleDefault)

}
