package browser

import "github.com/andrewarrow/feedback/markup"

type ChangeMenu struct {
	Selected int
	Colors   []string
}

func NewChangeMenu() *ChangeMenu {
	m := ChangeMenu{}
	m.Colors = markup.Colors
	return &m
}

func (m *ChangeMenu) Template() string {
	return "color_menu"
}
