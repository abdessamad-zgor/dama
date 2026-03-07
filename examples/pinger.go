package main

import (
	"fmt"

	"github.com/abdessamad-zgor/dama"
	"github.com/abdessamad-zgor/dama/keybinding"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/abdessamad-zgor/dama/traits"
)

type Input struct {
	dama.Widget
	traits.Editable
}

func main() {
	app, err := dama.NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app due to : %v\n", err))
	}

	url := Input {
		dama.NewWidget(),
		traits.NewEditable(),
	}
	url.SetTag('U')
	url.SetTitle("URL")
	url.SetModeKeybinding(event.InsertMode, "*", func (match keybinding.Match) {
		url.AddRune([]rune(match.Matched)[0])
	})

	url.SetKeybinding("<Up>", func (match keybinding.Match) {
		_ = match
		url.MoveCursor(traits.Top)
	})

	url.SetKeybinding("<Down>", func (match keybinding.Match) {
		_ = match
		url.MoveCursor(traits.Bottom)
	})

	url.SetKeybinding("<Left>", func (match keybinding.Match) {
		_ = match
		url.MoveCursor(traits.Left)
	})

	url.SetKeybinding("<Right>", func (match keybinding.Match) {
		_ = match
		url.MoveCursor(traits.Right)
	})

	output := Input {
		dama.NewWidget(),
		traits.NewEditable(),
	}
	output.SetTag('O')
	output.SetTitle("Output")

	output.SetKeybinding("<Up>", func (match keybinding.Match) {
		_ = match
		output.MoveCursor(traits.Top)
	})

	output.SetKeybinding("<Down>", func (match keybinding.Match) {
		_ = match
		output.MoveCursor(traits.Bottom)
	})

	output.SetKeybinding("<Left>", func (match keybinding.Match) {
		_ = match
		output.MoveCursor(traits.Left)
	})

	output.SetKeybinding("<Right>", func (match keybinding.Match) {
		_ = match
		output.MoveCursor(traits.Right)
	})

	app.AddElement(url, traits.Top)
	app.AddElement(output, traits.Center)

	app.Start()
}
