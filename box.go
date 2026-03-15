package dama

import (
	"fmt"
	"github.com/samazee/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	X       int
	Y       int
	Width   int
	Height  int
	Element Element
}

func (box Box) Render(screen tcell.Screen) {
	logger.Log(fmt.Sprintf("Element box: %+v", box.Element))
	boxStyle := box.Element.GetStyleProperties()

	borderStyle := tcell.StyleDefault
	if borderColor, ok := boxStyle.GetBorderColor(); ok {
		borderStyle = borderStyle.Foreground(borderColor)
	}
	if borderBold, ok := boxStyle.GetBorderBold(); ok {
		borderStyle = borderStyle.Bold(borderBold)
	}

	for xi := range box.Width {
		for yi := range box.Height {
			borderX, borderY := xi + box.X, yi + box.Y
			if borderX == box.X || borderX == box.X + box.Width - 1 {
				screen.SetContent(int(borderX), int(borderY), tcell.RuneVLine, nil, borderStyle)
			}
			if borderY == box.Y || borderY == box.Y + box.Height - 1 {
				screen.SetContent(int(borderX), int(borderY), tcell.RuneHLine, nil, borderStyle)
			}
		}
	}

	screen.SetContent(int(box.X), int(box.Y), tcell.RuneULCorner, nil, borderStyle)
	screen.SetContent(int(box.X+box.Width-1), int(box.Y), tcell.RuneURCorner, nil, borderStyle)
	screen.SetContent(int(box.X), int(box.Y+box.Height-1), tcell.RuneLLCorner, nil, borderStyle)
	screen.SetContent(int(box.X+box.Width-1), int(box.Y+box.Height-1), tcell.RuneLRCorner, nil, borderStyle)
}
