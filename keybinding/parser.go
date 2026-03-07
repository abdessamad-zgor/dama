package keybinding

import (
	"fmt"
	"regexp"
	"slices"
	"bytes"
	"errors"
	_ "strconv"
	_ "github.com/abdessamad-zgor/dama/utils"
	"github.com/abdessamad-zgor/dama/logger"
)

type Matcher = func (buffer string) Match

type Match struct {
	Pattern		string
	Matched 	string
	Buffer		string
	CatchAll	map[int]rune
}

func (match Match) IsFull() bool {
	//logger.Log(fmt.Sprintf("IsFull(): match=%+v", match))
	return len(match.Matched) > 0 && len(match.Matched) == len(match.Buffer) && len(match.Matched) == len(match.Pattern)
}

func (match Match) IsPartial() bool {
	//logger.Log(fmt.Sprintf("IsPartial(): match=%+v", match))
	return len(match.Matched) > 0 && len(match.Matched) == len(match.Buffer) && len(match.Matched) < len(match.Pattern)
}

func GetMatcherPattern(keybinding string) (string, error) {
	var matcherPattern bytes.Buffer

	keybindingRegex := regexp.MustCompile(`<(.+)>`)
	matches := keybindingRegex.FindAllSubmatchIndex([]byte(keybinding), -1)
	i := 0
	for i < len(keybinding) {
		isInMatch := slices.IndexFunc(matches, func (match []int) bool {
			return (i >= match[2] && i < match[3])
		})
		if isInMatch == -1 {
			matcherPattern.WriteString(string(keybinding[i]))
			i += 1
		} else {
			inMatch := matches[isInMatch]
			if slices.Contains(SpecialCharacters, keybinding[inMatch[0]: inMatch[1]]) {
				matcherPattern.WriteString(keybinding[inMatch[2]: inMatch[3]])
				i = inMatch[3]
			} else {
				return "", errors.New(fmt.Sprintf("%s is not a valid special character", keybinding[inMatch[0]: inMatch[1]]))
			}
		}
	}

	return matcherPattern.String(), nil
} 

func GetMatcher(pattern string) (Matcher, error) {
	noOp := func(buffer string) Match {
		return Match{}
	}
	if len(pattern) == 0 {
		return noOp, errors.New("connot have empty string as keybinding pattern")
	}

	matcherPattern, err := GetMatcherPattern(pattern)
	if err != nil {
		return noOp, err
	}
	logger.Log("matcher pattern: ", fmt.Sprintf("%+v", matcherPattern))

	matcher := func (buffer string) Match {
		match := Match{CatchAll: make(map[int]rune)}
		match.Buffer = buffer
		match.Pattern = matcherPattern
		for i, char := range buffer {
			if i < len(matcherPattern) {
				if rune(matcherPattern[i]) == '*' {
					match.CatchAll[i] = rune(buffer[i])
					match.Matched += string(buffer[i])
					continue
				} else if char == rune(matcherPattern[i]) {
					match.Matched += string(char)
					continue
				}
			}
			match.Matched = ""
			match.CatchAll = make(map[int]rune) 
			break
		}
		return match
	}
	return matcher, nil
}
