package traits

import (
	"strings"
)

type Cursor struct {
	Column int
	Line   int
}

type Editable interface {
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

func (editable *editable_s) GetContents() string {
    return editable.Contents
}

func (editable *editable_s) GetLines() []string {
	lines := strings.Split(editable.Contents, "\n")
	return lines
}
