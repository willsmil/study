package arithmetic

import "fmt"

// 已知sqrt(2)约等于1.414，要求不用数学库，求sqrt(2)精确到小数点后10位
func Sqrt2() {
	left := 1.414
	x := float64(10000)
	for b := 3; b < 11; b++ {
		for i := 0; i < 9; i++ {
			right := left + float64(i+1)/x
			if right*right > 2 {
				left = left + float64(i)/x
				fmt.Println(b, i, ">2", left)
				break
			}
			if i == 9 {
				left = left + float64(9)/x
			}
		}
		x = x * 10
	}
	fmt.Println(left)
}

func Sqrt() {
	left := 1.414
	x := float64(10000)
	for b := 3; b < 11; b++ {
		r := 9
		l := 0
		for true {
			tmp := (l + r) / 2
			right := left + float64(tmp)/x
			if tmp == l {
				left = right
				fmt.Println(b, tmp, ">2", left)
				break
			}
			if right*right > 2 {
				r = tmp
			} else {
				l = tmp
			}
		}

		x = x * 10
	}
	fmt.Println(left)
}
