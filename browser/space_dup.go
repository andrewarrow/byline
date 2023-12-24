package browser

import (
	"strings"
)

func (s *Space) Duplicate() {
	lines := strings.Split(s.Markup, "\n")
	spaces := getSpaces(lines[s.CurrentLine])
	maxCount := len(spaces)

	bufferCopy := []string{}
	for i := s.CurrentLine; i < len(lines); i++ {
		line := lines[i]
		spaces := getSpaces(line)
		count := len(spaces)
		if count < maxCount {
			break
		}
		bufferCopy = append(bufferCopy, line)
	}

	buffer := []string{}
	fired := false
	for i, line := range lines {
		spaces := getSpaces(line)
		count := len(spaces)
		if count <= maxCount && i >= s.CurrentLine && fired == false {
			fired = true
			buffer = append(buffer, strings.Join(bufferCopy, "\n"))
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
	s.Render()
}
