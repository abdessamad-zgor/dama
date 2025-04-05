package elements

import (
	"github.com/abdessamad-zgor/dama"
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Input struct {
	*dama.Widget
	*dama.Editable
	*dama.Scrollable
	*dama.Stylable
}

func NewInput() Input {
	input := Input{
		dama.NewWidget(),
		dama.NewEditable(),
		new(dama.Scrollable),
		&dama.Stylable{
			dama.DefaultStyle,
		},
	}

	input.SetEventListener(tcell.KeyRune, event.Key, func(context lcontext.Context, ievent event.Event) {
		keyEvent, _ := ievent.TEvent.(*tcell.EventKey)
		rune := keyEvent.Rune()
		input.AddRune(rune)
	})

	input.SetEventListener(tcell.KeyCR, event.Key, func(context lcontext.Context, ievent event.Event) {
		input.AddRune('\n')
	})

	input.BorderColor(tcell.ColorDefault)
	return input
}

func (input Input) ContentsToText() dama.Text {
	box := input.GetBox()

	return dama.Text{input.Contents, &box}
}

func (input Input) GetBox() dama.Box {
	box := input.Widget.GetBox()
	box.Element = input
	return box
}

func (input Input) Render(screen tcell.Screen, context lcontext.Context) {
	box := input.GetBox()
	logger.Logger.Println("box: ", box)
	box.Render(screen, context)

	text := input.ContentsToText()
	text.Render(screen)
	screen.ShowCursor(input.Cursor.Column+1+int(box.X), input.Cursor.Line+1+int(box.Y))

	screen.SetCursorStyle(tcell.CursorStyleDefault)
}
