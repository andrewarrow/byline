package browser

import "strings"

type Menu struct {
	Items    []string
	Selected int
	Search   string
}

func NewMenu(s string) *Menu {
	m := Menu{}
	m.Search = s
	m.Items = []string{}
	m.FillItems()
	m.Selected = 0
	return &m
}

func (m *Menu) Value() string {
	return m.Items[m.Selected]
}

func (m *Menu) Filter(s string) {
	m.Search += s
	m.FillItems()
}
func (m *Menu) Backspace() {
	m.Search = m.Search[0 : len(m.Search)-1]
	m.FillItems()
}
func (m *Menu) FillItems() {
	m.Items = []string{}
	for _, item := range allItems {
		if strings.Contains(item, m.Search) {
			m.Items = append(m.Items, item)
		}
	}
}

var allItems = []string{
	"text-center", "text-left", "text-right", "text-white", "text-black",
	"cursor-pointer",
}
