package tests

import (
	"testing"
	"github.com/abdessamad-zgor/dama/keybinding"
)

func TestKeybindingParser(t *testing.T) {
	charMatcher, err := keybinding.GetMatcher("<char>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match := charMatcher("abcd")
	t.Log(match)
	numMatcher, err := keybinding.GetMatcher("<num>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = numMatcher("1234")
	t.Log(match)
	textMatcher, err := keybinding.GetMatcher("<text>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = textMatcher("ewwwww")
	t.Log(match)

	aMatcher, err := keybinding.GetMatcher("a")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = aMatcher("a")
	t.Log(match)
}
