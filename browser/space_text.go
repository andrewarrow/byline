package browser

import "strings"

func (s *Space) SetText(t string) {
	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			buffer = append(buffer, spaces+t)
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
	s.Render()
}
