package dama

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/abdessamad-zgor/dama/logger"

	"strings"

	"github.com/gdamore/tcell/v2"
)

type EditableMode string

const (
	NormalMode EditableMode = "normal"
	VisualMode EditableMode = "visual"
	InsertMode EditableMode = "insert"
)

type Cursor struct {
	Column int
	Line   int
}

type DamaEditable interface {
	RemoveRune()
	AddRune(char rune)
	GetCursor() Cursor
	MoveCursor(direction Direction)
	GetContents() string
    GetLines() []string
	IsEditable() bool
	SetEditable(editable bool)
	SetMode(mode EditableMode)
	GetMode() EditableMode
}

type Editable struct {
	Element  Element
	Cursor   Cursor
	Contents string
	Mode     EditableMode
	Editable bool
}

func NewEditable() *Editable {
	editable := Editable{}
	editable.Mode = NormalMode
	return &editable
}

func (editable *Editable) RemoveRune() {
	if editable.Editable {
		i, line := editable.Cursor.Column, editable.Cursor.Line
		content := editable.Contents
		lines := strings.Split(content, "\n")
		runes := []rune(lines[line])
		if runes[i] == '\n' {
			editable.Cursor.Line -= 1
		}
		runes = append(runes[0:i-1], runes[i:]...)
		editable.Cursor.Column -= 1
	}
}

func (editable *Editable) AddRune(char rune) {
	if editable.Editable {
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
		logger.Logger.Print("line runes at char " + string(char) + ": ", runes)
	}
}

func (editable *Editable) GetCursor() Cursor {
	return editable.Cursor
}

func (editable *Editable) MoveCursor(direction Direction) {
		cursor := editable.Cursor
        lines := strings.Split(editable.Contents, "\n")
		switch direction {
		case Left:
			if cursor.Column > 0 {
				cursor.Column -= 1
			}
		case Right:
			if cursor.Column < len(lines[cursor.Line]) - 1 {
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
}

func (editable *Editable) GetContents() string {
    return editable.Contents
}

func (editable *Editable) GetLines() []string {
	lines := strings.Split(editable.Contents, "\n")
	return lines
}

func (_editable *Editable) SetEditable(editable bool) {
	_editable.Editable = editable
}

func (editable *Editable) IsEditable() bool {
	return editable.Editable 
}

func (editable *Editable) SetMode(mode EditableMode) {
	editable.Mode = mode
}

func (editable *Editable) GetMode() EditableMode {
	return editable.Mode
}

func KeyToDirection(key tcell.Key) Direction {
	direction := Center
	switch(key) {
	case tcell.KeyUp:
		direction = Top
	case tcell.KeyDown:
		direction = Bottom
	case tcell.KeyLeft:
		direction = Left
	case tcell.KeyRight:
		direction = Right
	default:
		direction = Center
	}
	return direction
}

func (editable *Editable) AttachModeSwitching(widget DamaWidget) {
	if editable.Editable {
		widget.SetKeybinding(tcell.KeyEsc, func (context lcontext.Context, eevent event.KeyEvent) {
			if editable.Mode == InsertMode || editable.Mode == VisualMode {
				editable.Mode = NormalMode
			}
		})

		widget.SetKeybinding(tcell.KeyRune, func (context lcontext.Context, eevent event.KeyEvent) {
			keyEvent, _ := eevent.TEvent.(*tcell.EventKey)
			keyRune := keyEvent.Rune()
			if editable.Mode == NormalMode {
				if keyRune == 'i' {
					editable.Mode = InsertMode
				} else if keyRune == 'v' {
					editable.Mode = VisualMode
				}
			} else if editable.Mode == InsertMode {
				editable.AddRune(keyRune)
			}
		})

		widget.SetKeybinding(tcell.KeyCR, func (context lcontext.Context, eevent event.KeyEvent) {
			if editable.Mode == InsertMode {
				editable.AddRune('\n')
			} else {
				editable.MoveCursor(Bottom)
			}
		})

		widget.SetKeybindings(func (context lcontext.Context, eevent event.KeyEvent) {
			direction := KeyToDirection(eevent.Key)
			editable.MoveCursor(direction)
		}, tcell.KeyUp, tcell.KeyDown, tcell.KeyLeft, tcell.KeyRight)
	}
}
