package browser

func (v *Vim) VisualArrows(k string) {
	if k == "ArrowDown" {
		v.Y++
		v.ToY = v.Y
	} else if k == "y" {
		v.VisualMode = false
	} else if k == "d" {
		v.Lines = append(v.Lines[0:v.FromY], v.Lines[v.ToY+1:]...)
		v.VisualMode = false
	} else if k == "ArrowUp" {
		v.Y--
		//fmt.Println(v.Y, v.FromY, v.ToY)
		v.FromY = v.Y
		//fmt.Println(v.Y, v.FromY, v.ToY)
	}
	if k == "Enter" {
		return
	}
}
