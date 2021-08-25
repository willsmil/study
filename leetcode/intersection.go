package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	var res []int
	m1 := make(map[int]struct{}, len(nums1))
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	m2 := make(map[int]struct{}, len(nums2))
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}
	for k, _ := range m2 {
		if _, ok := m1[k]; ok {
			res = append(res, k)
		}
	}
	return res
}
