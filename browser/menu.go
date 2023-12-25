package browser

type Menu struct {
	Items    []string
	Selected int
}

func NewMenu() *Menu {
	m := Menu{}
	m.Items = []string{"wefwef", "wefwefwe"}
	m.Selected = 0
	return &m
}

func (m *Menu) Value() string {
	return m.Items[m.Selected]
}
