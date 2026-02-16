package tests

import (
	"fmt"
	"testing"
	"time"
	"github.com/abdessamad-zgor/dama"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/keybinding"
)

func TestKeybindingEvent(t *testing.T) {
	widget1 := dama.NewWidget()
	widget1.SetTag('A')
	widget1.SetTitle("Widget 1")
	widget1.SetKeybinding("*", func (e keybinding.Match) {
		logger.Log(fmt.Sprintf("From keybinding callback: %+v", e))
	})
	app.AddElement(widget1, dama.Center)
	go app.Start()
	word := "Hello World!"
	for _, letter := range word {
		event := tcell.NewEventKey(tcell.KeyRune, rune(letter), tcell.ModNone)
		app.GetScreen().PostEvent(event)
		time.Sleep(500 * time.Millisecond)
	}
	app.Exit()
}
