package 剑指

import "strings"

func replaceSpace(s string) string {
	return strings.Replace(s," ","%20",-1)
}
