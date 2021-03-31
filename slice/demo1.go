package main

import (
	"fmt"
)

import (
	"strings"
)

func main() {
	testStr := "_ga=GA1.1.43310309.1616569846; _hjid=14d4261c-e4f3-44a2-b417-c16200be883e; _agent_id=9422; _enterprise_id=5869; _VERSION_=legacy;_ga_Q21FPJS3FB=GS1.1.1616988393.1.1.1616988912.0; _authenticated=9428d35619fb871da96c8cf6fafc0479da5a82a8881da522ce8b800c6062e815; MEIQIA_TRACK_ID=1qT6jsqvGp91HLODREHs63rx54o; _ga_ZPLPLJ4XQG=GS1.1.1617184132.22.0.1617184132.0"
	strArr := strings.Split(testStr, ";")
	strMap := make(map[string]string, 20)
	for _, v := range strArr {
		tmp := strings.Split(v, "=")
		for i, k := range tmp {
			strMap[k] = tmp[i+1]
			fmt.Println(strMap)
			break
		}
		//continue
	}
	//fmt.Println("strmap;:::",strMap)

	fmt.Println(strMap["_ga_Q21FPJS3FB"])

}
