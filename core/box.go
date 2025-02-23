package core

import (
	"github.com/gdamore/tcell/v2"
	lcontext "github.com/abdessamad-zgor/dama/context"
)

type Box struct {
	X      uint
	Y      uint
	Width  uint
	Height uint
    Widget DamaWidget
}

func (box Box) Render(screen tcell.Screen, context lcontext.Context) {
	boxStyle := box.Widget.GetContextStyle(context)
	for xi := range box.Width {
		for yi := range box.Height {
			borderX, borderY := xi+box.X, yi+box.Y
			if borderX == box.X || borderX == box.X+box.Width-1 {
				screen.SetContent(int(borderX), int(borderY), tcell.RuneVLine, nil, boxStyle)
			}
			if borderY == box.Y || borderY == box.Y+box.Height-1 {
				screen.SetContent(int(borderX), int(borderY), tcell.RuneHLine, nil, boxStyle)
			}
		}
	}

	screen.SetContent(int(box.X), int(box.Y), tcell.RuneULCorner, nil, boxStyle)
	screen.SetContent(int(box.X+box.Width-1) , int(box.Y), tcell.RuneURCorner, nil, boxStyle)
	screen.SetContent(int(box.X), int(box.Y+box.Height-1), tcell.RuneLLCorner, nil, boxStyle)
	screen.SetContent(int(box.X+box.Width-1), int(box.Y+box.Height-1), tcell.RuneLRCorner, nil, boxStyle)
}
