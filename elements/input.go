package elements

import (
	"github.com/abdessamad-zgor/dama"
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
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
	input.Mode = dama.NoMode

	input.BorderColor(tcell.ColorDefault)
	return input
}

func (input *Input) GetBox() dama.Box {
	box := input.Widget.GetBox()
	box.Element = input
	return box
}

func (input *Input) Focus() {
	input.BorderColor(tcell.ColorGrey)
	input.Focused = true
}

func (input *Input) Blur() {
	input.BorderColor(tcell.ColorDefault)
	input.Focused = false 
}

func (input *Input) OnKeyRune(context lcontext.Context, eevent event.KeyEvent) {
	input.Editable.OnKeyRune(context, eevent)
	if input.Mode == dama.NormalMode {
		input.BorderColor(tcell.ColorLime)
	} else if input.Mode == dama.InsertMode {
		input.BorderColor(tcell.ColorBlue)
	} else if input.Mode == dama.VisualMode  {
		input.BorderColor(tcell.ColorYellow)
	} else if input.Mode == dama.NoMode {
		input.BorderColor(tcell.ColorGrey)
	}
}

func (input *Input) Render(screen tcell.Screen, context lcontext.Context) {
	box := input.GetBox()
	box.Render(screen, context)
	input.RenderTag(screen)
	input.RenderTitle(screen)

	if input.GetMode() == dama.NoMode {
		hint := "(Press Enter for normal mode)"
		hintText := dama.Text{hint, box.X + box.Width - len(hint) - 1, box.Y, len(hint), 1}
		hintText.Render(screen)
	}

	text := dama.Text{input.Contents, box.X + 1, box.Y + 1, box.Width - 1, box.Height - 1}
	text.Render(screen)

	input.SetKeybinding(tcell.KeyRune, input.OnKeyRune)
	input.SetKeybinding(tcell.KeyCR, input.OnCarriageReturn)
	input.SetKeybindings(input.OnArrowKeys, tcell.KeyUp, tcell.KeyDown, tcell.KeyLeft, tcell.KeyRight)

	if input.IsFocused() {
		screen.ShowCursor(input.Cursor.Column+1+int(box.X), input.Cursor.Line+1+int(box.Y))
		screen.SetCursorStyle(tcell.CursorStyleDefault)
	}
}
