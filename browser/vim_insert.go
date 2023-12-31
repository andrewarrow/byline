package browser

import (
	"strings"
)

func (v *Vim) Replace(k string) {
	s := v.getLine()
	prefix := s[0 : v.X+v.FocusLevel]
	suffix := s[v.X+v.FocusLevel:]
	v.SavedLines[v.Y+v.FocusStart+v.Offset] = prefix[0:len(prefix)] + k + suffix[1:]
}

func singleSpace(s string) string {
	//fmt.Println("|", s, "|")
	spaces := getSpaces(s)
	r := spaces + strings.Join(strings.Fields(s), " ")
	//fmt.Println("|", r, "|")
	return r + "  "
}

func (v *Vim) Insert(k string) {
	if v.Menu != nil {
		if k == "ArrowUp" {
			v.Menu.Selected--
			Document.RenderToId("menu", "menu", v.Menu)
			return
		} else if k == "ArrowDown" {
			v.Menu.Selected++
			Document.RenderToId("menu", "menu", v.Menu)
			return
		}
	}
	if k == "Enter" && v.Menu != nil && v.Menu.Value() != "" {
		v.MenuDiv.Hide()
		s := v.getLine()
		prefix := s[0 : v.X+v.FocusLevel]
		suffix := s[v.X+v.FocusLevel:]
		newLine := prefix[0:len(prefix)-len(v.Menu.Search)] + v.Menu.Value() + " " + suffix
		v.SavedLines[v.Y+v.FocusStart+v.Offset] = newLine
		v.X = len(newLine) - 1 - v.FocusLevel
		v.Menu = nil
		return
	} else if k == "Enter" && v.Menu != nil {
		v.MenuDiv.Hide()
		v.InsertMode = false
		return
	} else if k == "Enter" {
		v.AddOneNewLine()
		return
	}
	if k == "ArrowUp" {
		return
	} else if k == "ArrowDown" {
		return
	} else if k == "ArrowRight" {
		return
	} else if k == "ArrowLeft" {
		return
	}

	s := v.getLine()
	//if strings.HasSuffix(s, "  ") && k == " " {
	//	return
	//}

	prefix := ""
	suffix := ""
	if len(s) > 0 {
		prefix = s[0 : v.X+v.FocusLevel]
		suffix = s[v.X+v.FocusLevel:]
	}

	//v.DebugLine = fmt.Sprintf("|%s|%s|%d|%d", prefix, suffix, v.X, v.X-v.FocusLevel)
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
	if len(tokens) > 1 {
		tokens = tokens[1:]
		last := tokens[len(tokens)-1]
		if last != "" {
			v.Menu = NewMenu(last)
			Document.RenderToId("menu", "menu", v.Menu)
			v.MenuDiv.Show()
		}
	}
	v.SavedLines[v.Y+v.FocusStart+v.Offset] = newLine
}

func (v *Vim) FocusLevelSpaces() string {
	buffer := []string{}
	for i := 0; i < v.FocusLevel; i++ {
		buffer = append(buffer, " ")
	}
	return strings.Join(buffer, "")
}
