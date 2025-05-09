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
}

func KeyToDirection(key tcell.Key) dama.Direction {
	direction := dama.Center
	switch(key) {
	case tcell.KeyUp:
		direction = dama.Top
	case tcell.KeyDown:
		direction = dama.Bottom
	case tcell.KeyLeft:
		direction = dama.Left
	case tcell.KeyRight:
		direction = dama.Right
	default:
		direction = dama.Center
	}
	return direction
}

func NewInput() *Input {
	input := &Input{
		dama.NewWidget(),
		dama.NewEditable(),
		new(dama.Scrollable),
	}

	input.SetEventListener(tcell.KeyRune, event.Key, func(context lcontext.Context, ievent event.Event) {
		keyEvent, _ := ievent.TEvent.(*tcell.EventKey)
		rune := keyEvent.Rune()
		input.AddRune(rune)
		logger.Logger.Println("Inside input listener")
	})

	input.SetEventListener(tcell.KeyCR, event.CR, func(context lcontext.Context, ievent event.Event) {
		input.AddRune('\n')
		logger.Logger.Println("Inside input listener")
	})

	input.SetEventListener(tcell.KeyUp, event.MoveCursor, func(context lcontext.Context, ievent event.Event) {
		keyEvent, _ := ievent.TEvent.(*tcell.EventKey)
		direction := KeyToDirection(keyEvent.Key())
		input.MoveCursor(direction)
	})

	input.SetEventListener(tcell.KeyDown, event.MoveCursor, func(context lcontext.Context, ievent event.Event) {
		keyEvent, _ := ievent.TEvent.(*tcell.EventKey)
		direction := KeyToDirection(keyEvent.Key())
		input.MoveCursor(direction)
	})

	input.SetEventListener(tcell.KeyLeft, event.MoveCursor, func(context lcontext.Context, ievent event.Event) {
		keyEvent, _ := ievent.TEvent.(*tcell.EventKey)
		direction := KeyToDirection(keyEvent.Key())
		input.MoveCursor(direction)
	})

	input.SetEventListener(tcell.KeyRight, event.MoveCursor, func(context lcontext.Context, ievent event.Event) {
		keyEvent, _ := ievent.TEvent.(*tcell.EventKey)
		direction := KeyToDirection(keyEvent.Key())
		input.MoveCursor(direction)
	})

	input.BorderColor(tcell.ColorDefault)
	return input
}

func (input *Input) ContentsToText() dama.Text {
	box := input.GetBox()

	return dama.Text{input.Contents, &box}
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

	text := input.ContentsToText()
	text.Render(screen)
	screen.ShowCursor(input.Cursor.Column+1+int(box.X), input.Cursor.Line+1+int(box.Y))

	screen.SetCursorStyle(tcell.CursorStyleDefault)
}
