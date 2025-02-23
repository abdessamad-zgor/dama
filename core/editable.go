package core

import "strings"

type Cursor struct {
	Column int
	Line   int
}

type Editable interface {
	RemoveRune()
	AddRune(char rune)
	GetCursor() Cursor
	MoveCursor(direction Direction)
	GetContent() []string
}

type Buffer struct {
	Cursor   Cursor
	Contents string
}

func (buffer *Buffer) RemoveRune() {
    i, line := buffer.Cursor.Column, buffer.Cursor.Line
	content := buffer.Contents
    contentLines := strings.Split(content, "\n")
	contentRunes := []rune(contentLines[line])
	contentRunes = append(contentRunes[0:i-1], contentRunes[i:]...)
}

func (buffer *Buffer) AddRune(char rune) {
    i, line := buffer.Cursor.Column, buffer.Cursor.Line
	content := buffer.Contents
    contentLines := strings.Split(content, "\n")
	contentRunes := []rune(contentLines[line])
	contentRunes = append(contentRunes[:i+1], contentRunes[i:]...)
	contentRunes[i] = char
    contentLines[i] = string(contentRunes)
	buffer.Contents = strings.Join(contentLines, "\n")
}

func (buffer *Buffer) GetCursor() Cursor {
	return buffer.Cursor
}

func (buffer *Buffer) MoveCursor(direction Direction) {
	switch direction {
	case Left:
	case Right:
	case Top:
	case Bottom:
	}
}

func (buffer *Buffer) GetContent() []string {
    lines := strings.Split(buffer.Contents, "\n")
    return lines
}
