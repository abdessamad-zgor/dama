package tests

import (
	"testing"
	"github.com/abdessamad-zgor/dama/keystroke"
)

func TestKeystrokeParser(t *testing.T) {
	charMatcher, err := keystroke.GetMatcher("<char>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match := charMatcher("abcd")
	t.Log(match)
	numMatcher, err := keystroke.GetMatcher("<num>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = numMatcher("1234")
	t.Log(match)
	textMatcher, err := keystroke.GetMatcher("<text>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = textMatcher("ewwwww")
	t.Log(match)
}
