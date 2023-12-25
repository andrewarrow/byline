package browser

import "github.com/andrewarrow/feedback/markup"

type ChangeMenu struct {
	Selected int
	Colors   []string
	Value    int
}

func NewChangeMenu() *ChangeMenu {
	m := ChangeMenu{}
	m.Value = 100
	m.Colors = markup.Colors
	return &m
}

func (m *ChangeMenu) Template() string {
	return "color_menu"
}
