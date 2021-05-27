package 剑指

//题解：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solution/mian-shi-ti-11-xuan-zhuan-shu-zu-de-zui-xiao-shu-3/
func minArray(numbers []int) int {
	n := len(numbers)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return numbers[0]
	}

	left,right := 0,n -1
	for left < right{
		mid := left + (right - left)/2
		if numbers[mid] < numbers[right]{
			right = mid
		}else if numbers[mid] > numbers[right] {
			left = mid + 1
		}else {//这里是numbers[mid] == numbers[right],两个值相等时，将right往左移
			right--
		}
	}
	return numbers[left]
}
