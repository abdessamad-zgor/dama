package dama

import (
	"github.com/gdamore/tcell/v2"
)

type Text struct {
    Text 	string
    X 		int
	Y 		int
	Width 	int
	Height 	int
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
        charX := int(text.X+(column))
        charY := int(text.Y+(line))
        if charX > (text.X+text.Width) || charY > (text.Y+text.Height)  {
            continue;
        }
        screen.SetContent(charX, charY, char, nil, tcell.StyleDefault)
        column +=1
    }
}
