package browser

import "strings"

func (v *Vim) FixExtraSpaces() {
	buffer := []string{}
	for _, line := range v.SavedLines {
		spaces := getSpaces(line)
		r := spaces + strings.Join(strings.Fields(line), " ")
		buffer = append(buffer, r+" ")
	}
	v.SavedLines = buffer
}
