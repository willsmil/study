package leetcode

func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	var idx int
	map1 := make(map[string]int)
	for i, l := range list1 {
		map1[l] = i
	}
	for i, l := range list2 {
		if _, ok := map1[l]; ok {
			if len(res) == 0 {
				res = append(res, l)
				idx = map1[l] + i
			} else {
				if map1[l]+i < idx {
					res = []string{l}
					idx = map1[l] + i
				} else if map1[l]+i == idx {
					res = append(res, l)
				}
			}
		}
	}
	return res
}
