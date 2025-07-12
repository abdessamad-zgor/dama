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
	go func(s tcell.Screen) {
		for {
			logger.Logger.Println("before failure")
			s.Show()
			logger.Logger.Println("after failure")
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				key := ev.Key()
				switch key {
				case tcell.KeyRune:
					logger.Logger.Println("key: ", key)
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					logger.Logger.Println("key: ", key)
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
