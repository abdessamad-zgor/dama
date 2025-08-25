package traits

import (
	"strings"
)

type EditMode string

const (
	NoMode 		EditMode = "no-mode"
	InsertMode 	EditMode = "insert"
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
	SetMode(mode EditMode)
	GetMode() EditMode
}

type Editable struct {
	Cursor   Cursor
	Contents string
	Mode     EditMode
}

func NewEditable() *Editable {
	editable := Editable{}
	editable.Mode = NoMode
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

func (editable *Editable) SetMode(mode EditMode) {
	editable.Mode = mode
}

func (editable *Editable) GetMode() EditMode {
	return editable.Mode
}
