package browser

type Menu struct {
	Items    []string
	Selected int
}

func NewMenu() *Menu {
	m := Menu{}
	m.Items = []string{"text-center", "text-left", "text-right", "text-white", "text-black"}
	m.Selected = 0
	return &m
}

func (m *Menu) Value() string {
	return m.Items[m.Selected]
}
