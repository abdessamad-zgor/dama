package tests

import (
	"testing"
	"github.com/samazee/dama"
)

func TestKeybindingParser(t *testing.T) {
	charMatcher, err := dama.GetMatcher("<char>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match := charMatcher("abcd")
	t.Log(match)
	numMatcher, err := dama.GetMatcher("<num>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = numMatcher("1234")
	t.Log(match)
	textMatcher, err := dama.GetMatcher("<text>")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = textMatcher("ewwwww")
	t.Log(match)

	aMatcher, err := dama.GetMatcher("a")
	if err != nil {
		t.Errorf("Failed to create matcher due to %v", err)
	}
	match = aMatcher("a")
	t.Log(match)
}
