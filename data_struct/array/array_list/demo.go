package main

import (
	"encoding/json"
	"fmt"
)

var js = `[{"aaa":"3","bbb":199,"ccc":93.32},{"aaa":"2","bbb":199.16,"ccc":33},{"aaa":"6","bbb":199.1600,"ccc":22}]`

func main() {
	var months = []string{"1","2","3","4","5","6","7"}

	type data struct {
		Aaa string  `json:"aaa"`
		Bbb float64 `json:"bbb"`
		Ccc float64 `json:"ccc"`
	}

	var ds []data
	err := json.Unmarshal([]byte(js), &ds)
	if err != nil {
		fmt.Println("un err:", err)
	}

	var bb = make([]float64,len(months))
	var cc = make([]float64,len(months))
	for i := 0; i < len(months); i++ {
		for _,v := range ds{
			if months[i] == v.Aaa {
				bb[i] = v.Bbb
				cc[i] = v.Ccc
			}
		}
	}

	fmt.Println("bb:",bb)
	fmt.Println("cc:",cc)
}
