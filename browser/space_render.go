package browser

import (
	"fmt"
	"strings"
)

func (s *Space) Render() {
	lines := strings.Split(s.Markup, "\n")
	for i, line := range lines {

		s := fmt.Sprintf(strings.ReplaceAll(line, " ", "&nbsp;"))
		p := Document.NewTag("p", s)
		p.Set("id", fmt.Sprintf("line%d", i+1))
	}
}
