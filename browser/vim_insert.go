package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) Replace(k string) {
	s := v.getLine()
	prefix := s[0 : v.X+v.FocusLevel]
	suffix := s[v.X+v.FocusLevel:]
	v.SavedLines[v.Y+v.FocusStart+v.Offset] = prefix[0:len(prefix)] + k + suffix[1:]
}

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
	newLine := prefix + k + suffix
	tokens := strings.Split(strings.TrimSpace(newLine), " ")
	last := tokens[len(tokens)-1]
	fmt.Println(tokens, last, len(last))
	menu := NewMenu(last)
	Document.RenderToId("menu", "menu", menu)
	v.Menu.Show()
	v.SavedLines[v.Y+v.FocusStart+v.Offset] = newLine
}

func (v *Vim) FocusLevelSpaces() string {
	buffer := []string{}
	for i := 0; i < v.FocusLevel; i++ {
		buffer = append(buffer, " ")
	}
	return strings.Join(buffer, "")
}
