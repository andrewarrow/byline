package browser

import "strings"

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

	s := v.getLineAdjustForFocus()
	v.DebugLine = s
	if k == "Backspace" {
		prefix := s[0 : v.X-1]
		suffix := s[v.X:]
		v.SavedLines[v.Y+v.FocusStart] = prefix + suffix
		v.X--
		return
	}

	prefix := s[0:v.X]
	suffix := s[v.X:]
	v.X++
	v.SavedLines[v.Y+v.FocusStart] = prefix + k + suffix
}

func (v *Vim) FocusLevelSpaces() string {
	buffer := []string{}
	for i := 0; i < v.FocusLevel-2; i++ {
		buffer = append(buffer, " ")
	}
	return strings.Join(buffer, "")
}
