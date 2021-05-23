package leetcode

func minOperations1(logs []string) int {
	if len(logs) == 0 {
		return 0
	}
	sum :=0
	for _,v := range logs{
		if v == "../" {
			if sum > 0{
				sum--
			}
		}else if v == "./"{
		}else {
			sum++
		}
	}
	return sum
}
