package leetcode

func subtractProductAndSum(n int) int {
	var a int
	b:=1
	for n>0 {
		var m int
		 m = n%10
		 n/=10
		 a+= m
		 b*=m
	}
	return b-a
}