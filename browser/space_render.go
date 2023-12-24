package browser

import (
	"fmt"
	"strings"
)

func (s *Space) Render() {
	lines := strings.Split(s.Markup, "\n")
	buffer := []string{}
	for i, line := range lines {

		p := fmt.Sprintf(`p id=line%d`+"\n", i+i)
		s := fmt.Sprintf(strings.ReplaceAll(line, " ", "&nbsp;"))
		buffer = append(buffer, p+s)
	}
	s.Left.Set("innerHTML", strings.Join(buffer, "\n"))
}
