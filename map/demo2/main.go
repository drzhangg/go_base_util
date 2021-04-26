package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//var ms []map[string]interface{}
	//for i := 0; i < 10; i++ {
	//	m := make(map[string]interface{})
	//	m["key_"+fmt.Sprint(i)] = i
	//	ms = append(ms, m)
	//}
	//fmt.Println(ms)

	////s1 := reg.ReplaceAllString(s,"")
	//ss := reg.FindStringSubmatch(s)
	//fmt.Println(len(ss))
	//fmt.Println(ss)
	//fmt.Println(s1)

	//reg1 := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/(([\w-]+).([\w]+))$`)
	//s2 := reg1.FindStringSubmatch(s)
	//fmt.Println(s2)

	//flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	//params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang1-goque2ry-examples1-selector.html")
	//
	//for _,param :=range params {
	//	fmt.Println(param)
	//}

	//files := regexp.MustCompile(`^/([\w-]+)/(([\w-]+).([\w]+))$`)
	//pp := files.FindStringSubmatch("/us2r-loca3l/ab123c.sh")
	//fmt.Println(pp)

	//strings.Replace(url.Path, "/"+config.OssConfig.FileBucket+"/", "", 1)

	//sss := "We are happy."
	//ns := strings.Replace(sss," ","%20",-1)
	//fmt.Println(ns)
	//for k, v := range sss {
	//	fmt.Println("v:", string(v))
	//	if string(v) == " "{
	//		//strings.
	//		fmt.Println("k:",k)
	//	}
	//}

	//var ms = make([]map[int]string,0)
	var m = make(map[int]string)
	for i := 0; i < 5; i++ {

		m[i] = "int:" + strconv.Itoa(i)
		//ms = append(ms, m)
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v)
	}

	nums := []int{9,72,34,29,-49,-22,-77,-17,-66,-75,-44,-30,-24}

	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
	}


	if sum >0{
		fmt.Println()
	}else if sum == 0 {
		fmt.Println()
	}else {
		fmt.Println(-1)
	}

}

func IsIP(ip string) (b bool) {
	m, _ := regexp.MatchString(`^http://[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`, ip)
	if !m {
		return false
	} else {
		return true
	}
}
