package jz

func replaceSpace(s string) string {
	// write code here
	var res []rune
	for _, i := range s {
		if i == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}
