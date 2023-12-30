package browser

func (v *Vim) GrowTag(k string) {
	if k == "ArrowUp" {
	} else if k == "ArrowDown" {
	} else if k == "ArrowRight" {
	} else if k == "ArrowLeft" {
	}
	v.Render()
	leaveInsertMode()
}
