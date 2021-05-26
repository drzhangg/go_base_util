package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for _,v := range nums1{
		m1[v] = struct{}{}
	}

	result := []int{}
	for _,v := range nums2{
		m2[v] = struct{}{}
	}

	for k,_ := range m2{
		if _,ok := m1[k];ok{
			result = append(result, k)
		}
	}
	return result
}
