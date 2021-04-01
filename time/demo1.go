package main

import (
	"fmt"
	"time"
)

func main() {

	var months []string
	var dataMap = make(map[string]int)
	for i := 30; i >= 0; i-- {
		t := time.Now()
		yes := t.AddDate(0, 0, -i)
		months = append(months, yes.Format("2006-01-02"))
		dataMap[yes.Format("2006-01-02")] = 0
	}

	fmt.Println(months)
	fmt.Println(dataMap)

}
