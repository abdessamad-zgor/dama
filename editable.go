package dama

import (
	_ "github.com/abdessamad-zgor/dama/context"
	_ "github.com/abdessamad-zgor/dama/logger"

	"strings"

	_ "github.com/gdamore/tcell/v2"
)

type Cursor struct {
	Column int
	Line   int
}

type DamaEditable interface {
	RemoveRune()
	AddRune(char rune)
	GetCursor() *Cursor
	MoveCursor(direction Direction)
	GetContents() string
    GetLines() []string
	SetEditable(editable bool)
}

type Editable struct {
	Element  Element
	Cursor   *Cursor
	Contents string
}

func NewEditable() *Editable {
	editable := Editable{}
	editable.Cursor = new(Cursor)
	return &editable
}

func (editable *Editable) RemoveRune() {
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

func (editable *Editable) AddRune(char rune) {
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

func (editable *Editable) GetCursor() *Cursor {
	return editable.Cursor
}

func (editable *Editable) MoveCursor(direction Direction) {
	if editable.Cursor != nil {
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
}

func (editable *Editable) GetContents() string {
    return editable.Contents
}

func (editable *Editable) GetLines() []string {
	lines := strings.Split(editable.Contents, "\n")
	return lines
}

func (_editable *Editable) SetEditable(editable bool) {
	if editable {
		_editable.Cursor = new(Cursor)
	} else {
		_editable.Cursor = nil
	}
}
