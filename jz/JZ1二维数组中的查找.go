package jz

// 1 2 3 4
// 2 4 7 9
func Find(target int, array [][]int) bool {
	height := len(array)
	if height == 0 {
		return false
	}
	width := len(array[0])
	if width == 0 {
		return false
	}
	if array[0][0] > target || array[height-1][width-1] < target {
		return false
	}
	i := 0
	j := width - 1
	for i < height && j >= 0 {
		if array[i][j] == target {
			return true
		} else if array[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
