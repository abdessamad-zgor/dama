package dama

import (
	"strings"
	devent "github.com/abdessamad-zgor/dama/event"
	dkeybinding "github.com/abdessamad-zgor/dama/keybinding"
	"github.com/gdamore/tcell/v2"
)

type Trait interface {
	Render(widget Widget, screen tcell.Screen)
	GetTraitKeybindings() []devent.Event
}

type Cursor struct {
	Column int
	Line   int
}

type Editable interface {
	Trait
	AddRune(char rune)
	RemoveRune()
	GetCursor() Cursor
	MoveCursor(direction Direction)
	GetContents() string
    GetLines() []string
}

type editable_s struct {
	Cursor   Cursor
	Contents string
}

func NewEditable() Editable {
	editable := editable_s{}
	return &editable
}

func (editable *editable_s) RemoveRune() {
	i, line := editable.Cursor.Column - 1, editable.Cursor.Line
	content := editable.Contents
	lines := strings.Split(content, "\n")
	runes := []rune(lines[line])
	if runes[i] == '\n' {
		editable.Cursor.Line -= 1
	}
	runes = append(runes[0:i], runes[i + 1:]...)
	editable.Cursor.Column -= 1
	lines = append(lines[:line], string(runes))
	editable.Contents = strings.Join(lines, "\n")
}

func (editable *editable_s) AddRune(char rune) {
	i, line := editable.Cursor.Column, editable.Cursor.Line
	content := editable.Contents
	lines := strings.Split(content, "\n")
	runes := []rune(lines[line])
	if i >= len(runes) {
		runes = append(runes, char)
		lines[line] = string(runes)
	} else {
		runes = append(runes[:i+1], runes[i:]...)
		runes[i] = char
		lines[line] = string(runes)
	}
	editable.Contents = strings.Join(lines, "\n")
	editable.Cursor.Column += 1
	if char == '\n' {
		editable.Cursor.Line += 1
		editable.Cursor.Column = 0
	}
}

func (editable *editable_s) GetCursor() Cursor {
	return editable.Cursor
}

func (editable *editable_s) MoveCursor(direction Direction) {
	cursor := editable.Cursor
	lines := strings.Split(editable.Contents, "\n")
	switch direction {
	case Left:
		if cursor.Column > 0 {
			cursor.Column -= 1
		}
	case Right:
		if cursor.Column < len(lines[cursor.Line]) {
			cursor.Column += 1
		}
	case Top:
		if cursor.Line > 0 {
			cursor.Line -= 1
			if cursor.Column >= len(lines[cursor.Line]) {
				cursor.Column = len(lines[cursor.Line])
			}
		}
	case Bottom:
		if cursor.Line < len(lines) - 1 {
			cursor.Line += 1
			if cursor.Column >= len(lines[cursor.Line]) {
				cursor.Column = len(lines[cursor.Line])
			}
		}
	}
	editable.Cursor = cursor
}

func (editable *editable_s) GetContents() string {
    return editable.Contents
}

func (editable *editable_s) GetLines() []string {
	lines := strings.Split(editable.Contents, "\n")
	return lines
}

func (editable *editable_s) Render(widget Widget, screen tcell.Screen) {
	widgetBox := widget.GetBox()
	if widget.IsFocused() {
		if widget.GetMode() == devent.InsertMode {
			screen.SetCursorStyle(tcell.CursorStyleSteadyBar)
		} else if widget.GetMode() == devent.NormalMode {
			screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
		}
		screen.ShowCursor(widgetBox.X + editable.Cursor.Column + 1, widgetBox.Y + editable.Cursor.Line + 1)
	}
	text := NewText(widgetBox.X + 1, widgetBox.Y + 1, widgetBox.Width, widgetBox.Height, editable.Contents)
	text.Render(screen)
}

func (editable *editable_s) GetTraitKeybindings() []devent.Event {
	keybindings := []devent.Event{}
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.InsertMode, "*", func (match dkeybinding.Match) {
		editable.AddRune([]rune(match.Matched)[0])
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.InsertMode, "<BS>", func (match dkeybinding.Match) {
		_ = match
		editable.RemoveRune()
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.NormalMode, "<Up>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Top)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.NormalMode, "<Down>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Bottom)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.NormalMode, "<Left>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Left)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.NormalMode, "<Right>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Right)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.InsertMode, "<Up>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Top)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.InsertMode, "<Down>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Bottom)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.InsertMode, "<Left>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Left)
	}))
	keybindings = append(keybindings, devent.KeybindingToEvent(devent.InsertMode, "<Right>", func (match dkeybinding.Match) {
		_ = match
		editable.MoveCursor(Right)
	}))
	return keybindings
}
