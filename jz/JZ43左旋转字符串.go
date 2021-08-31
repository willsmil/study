package jz

func LeftRotateString(str string, n int) string {
	// write code here
	ss := []byte(str)
	if n > len(ss) {
		reverse(ss, 0, len(ss)-1)
	} else {
		reverse(ss, 0, n-1)
		reverse(ss, n, len(ss)-1)
		reverse(ss, 0, len(ss)-1)
	}
	return string(ss)
}

func reverse(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
