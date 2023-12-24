package browser

import (
	"strings"
)

func (s *Space) Duplicate() {
	lines := strings.Split(s.Markup, "\n")
	spaces := getSpaces(lines[s.CurrentLine])
	maxCount := len(spaces)

	bufferCopy := []string{}
	count := 0
	for i := s.CurrentLine; i < len(lines); i++ {
		line := lines[i]
		spaces := getSpaces(line)
		count = len(spaces)
		if count < maxCount {
			break
		}
		//fmt.Println(count, line, maxCount)
		bufferCopy = append(bufferCopy, line)
	}
	if count == maxCount {
		bufferCopy = bufferCopy[0 : len(bufferCopy)-1]
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
