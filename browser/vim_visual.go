package browser

func (v *Vim) VisualArrows(k string) {
	if k == "ArrowDown" {
		v.Y++
		v.ToY = v.Y
	} else if k == "ArrowUp" {
		v.Y--
		v.FromY = v.Y
	}
	if k == "Enter" {
		return
	}
}
