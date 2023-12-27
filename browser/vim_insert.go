package browser

import (
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

	s := v.getLine()
	current := getSpaces(v.OffsetLines[v.Y])
	correct := current + v.FocusLevelSpaces()
	correctLength := len(correct)

	//delta := v.FocusLevel - len(current) - 2
	//delta = v.X + delta
	//prefix := s[delta:delta+1]
	//suffix := s[delta:]
	prefix := s[0 : correctLength+v.X-1]
	suffix := s[correctLength+v.X:]
	//prefix := v.SavedLines[v.Y+v.FocusStart]
	//v.DebugLine = fmt.Sprintf("%s|%s", prefix, suffix)

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
