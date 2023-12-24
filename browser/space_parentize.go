package browser

import "strings"

func (s *Space) Parentize() {
	buffer := []string{}
	add := ""
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			buffer = append(buffer, spaces+"div")
			add = "  "
			buffer = append(buffer, add+line)
			continue
		}
		buffer = append(buffer, add+line)
	}
	s.Markup = strings.Join(buffer, "\n")
}
