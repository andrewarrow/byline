package browser

type Menu struct {
	Items    []string
	Selected int
	Search   string
}

func NewMenu(s string) *Menu {
	m := Menu{}
	m.Search = s
	m.Items = []string{"text-center", "text-left", "text-right", "text-white", "text-black"}
	m.Selected = 0
	return &m
}

func (m *Menu) Value() string {
	return m.Items[m.Selected]
}

func (m *Menu) Filter(s string) {
	m.Search += s
	m.Items = []string{"foo-text-right", "text-white", "text-black"}
}
func (m *Menu) Backspace() {
	m.Search = m.Search[0 : len(m.Search)-1]
	m.Items = []string{"bar-text-right", "text-white", "text-black"}
}
