package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//var ms []map[string]interface{}
	//for i := 0; i < 10; i++ {
	//	m := make(map[string]interface{})
	//	m["key_"+fmt.Sprint(i)] = i
	//	ms = append(ms, m)
	//}
	//fmt.Println(ms)


	var s = "http://100.73.70.169:7777/file-manage-bucket/2021-03-31/b2fe0403-a420-4aab-8a54-2ad7e6d1523e.s"


	as := strings.SplitN(s,"/",-1)
	fmt.Println(as)


	reg := regexp.MustCompile(`^http://100.73.70.169:7777/([\w-]+)/([\w-]+)/([\w-]+).([\w]+)$`)
	//s1 := reg.ReplaceAllString(s,"")
	ss := reg.FindStringSubmatch(s)
	fmt.Println(len(ss))
	fmt.Println(ss)
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

}

func IsIP(ip string) (b bool) {
	m, _ := regexp.MatchString(`^http://[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`, ip)
	if !m {
		return false
	} else {
		return true
	}
}
