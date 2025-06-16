package kparse

import (
	"regexp"
	"github.com/gdamore/tcell/v2"
)

type Keystroke string

type KeyToken struct {
	Modifier 	string
	Char 		rune
	Index		int
}

type KeystrokeParser func(string, []KeyToken) (string, []KeyToken)

const pCtrl KeystrokeParser = func (input string, tokens []KeyToken) (string, []KeyToken) {
	ctrlRegex := regexp.MustCompile(`<C-(\w)>`)
	matches := ctrlRegex.FindAllStringSubmatchIndex(input)
	if len(matches) != 0 {
		for match := range matches {
			 tokens = append(tokens, 
		}
	}
}
