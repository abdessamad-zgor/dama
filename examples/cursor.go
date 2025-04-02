//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault)
	s.Clear()

	text := "This demonstrates cursor styles.  Press 0 through 6 to change the style."
	x := 1
	for _, r := range text {
		s.SetCell(x, 1, tcell.StyleDefault, r)
		x++
	}
	s.SetContent(2, 2, '0', nil, tcell.StyleDefault)
	s.SetCursorStyle(tcell.CursorStyleDefault)
	s.ShowCursor(3, 2)
	quit := make(chan struct{})
	style := tcell.StyleDefault
	go func(s tcell.Screen) {
		for {
            logger.Logger.Println("inside cursor example")
			s.Show()
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyRune:
					switch ev.Rune() {
					case '0':
						s.SetContent(2, 2, '0', nil, style)
						s.SetCursorStyle(tcell.CursorStyleDefault, tcell.ColorReset)
					case '1':
						s.SetContent(2, 2, '1', nil, style)
						s.SetCursorStyle(tcell.CursorStyleBlinkingBlock, tcell.ColorGreen)
					case '2':
						s.SetCell(2, 2, tcell.StyleDefault, '2')
						s.SetCursorStyle(tcell.CursorStyleSteadyBlock, tcell.ColorBlue)
					case '3':
						s.SetCell(2, 2, tcell.StyleDefault, '3')
						s.SetCursorStyle(tcell.CursorStyleBlinkingUnderline, tcell.ColorRed)
					case '4':
						s.SetCell(2, 2, tcell.StyleDefault, '4')
						s.SetCursorStyle(tcell.CursorStyleSteadyUnderline, tcell.ColorOrange)
					case '5':
						s.SetCell(2, 2, tcell.StyleDefault, '5')
						s.SetCursorStyle(tcell.CursorStyleBlinkingBar, tcell.ColorYellow)
					case '6':
						s.SetCell(2, 2, tcell.StyleDefault, '6')
						s.SetCursorStyle(tcell.CursorStyleSteadyBar, tcell.ColorPink)
					}

				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}(s)
	<-quit
	s.Fini()
}
