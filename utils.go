package pool

// distribute 任务分配
func distribute(from, to, groupNum int) [][2]int {

	var r [][2]int

	if (to - from) <= groupNum {
		panic("分配组数过多而任务跨度过小，无法进行均匀分配")
	}

	// 分块长度
	part := (to - from) / groupNum

	// 剩余分配部分
	restPart := (to - from) - part*groupNum

	// 上一次分配终点
	end := from

	for i := 1; i <= groupNum; i++ {

		// 分配宽度
		width := part

		if i <= restPart+1 {
			width++
		}

		temp := [2]int{
			end,
			end + width - 1,
		}

		end = end + width

		r = append(r, temp)
	}

	return r
}

// SetTask 设置任务区间
func (a *App) SetTask(from, to int) *App {
	a.task = [2]int{from, to}
	return a
}

// Task 获取子任务区间
func (a *App) Task(index int) (int, int) {
	result := distribute(a.task[0], a.task[1], a.Routine)[index]
	return result[0], result[1]
}
