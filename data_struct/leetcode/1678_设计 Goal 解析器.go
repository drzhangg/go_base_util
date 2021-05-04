package leetcode

import "strings"

func interpret(command string) string {
	newStr := strings.Split(command,"")
	var r []string
	for i := 0; i < len(newStr); i++ {
		if newStr[i] == "G"{
			r = append(r, "G")
		}else if newStr[i] =="(" && newStr[i+1] ==")"{
			r = append(r, "o")
		}else if newStr[i]=="(" && newStr[i+1] =="a" {
			r = append(r, "al")
		}
	}
	return strings.Join(r,"")
}
