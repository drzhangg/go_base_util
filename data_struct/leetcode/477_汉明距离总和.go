package leetcode

func totalHammingDistance(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum += getSum(nums[i],nums[j])
		}
	}
	return sum
}

func getSum(x, y int) int {
	num := x ^ y
	var count int
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
