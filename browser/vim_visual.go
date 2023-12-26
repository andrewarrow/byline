package browser

func (v *Vim) VisualArrows(k string) {
	if k == "ArrowDown" {
		v.Y++
		v.ToY = v.Y
	}
	if k == "Enter" {
		return
	}
}
