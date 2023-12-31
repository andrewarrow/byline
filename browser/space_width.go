package browser

import (
	"strings"
)

type IncreaseDecrease struct {
	List   []string
	Prefix string
}

func (s *Space) Padding(val int) {
	s.IdLogic("p", spacing, val)
}

func (s *Space) Width(val int) {
	s.IdLogic("w", fractions, val)
}

func (s *Space) IdLogic(prefix string, list []string, val int) {
	id := IncreaseDecrease{}
	id.List = list
	id.Prefix = prefix

	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			m, tag := makeClassMap(line)
			w := id.find(m)
			if w+val >= 0 && w+val < len(id.List) {
				newW := id.List[w+val]
				m[id.Prefix+"-"+newW] = true
			}
			buffer = append(buffer, spaces+tag+" "+makeClasses(m))
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
}

func (id *IncreaseDecrease) find(m map[string]bool) int {
	for i, s := range id.List {
		key := id.Prefix + "-" + s
		if m[key] == true {
			m[key] = false
			return i
		}
	}

	return 0
}

var spacing = []string{
	"",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"8",
	"10",
	"12",
	"16",
	"20",
	"24",
	"32",
	"40",
	"48",
	"56",
	"64",
	"72",
	"80",
	"96",
}

var fractions = []string{
	"",
	"1/12",
	"1/6",
	"1/5",
	"1/4",
	"1/3",
	"2/5",
	"5/12",
	"3/5",
	"1/2",
	"7/12",
	"2/3",
	"3/4",
	"4/5",
	"5/6",
	"7/8",
	"11/12",
	"full",
}

var textSizes = []string{
	"",
	"xs",
	"sm",
	"base",
	"lg",
	"xl",
	"2xl",
	"3xl",
	"4xl",
	"5xl",
	"6xl",
}
