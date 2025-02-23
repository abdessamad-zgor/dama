package core

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/gdamore/tcell/v2"
)

type DamaApp interface {
	DamaContainer
	Init() error
	Render()
	StartEventLoop()
}

type App struct {
	EventChannel chan event.Event
	Screen       tcell.Screen
	*Container
	event.EventMap
	event.Keybindings
	lcontext.Context
}

func (app *App) Init() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	app.Screen = screen
	if err = screen.Init(); err != nil {
		return err
	}
	width, height := app.Screen.Size()
    app.EventChannel = make(chan event.Event)
	app.X = 0
	app.Y = 0
	app.Width = uint(width)
	app.Height = uint(height)
	return nil
}

func (app *App) StartEventLoop() {

}

func (app *App) Render() {
    widgets := app.GetWidgets()
    for _, widget := range widgets {
        widget
    }
}
