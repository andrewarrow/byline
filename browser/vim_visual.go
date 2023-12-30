package browser

func (v *Vim) VisualArrows(k string) {
	if k == "ArrowDown" {
		v.Y++
		v.EndY++
	} else if k == "y" {
		v.VisualMode = false
		//start, end := blockOfLines(v.StartY, v.EndY)
		//lines := v.Lines[start : end+1]
		//v.Yanked = append([]string{}, lines...)
	} else if k == ">" {
		v.VisualMode = false
		//start, end := blockOfLines(v.StartY, v.EndY)
		//op := NewOperation("indent_lines")
		//op.From = start
		//op.To = end
		//vim.RunOp(op)
	} else if k == "d" {
		v.VisualMode = false
		start, end := blockOfLines(v.StartY, v.EndY)
		lines := v.OffsetLines[start : end+1]
		v.Yanked = append([]string{}, lines...)

		op := NewOperation("remove_lines")
		op.Data = v.Yanked
		op.InsertY = start + v.Offset
		vim.RunOp(op)
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
