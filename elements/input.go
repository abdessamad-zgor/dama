package elements

import (
	"github.com/abdessamad-zgor/dama"
	lcontext "github.com/abdessamad-zgor/dama/context"
	_ "github.com/abdessamad-zgor/dama/event"
	_ "github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Input struct {
	*dama.Widget
	*dama.Editable
	*dama.Scrollable
}


func NewInput() *Input {
	input := &Input{
		dama.NewWidget(),
		dama.NewEditable(),
		new(dama.Scrollable),
	}
	input.SetEditable(true)

	input.AttachModeSwitching(input)

	input.BorderColor(tcell.ColorDefault)
	return input
}

func (input *Input) GetBox() dama.Box {
	box := input.Widget.GetBox()
	box.Element = input
	return box
}

func (input *Input) Focus() {
	input.BorderColor(tcell.ColorLime)
}

func (input *Input) Blur() {
	input.BorderColor(tcell.ColorDefault)
}

func (input *Input) Render(screen tcell.Screen, context lcontext.Context) {
	box := input.GetBox()
	box.Render(screen, context)
	input.RenderTag(screen)
	input.RenderTitle(screen)

	text := dama.Text{input.Contents, &box}
	text.Render(screen)
	screen.ShowCursor(input.Cursor.Column+1+int(box.X), input.Cursor.Line+1+int(box.Y))

	screen.SetCursorStyle(tcell.CursorStyleDefault)
}
