package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) Render() {
	v.Editor.Set("innerHTML", "")
	for i, line := range v.Lines {
		str := fmt.Sprintf(strings.ReplaceAll(line, " ", "&nbsp;"))
		p := Document.NewTag("p", str)
		p.Set("id", fmt.Sprintf("line%d", i+1))
		//p.AddClass("whitespace-nowrap")
		v.Editor.AppendChild(p.JValue)
	}
}
