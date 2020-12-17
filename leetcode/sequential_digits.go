package leetcode

// 1291. 顺次数
// 我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。
// 请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）。
//
// 输出：low = 100, high = 300
// 输出：[123,234]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sequential-digits
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func sequentialDigits(low int, high int) []int {
	var res []int
	value := firstSequentialDigits(low)
	if value > high {
		return res
	}
	res = append(res, value)
	bit := getBit(low)
	for {
		if bit >= 9 {
			break
		}
		step := getStep(getBit(value))
		value = value + step
		if value%10 == 0 {
			bit++
			value = getSequentialDigit(1, bit)
		}
		if value > high {
			break
		}
		res = append(res, value)
	}
	return res
}

func getStep(bit int) int {
	num := 1
	for i := 1; i < bit; i++ {
		num = num*10 + 1
	}
	return num
}

func firstSequentialDigits(low int) int {
	b := getBit(low)
	h := low
	for i := 1; i < b; i++ {
		h = h / 10
	}
	num := getSequentialDigit(h, b)

	if num < low {
		num = getSequentialDigit(h+1, b)
	}
	return num
}

func getSequentialDigit(start, bit int) int {
	if start+bit > 10 {
		return getSequentialDigit(1, bit+1)
	}
	num := start
	for i := 1; i < bit; i++ {
		num = num*10 + start + i
	}
	return num
}

func getBit(l int) int {
	i := 1
	for ; l/10 >= 1; i++ {
		l = l / 10
	}
	return i
}
