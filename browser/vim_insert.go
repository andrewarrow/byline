package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) Insert(k string) {
	//chars := []rune(v.Lines[v.Y])
	//replacementRune := []rune(k)[0]
	//chars[v.X] = replacementRune
	if k == "ArrowUp" || k == "ArrowDown" || k == "ArrowLeft" || k == "ArrowRight" {
		return
	}
	if k == "Enter" {
		return
	}

	prefix := ""
	suffix := ""
	correctLength := 0
	s := v.getLine()
	if v.FocusLevel == 0 {
		prefix = s[0:v.X]
		suffix = s[v.X:]
		v.DebugLine = fmt.Sprintf("|%s|%s|%d|%d", prefix, suffix, correctLength, v.X)
	} else {
		current := getSpaces(v.OffsetLines[v.Y])
		correct := current + v.FocusLevelSpaces()
		correctLength = len(correct) + 1

		prefix = s[0 : correctLength+v.X-1]
		suffix = s[correctLength-1+v.X:]
	}

	//|div_p-|3|1|6
	//|__div_bg-red-900_p-3_|rounded_|3|19

	if k == "Backspace" {
		v.SavedLines[v.Y+v.FocusStart] = prefix + suffix
		v.X--
		return
	}

	v.X++
	v.SavedLines[v.Y+v.FocusStart] = prefix + k + suffix
}

func (v *Vim) FocusLevelSpaces() string {
	buffer := []string{}
	for i := 0; i < v.FocusLevel; i++ {
		buffer = append(buffer, " ")
	}
	return strings.Join(buffer, "")
}
