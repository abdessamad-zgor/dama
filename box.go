package dama

import (
	_ "github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	X       int
	Y       int
	Width   int
	Height  int
	Element DamaElement
}

func (box Box) Render(screen tcell.Screen) {
	boxStyle := box.Element.GetStyle()

	if boxStyle.Border != nil {
		borderStyle := tcell.StyleDefault.Foreground(boxStyle.Border.Color).Bold(boxStyle.Border.Bold)
		for xi := range box.Width {
			for yi := range box.Height {
				borderX, borderY := xi+box.X, yi+box.Y
				if borderX == box.X || borderX == box.X+box.Width-1 {
					screen.SetContent(int(borderX), int(borderY), tcell.RuneVLine, nil, borderStyle)
				}
				if borderY == box.Y || borderY == box.Y+box.Height-1 {
					screen.SetContent(int(borderX), int(borderY), tcell.RuneHLine, nil, borderStyle)
				}
			}
		}

		screen.SetContent(int(box.X), int(box.Y), tcell.RuneULCorner, nil, borderStyle)
		screen.SetContent(int(box.X+box.Width-1), int(box.Y), tcell.RuneURCorner, nil, borderStyle)
		screen.SetContent(int(box.X), int(box.Y+box.Height-1), tcell.RuneLLCorner, nil, borderStyle)
		screen.SetContent(int(box.X+box.Width-1), int(box.Y+box.Height-1), tcell.RuneLRCorner, nil, borderStyle)
	}

}
