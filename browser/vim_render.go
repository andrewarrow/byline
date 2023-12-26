package browser

import (
	"fmt"
)

func (v *Vim) Render() {
	v.Editor.Set("innerHTML", "")
	for i, line := range v.Lines {
		p := Document.NewTag("p", "")
		p.Set("id", fmt.Sprintf("p%d", i+1))
		if v.VisualMode && i == v.Y {
			p.AddClass("bg-gray-600")
		}
		//p.AddClass("whitespace-nowrap")
		v.Editor.AppendChild(p.JValue)

		for j, char := range line {
			s := fmt.Sprintf("%c", char)
			if s == " " {
				s = "&nbsp;"
			}
			span := Document.NewTag("span", s)
			span.Set("id", fmt.Sprintf("s%d-%d", i+1, j+1))
			if v.X == j && v.Y == i {
				span.AddClass("bg-white")
				span.AddClass("text-black")
			}
			p.AppendChild(span.JValue)
		}
	}
}
