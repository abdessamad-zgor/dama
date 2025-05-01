package elements

import (
	"github.com/abdessamad-zgor/dama"
	_ "github.com/abdessamad-zgor/dama/event"
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

func (textarea *TextArea) Focus() {
	textarea.BorderColor(tcell.ColorGrey)
	textarea.Focused = true
}

func (textarea *TextArea) Blur() {
	textarea.BorderColor(tcell.ColorDefault)
	textarea.Focused = false
}

func (textarea *TextArea) Render(screen tcell.Screen, context lcontext.Context) {
	box := textarea.GetBox()
	box.Render(screen, context)
	textarea.RenderTag(screen)
	textarea.RenderTitle(screen)

	text := dama.Text{textarea.Contents, box.X + 1, box.Y + 1, box.Width - 1, box.Height - 1 }
	text.Render(screen)

	textarea.SetKeybinding(tcell.KeyCR, textarea.OnCarriageReturn)
	textarea.SetKeybindings(textarea.OnArrowKeys, tcell.KeyUp, tcell.KeyDown, tcell.KeyLeft, tcell.KeyRight)

	if textarea.IsFocused() {
		screen.ShowCursor(textarea.Cursor.Column+1+int(box.X), textarea.Cursor.Line+1+int(box.Y))
		screen.SetCursorStyle(tcell.CursorStyleDefault)
	}
}
