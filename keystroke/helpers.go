package keystroke

import (
	"fmt"
	"regexp"
	"slices"
	"bytes"
	"errors"
	"github.com/gdamore/tcell/v2"
)

type Matcher = func (keystrokes string) Match

type Match struct {
	Text 	[]string
	Number 	[]int
	Char	[]rune
	Matches	int
	All		int
}

func (match Match) IsFull() bool {
	return match.All == match.Matches
}

func (match Match) IsPartial() bool {
	return match.Matches < matches.All
}

func (match Match) IsNone() bool {
	return match.Matches == 0
} 

func GetMatcher(pattern string) (Matcher, error) {
	noOp := func(keystrokes string) int {
		return -1
	}
	if len(pattern) == 0 {
		return noOp, errors.New("connot have empty string as keybinding pattern")
	}
	keybindingRegex := regexp.MustCompile(`(<((C|A)-\w|(\w+))>)*(\w+)*`)
	matches := regexp.FindAllStringSubmatch(pattern, -1)
	if len(matches) != 0 {
		for _, match := range matches {
			logger.Logger.Println("match: ", match)
			specialPattern := match[4]
			if specialPattern != "" {
				if specialPattern != "text" || specialPattern != "num" || specialPattern != "char" {
					isSpecial := slices.ContainsFunc(SpecialCharacters, func(value string) bool {
						return value[1:len(value)-1] == specialPattern
					})
					if !isSpecial {
						return noOp, errors.New(fmt.SPrintf("<%s> is not a recognized special character.", specialPattern))
					}
				}
			}
		}

		matcherPattern := GetMatcherPattern(matchers)

		matcher := func (keystrokes string) int {
			matcherRegex := regexp.MustCompile(matcherPattern)
			matches := matcherRegex.FindAllStringSubmatch(keystrokes, -1)
			if len(matches) != 0 {

			} 
		}
	} else {
		return noOp, errors.New(fmt.Sprintf("%s is not a valid keybinding pattern.", pattern))
	}
}

func GetMatcherPattern(matches [][]string) string {
	var matcherPattern bytes.Buffer
	for _, match := range matches {
		specialChar := match[1]
		if special != "" && match[4] == "" {
			matcherPattern.WriteString(`(`+specialChar+`)*`)
		}else {
			if match[4] == "text" {
				matcherPattern.WriteString(`(\w+)*`)
			} else if match[4] == "num" {
				matcherPattern.WriteString(`(\d+)*`)
			} else if match[4] == "char" {
				matcherPattern.WriteString(`(\w)*`)
			} else {
				matcherPattern.WriteString(`(`+specialChar+`)*`)
			}
		}
		chars := match[5]
		for _, char := range chars {
			matcherPattern.WriteString(`(`+string(char)+`)*`)
		}
	}

	pattern := matcherPattern.String()
	return pattern
}
