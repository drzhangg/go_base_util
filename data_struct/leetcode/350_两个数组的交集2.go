package leetcode

import "sort"

/*
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
*/
// 题解：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/solution/liang-ge-shu-zu-de-jiao-ji-ii-by-leetcode-solution/
//1.排序+双指针
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	index1, index2 := 0, 0

	result := []int{}

	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			result = append(result, nums2[index2])
			index1++
			index2++
		}
	}
	return result
}

//map对比
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		intersect2(nums2, nums1)
	}

	var m1 = make(map[int]int)
	for _, v := range nums1 {
		m1[v]++
	}

	var result []int
	for _,v := range nums2{
		if m1[v] > 0 {
			m1[v]--
			result = append(result, v)
		}
	}
	return result
}
