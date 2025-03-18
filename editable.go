package dama

import "strings"

type Cursor struct {
	Column int
	Line   int
}

type DamaEditable interface {
	RemoveRune()
	AddRune(char rune)
	GetCursor() Cursor
	MoveCursor(direction Direction)
	GetContent() []string
}

type Editable struct {
	Cursor   Cursor
	Contents string
}

func (editable *Editable) RemoveRune() {
    i, line := editable.Cursor.Column, editable.Cursor.Line
	content := editable.Contents
    contentLines := strings.Split(content, "\n")
	contentRunes := []rune(contentLines[line])
	contentRunes = append(contentRunes[0:i-1], contentRunes[i:]...)
}

func (editable *Editable) AddRune(char rune) {
    i, line := editable.Cursor.Column, editable.Cursor.Line
	content := editable.Contents
    contentLines := strings.Split(content, "\n")
	contentRunes := []rune(contentLines[line])
	contentRunes = append(contentRunes[:i+1], contentRunes[i:]...)
	contentRunes[i] = char
    contentLines[i] = string(contentRunes)
	editable.Contents = strings.Join(contentLines, "\n")
}

func (editable *Editable) GetCursor() Cursor {
	return editable.Cursor
}

func (editable *Editable) MoveCursor(direction Direction) {
	switch direction {
	case Left:
	case Right:
	case Top:
	case Bottom:
	}
}

func (editable *Editable) GetContent() []string {
    lines := strings.Split(editable.Contents, "\n")
    return lines
}
