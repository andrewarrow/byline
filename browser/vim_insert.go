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
	s := v.getLine()
	prefix = s[0 : v.X+v.FocusLevel]
	suffix = s[v.X+v.FocusLevel:]

	v.DebugLine = fmt.Sprintf("|%s|%s|%d|%d", prefix, suffix, v.X, v.X-v.FocusLevel)
	//|div_p-|3|1|6
	//|__div_bg-red-900_p-3_|rounded_|3|19

	if k == "Backspace" {
		v.SavedLines[v.Y+v.FocusStart+v.Offset] = prefix[0:len(prefix)-1] + suffix
		v.X--
		return
	}

	v.X++
	v.SavedLines[v.Y+v.FocusStart+v.Offset] = prefix + k + suffix
}

func (v *Vim) FocusLevelSpaces() string {
	buffer := []string{}
	for i := 0; i < v.FocusLevel; i++ {
		buffer = append(buffer, " ")
	}
	return strings.Join(buffer, "")
}
