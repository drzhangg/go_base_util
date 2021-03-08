package 数组

func intersect(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2,nums1)
	}

	m := map[int]int{}
	for _,v := range nums1{
		m[v]++
	}

	intersection := []int{}
	for _,v := range nums2{
		if m[v] > 0{
			intersection = append(intersection, v)
			v--
		}
	}
	return intersection
}
