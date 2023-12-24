package browser

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

func (s *Space) Render() {
	s.Left.Set("innerHTML", "")
	lines := strings.Split(s.Markup, "\n")
	for i, line := range lines {

		str := fmt.Sprintf(strings.ReplaceAll(line, " ", "&nbsp;"))
		p := Document.NewTag("p", str)
		p.Set("id", fmt.Sprintf("line%d", i+1))
		s.Left.AppendChild(p)
	}

	m := map[string]any{}
	h := markup.ToHTMLFromLines(m, lines)
	s.Right.Set("innerHTML", h)
}
