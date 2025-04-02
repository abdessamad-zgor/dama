package dama

import (
	"github.com/gdamore/tcell/v2"
)

type Text struct {
    Text string
    Box  *Box
}

func (text Text) Render(screen tcell.Screen) {
    runes := []rune(text.Text)
    line := 0
    column := 0
    for _, char := range runes {
        if char =='\n' {
            line += 1
            column = 0
        }
        charX := int(text.Box.X+uint(column+1))
        charY := int(text.Box.Y+uint(line+1))
        if charX >= int(text.Box.Width) || charY >= int(text.Box.Height)  {
            continue;
        }
        screen.SetContent(charX, charY, char, nil, tcell.StyleDefault)
        column +=1
    }
}
