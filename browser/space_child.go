package browser

import "strings"

func (s *Space) Child() {
	buffer := []string{}
	for i, line := range s.Lines {
		buffer = append(buffer, line)
		if i == s.CurrentLine {
			buffer = append(buffer, "  "+line)
		}
	}
	s.Markup = strings.Join(buffer, "\n")
}
