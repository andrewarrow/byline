package browser

func (v *Vim) VisualArrows(k string) {
	if k == "ArrowDown" {
		v.Y++
		v.EndY++
	} else if k == "y" {
		v.VisualMode = false
		start, end := blockOfLines(v.StartY, v.EndY)
		lines := v.Lines[start : end+1]
		v.Yanked = append([]string{}, lines...)
	} else if k == "d" {
		/*
			v.DeletedLines = append([]string{}, v.Lines[v.FromY:v.ToY+1]...)
			fmt.Println("1", v.DeletedLines)
			v.Lines = append(v.Lines[0:v.FromY], v.Lines[v.ToY+1:]...)
			v.VisualMode = false
		*/
	} else if k == "ArrowUp" {
		v.Y--
		//fmt.Println(v.Y, v.FromY, v.ToY)
		v.EndY--
		//fmt.Println(v.Y, v.FromY, v.ToY)
	}
	if k == "Enter" {
		return
	}
}

func blockOfLines(start, end int) (int, int) {
	if start > end {
		return end, start
	} else if start < end {
		return start, end
	}
	return start, start
}
