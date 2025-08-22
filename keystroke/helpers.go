package keystroke

import (
	"fmt"
	"regexp"
	"slices"
	"bytes"
	"errors"
	"github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
)

type Matcher = func (keystrokes string) Match

type Match struct {
	Texts 	[]string
	Numbers []int
	Chars	[]rune
	Matches	int
	All		int
}

type MatcherPatternResult struct {
	pattern string
	texts	[]int
	nums	[]int
	chars	[]int
	all		int
}

func (match Match) IsFull() bool {
	utils.Assert(match.All >= match.Matches, "there was something wrong with the counting off all matches")
	return match.All == match.Matches
}

func (match Match) IsPartial() bool {
	utils.Assert(match.All >= match.Matches, "there was something wrong with the counting off all matches")
	return match.Matches < matches.All
}

func (match Match) IsNone() bool {
	utils.Assert(match.All >= match.Matches, "there was something wrong with the counting off all matches")
	return match.Matches == 0
} 

func GetMatcher(pattern string) (Matcher, error) {
	noOp := func(keystrokes string) Match {
		return Match{}
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

		matcherPattern := GetMatcherPattern(matches)
		logger.Logger.Println("matcher pattern: " + matcherPattern)

		matcher := func (keystrokes string) Match {
			match := Match{}
			match.All = matcherPatten.all
			matcherRegex := regexp.MustCompile(matcherPattern.pattern)
			matches := matcherRegex.FindAllStringSubmatch(keystrokes, -1)
			utils.Assert(len(matches) == 1, "keybinding matchers got multiple matches, probably the keystroke buffer didn't get emptied")
			if len(matches) != 0 {
				for i, _m := range matches[0] {
					if _m != "" {
						if slices.Contains(matcherPattern.texts, i) {
							match.Texts = append(match.Texts, _m)
							match.Matches += 1
						} else if slices.Contains(matcherPattern.nums, i) {
							match.Numbers = append(match.Numbers, _m)
							match.Matches += 1
						} else if slices.Contains(matcherPattern.chars, i) {
							match.Chars = append(match.Chars, _m)
							match.Matches += 1
						} else {
							match.Matches += 1
						}
					} else {
						break
					}
				}
			}

			return match
		}

		return matcher, nil
	} else {
		return noOp, errors.New(fmt.Sprintf("%s is not a valid keybinding pattern.", pattern))
	}
}

func GetMatcherPattern(matches [][]string) PatternMatcherResult {
	var matcherPattern bytes.Buffer
	result := PatternMatcherResult{}
	for i, match := range matches {
		specialChar := match[1]
		if special != "" && match[4] == "" {
			matcherPattern.WriteString(`(`+specialChar+`)*`)
			result.all += 1
		}else {
			if match[4] == "text" {
				matcherPattern.WriteString(`(\w+)*`)
				result.texts = append(result.texts, i)
				result.all += 1
			} else if match[4] == "num" {
				matcherPattern.WriteString(`(\d+)*`)
				result.nums = append(result.nums, i)
				result.all += 1
			} else if match[4] == "char" {
				matcherPattern.WriteString(`(\w)*`)
				result.chars = append(result.chars, i)
				result.all += 1
			} else {
				matcherPattern.WriteString(`(`+specialChar+`)*`)
				result.all += 1
			}
		}
		chars := match[5]
		for _, char := range chars {
			matcherPattern.WriteString(`(`+string(char)+`)*`)
			result.all += 1
		}
	}

	pattern := matcherPattern.String()
	result.pattern = pattern
	return result
}
