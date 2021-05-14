package leetcode

func intToRoman(num int) string {
	var roman = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string
	for i := 0; i < len(roman); i++ {
		for num >= values[i] {
			num -= values[i]
			result += roman[i]
		}
	}
	return result
}
